#!/bin/bash

GITHUB_USER="..."
GITHUB_TOKEN="..."

JWT_SECRET="..." # openssl rand -base64 32

DB_USER="..."
DB_PASS="..."
DB_NAME="..."

DOMAIN_NAME="..."
EMAIL="..."

NEXTAUTH_URL="..."
GOOGLE_MAPS_KEY="..."

API_URL="http://localhost:3333"
FRONT_URL="http://localhost:3000"

###
## UPDATE SYSTEM
###

echo "[SCRIPT] Updating packages"
sudo apt update

echo "[SCRIPT] Upgrading packages"
sudo apt upgrade


###
## CHECKING FOR AVAILABLE PACKAGES
###

echo "[SCRIPT] Searching for GIT"
HAS_GIT=false

if [ "$(dpkg -l git)" ]; then
  HAS_GIT=true
fi

if [ $HAS_GIT != true ]; then
  echo "[SCRIPT] GIT isn't available, installing..."
  sudo apt install -y git
fi

# ------------------------------------------------- #

echo "[SCRIPT] Searching for docker"
HAS_DOCKER=false

if [ "$(dpkg -l docker)" ]; then
  HAS_DOCKER=true
if

if [ $HAS_DOCKER != true ]; then
  echo "[SCRIPT] Docker isn't available, installing..."

  sudo apt install apt-transport-https ca-certificates curl software-properties-common -y
  curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" -y
  sudo apt update && sudo apt install docker-ce -y

  sudo rm -f /usr/local/bin/docker-compose
  sudo curl -L "https://github.com/docker/compose/releases/download/v2.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

  # Wait for the file to be fully downloaded before proceeding
  if [ ! -f /usr/local/bin/docker-compose ]; then
    echo "Docker Compose download failed. Exiting."
    exit 1
  fi

  sudo chmod +x /usr/local/bin/docker-compose

  docker-compose --version
  if [ $? -ne 0 ]; then
    echo "Docker Compose installation failed. Exiting."
    exit 1
  fi

  # Ensure Docker starts on boot and start Docker service
  sudo systemctl enable docker
  sudo systemctl start docker
fi

echo "[SCRIPT] Searching for NGINX"
HAS_NGINX=false

if [ "$(dpkg -l nginx)" ]; then
  HAS_NGINX=true
if

if [ $HAS_NGINX != true ]; then
  echo "[SCRIPT] Docker isn't available, installing..."
  sudo apt install nginx -y

  sudo rm -f /etc/nginx/sites-available/myapp
  sudo rm -f /etc/nginx/sites-enabled/myapp

  sudo systemctl stop nginx

  sudo apt install certbot -y
  sudo certbot certonly --standalone -d $DOMAIN_NAME --non-interactive --agree-tos -m $EMAIL

  # Ensure SSL files exist or generate them
  if [ ! -f /etc/letsencrypt/options-ssl-nginx.conf ]; then
    sudo wget https://raw.githubusercontent.com/certbot/certbot/main/certbot-nginx/certbot_nginx/_internal/tls_configs/options-ssl-nginx.conf -P /etc/letsencrypt/
  fi

  if [ ! -f /etc/letsencrypt/ssl-dhparams.pem ]; then
    sudo openssl dhparam -out /etc/letsencrypt/ssl-dhparams.pem 2048
  fi

  sudo cat > /etc/nginx/sites-available/myapp <<EOL
  limit_req_zone \$binary_remote_addr zone=mylimit:10m rate=10r/s;

  server {
      listen 80;
      server_name $DOMAIN_NAME;

      # Redirect all HTTP requests to HTTPS
      return 301 https://\$host\$request_uri;
  }

  server {
      listen 443 ssl;
      server_name $DOMAIN_NAME;

      ssl_certificate /etc/letsencrypt/live/$DOMAIN_NAME/fullchain.pem;
      ssl_certificate_key /etc/letsencrypt/live/$DOMAIN_NAME/privkey.pem;
      include /etc/letsencrypt/options-ssl-nginx.conf;
      ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

      # Enable rate limiting
      limit_req zone=mylimit burst=20 nodelay;

      location / {
          proxy_pass $FRONT_URL;
          proxy_http_version 1.1;
          proxy_set_header Upgrade \$http_upgrade;
          proxy_set_header Connection 'upgrade';
          proxy_set_header Host \$host;
          proxy_cache_bypass \$http_upgrade;

          # Disable buffering for streaming support
          proxy_buffering off;
          proxy_set_header X-Accel-Buffering no;
      }

      location /api/ {
          proxy_pass $API_URL;
          proxy_http_version 1.1;
          proxy_set_header Upgrade \$http_upgrade;
          proxy_set_header Connection 'upgrade';
          proxy_set_header Host \$host;
          proxy_cache_bypass \$http_upgrade;
      }
  }
  EOL

  # Create symbolic link if it doesn't already exist
  sudo ln -s /etc/nginx/sites-available/myapp /etc/nginx/sites-enabled/myapp

  # Restart Nginx to apply the new configuration
  sudo systemctl restart nginx
fi

# -------------------------------------------------- #

###
## CLONING REPOSITORIES
###

BACKEND_PATH="2ndevs/imobiliaria-terreno-api"
FRONTEND_PATH="2ndevs/imobiliaria-terreno-web"

BACKEND_DIR="backend"
FRONTEND_DIR="frontend"

if [ -d $BACKEND_DIR ]; then
  echo "[SCRIPT] Removing old backend folder"
  rm -rf $BACKEND_DIR
fi

if [ -d $FRONTEND_DIR ]; then
  echo "[SCRIPT] Removing old frontend folder"
  rm -rf $FRONTEND_DIR
fi

echo "[SCRIPT] Cloning backend repository"
git clone https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/${BACKEND_PATH}
mv imobiliaria-terreno-api $BACKEND_DIR

echo "[SCRIPT] Setting up backend environment variables"
echo "APP_PORT=3333" > "$BACKEND_DIR/.env"
echo "NODE_ENV='production'" >> "$BACKEND_DIR/.env"
echo "JWT_SECRET=$JWT_SECRET" >> "$BACKEND_DIR/.env" # openssl rand -base64 32
echo "POSTGRES_USER=$DB_USER" >> "$BACKEND_DIR/.env"
echo "POSTGRES_PASSWORD=$DB_PASS" >> "$BACKEND_DIR/.env"
echo "POSTGRES_DB=$DB_NAME" >> "$BACKEND_DIR/.env"
echo "POSTGRES_HOST='localhost'" >> "$BACKEND_DIR/.env"
echo "POSTGRES_PORT=5432" >> "$BACKEND_DIR/.env"
echo "DATABASE_URL=postgres://\${POSTGRES_USER}:\${POSTGRES_PASSWORD}@localhost:5432/\${POSTGRES_DB}" >> "$BACKEND_DIR/.env"

echo "[SCRIPT] Running database docker"
docker compose -f $BACKEND_DIR/docker-compose.yml up -d

if ! sudo docker-compose ps | grep "Up"; then
  echo "Docker containers failed to start. Check logs with 'docker-compose logs'."
  exit 1
fi

echo "[SCRIPT] Cloning frontend repository"
git clone https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/${FRONTEND_PATH}
mv imobiliaria-terreno-web $FRONTEND_DIR

echo "[SCRIPT] Setting up frontend environment variables"
echo "NEXT_PUBLIC_GOOGLE_MAPS_KEY=$GOOGLE_MAPS_KEY" >> "$FRONTEND_DIR/.env"
echo "NEXTAUTH_URL=$NEXTAUTH_URL" >> "$FRONTEND_DIR/.env"
