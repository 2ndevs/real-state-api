<br/>

<div align="center">
  <strong>
    PREVIEW ONLY PROJECT, READ MORE <a href="https://github.com/2ndevs/real-state-api/blob/main/LICENSE">HERE</a>
  </strong>
</div>

<br/>

# Real Estate API

This project is an API designed for a real estate website where users can search for their ideal property in a specific city. The API is built using **Bun** and **Fastify**, with a **PostgreSQL** database running in a **Docker Compose** environment. The API provides separate routes for web users and administrators, allowing for property management and report generation.

---

## **Table of Contents**

- [Project Overview](#project-overview)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
  - [Web Routes](#web-routes)
  - [Admin Routes](#admin-routes)
- [Testing](#testing)

---

## **Project Overview**

This API allows users to search for properties, view detailed property information, and contact real estate agents. Admin users can manage property listings, generate reports, and control user access. Google Maps integration enables an interactive map with property listings.

---

## **Technologies Used**

- **Go**: A fast garbage collected language and memory efficient.
- **CHI**: Low-overhead backend framework for building APIs.
- **PostgreSQL**: Relational database management system.
- **Docker Compose**: Tool to define and manage multi-container Docker applications.
- **JWT**: JSON Web Tokens for authentication and secure access.

---

## **Getting Started**

### Prerequisites

- **Go**: Install Bun by following the [official guide](https://bun.sh/).
- **Docker**: Make sure Docker and Docker Compose are installed on your machine.

---

## **Environment Variables**

Before running the project, you need to configure the environment variables in a `.env` file. Here's an example configuration:

```env
# .env

# Application port
APP_PORT=3333

# Node environment
NODE_ENV='development'

# JWT secret key (generate a secure key using `openssl rand -base64 32`)
JWT_SECRET='your_random_secret'

# PostgreSQL database configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=realestate_db
```

### **Generating JWT Secret**

To generate a secure JWT secret for production, use the following command:

```bash
openssl rand -base64 32
```

---

## **Running the Project**

### Step 1: Clone the Repository

```bash
git clone https://github.com/your-repo/real-estate-api.git
cd real-estate-api
```

### Step 2: Install Dependencies

```bash
go get
```

### Step 3: Set Up the Database

Ensure Docker is running and use the Docker Compose configuration to set up the PostgreSQL database:

```bash
docker-compose up -d
```

### Step 4: Run Migrations

After the database is running, run the migrations to set up the tables:

```bash
#TODO
```

### Step 5: Start the API

To run the server in development mode:

```bash
go run .
```

The API will be available at `http://localhost:3333`.

---

## **API Endpoints**

### **Web Routes**

These routes are accessible to regular users of the real estate site.

#### `GET /properties`

- **Description**: Get a list of properties available for sale or rent.
- **Query Parameters**:
  - `city`: Filter by city.
  - `price_min` & `price_max`: Filter by price range.
  - `type`: Filter by property type (e.g., house, apartment).
- **Response**: Array of properties.

#### `GET /properties/:id`

- **Description**: Get detailed information about a specific property.
- **URL Parameters**: 
  - `id`: Property ID.
- **Response**: Detailed property information, including location, features, and price.

#### `POST /contact`

- **Description**: Send a message to the agent responsible for the property.
- **Body**:
  - `propertyId`: ID of the property the message is about.
  - `name`: Sender's name.
  - `email`: Sender's email.
  - `message`: Message content.
- **Response**: Confirmation of message sent.

### **Admin Routes**

These routes require JWT authentication and are for administrative users.

#### `POST /admin/properties`

- **Description**: Create a new property listing.
- **Body**:
  - `name`: Property name.
  - `description`: Property description.
  - `type`: Property type (e.g., house, apartment).
  - `price`: Property price.
  - `location`: Property address.
  - `features`: Array of property features (e.g., number of rooms, area).
- **Response**: Newly created property.

#### `PUT /admin/properties/:id`

- **Description**: Update the information of an existing property.
- **URL Parameters**: 
  - `id`: Property ID.
- **Body**:
  - `name`: Property name (optional).
  - `description`: Property description (optional).
  - `price`: Property price (optional).
  - Any other field to update.
- **Response**: Updated property.

#### `DELETE /admin/properties/:id`

- **Description**: Delete a property listing.
- **URL Parameters**: 
  - `id`: Property ID.
- **Response**: Confirmation of deletion.

#### `GET /admin/reports`

- **Description**: Generate property reports based on filters like status (available, sold), date range, etc.
- **Response**: Report data (e.g., PDF or Excel export).

---

## **Testing**

You can run tests for the API endpoints using:

```bash
go test
```
