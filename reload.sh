#!/bin/bash

echo "[SCRIPT] Rebuilding and restarting Docker containers..."
sudo docker-compose down
sudo docker-compose up --build -d

if ! sudo docker-compose ps | grep "Up"; then
  echo "[SCRIPT] Docker containers failed to start. Check logs with 'docker-compose logs'."
  exit 1
fi

echo "[SCRIPT] Update complete."
