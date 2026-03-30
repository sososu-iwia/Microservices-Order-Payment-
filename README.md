# AP2 Assignment 2 — Microservices (Order & Payment)

##  Overview

This project implements a simple **microservices architecture** using Go, PostgreSQL, and Docker.

It consists of two independent services:

* **Order Service** — manages orders
* **Payment Service** — processes payments for orders

Each service has its own database and communicates through HTTP APIs.

---

##  Architecture

* **Order Service**

    * Creates and retrieves orders
    * Triggers payment creation

* **Payment Service**

    * Stores and retrieves payment information
    * Linked to orders via `order_id`

* **PostgreSQL**

    * Separate database for each service:

        * `order_db`
        * `payment_db`

* **Docker Compose**

    * Runs all services and databases together

---

##  How to Run

### 1. Clone the repository

```bash
git clone <your-repo-link>
cd ap2_assignment 2
```

### 2. Start the project

```bash
docker compose up --build
```

---

##  Services

| Service         | URL                   |
| --------------- | --------------------- |
| Order Service   | http://localhost:8081 |
| Payment Service | http://localhost:8082 |

---

##  API Endpoints

###  Order Service (Port 8081)

| Method | Endpoint            | Description      |
| ------ | ------------------- | ---------------- |
| POST   | /orders             | Create new order |
| GET    | /orders/{id}        | Get order by ID  |
| PATCH  | /orders/{id}/cancel | Cancel order     |
| GET    | /health             | Health check     |

---

###  Payment Service (Port 8082)

| Method | Endpoint            | Description             |
| ------ | ------------------- | ----------------------- |
| POST   | /payments           | Create payment          |
| GET    | /payments/{orderID} | Get payment by order ID |
| GET    | /health             | Health check            |

---

## Example Usage

###  Check services

```bash
curl http://localhost:8081/health
curl http://localhost:8082/health
```

---

###  Create Order

```bash
curl -X POST http://localhost:8081/orders \
  -H "Content-Type: application/json" \
  -H "Idempotency-Key: test-key-1" \
  -d '{
    "customer_id": "cust_1",
    "item_name": "iPhone",
    "amount": 100000
  }'
```

---

### Get Order

```bash
curl http://localhost:8081/orders/{order_id}
```

---

###  Get Payment

```bash
curl http://localhost:8082/payments/{order_id}
```

---

## How It Works

1. A client sends a request to create an order (`POST /orders`)
2. Order Service:

    * Saves order in `order_db`
    * Automatically triggers payment creation
3. Payment Service:

    * Creates a payment record in `payment_db`
4. Both services store and return data via REST API

---

##  Database

Each service uses its own PostgreSQL database:

* `order_db`

    * table: `orders`

* `payment_db`

    * table: `payments`

Migrations are applied automatically on startup.

---

##  Key Features

* Microservices architecture
* Independent databases
* REST API with Gin (Go)
* Dockerized environment
* Database migrations
* Idempotency support via headers

---

##  Technologies

* Go (Golang)
* Gin Web Framework
* PostgreSQL
* Docker & Docker Compose

---

##  Project Structure

```
ap2_assignment 2/
│
├── docker-compose.yml
│
├── order-service/
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── Dockerfile
│   └── go.mod
│
├── payment-service/
│   ├── cmd/
│   ├── internal/
│   ├── migrations/
│   ├── Dockerfile
│   └── go.mod
```

---

## Testing

The project can be tested using:

* curl (CLI)
* Postman

---

##  Conclusion

This project demonstrates:

* Microservices design
* Service-to-service interaction
* Database separation
* Docker-based deployment

---

##  Author

Damir Erbolatov
