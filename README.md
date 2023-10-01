# Application Packaging

## Introduction

The Application Packaging application is designed with Domain Driven Design (DDD) principles in mind. Its main objective is to optimize the packaging of items to fulfill orders while adhering to specific rules.

### Rules

The application operates based on the following rules:

1. **Rule 1:** Only whole packs can be sent. Packs cannot be broken open.
2. **Rule 2:** Within the constraints of Rule 1, send out no more items than necessary to fulfill the order.
3. **Rule 3:** Within the constraints of Rules 1 & 2, send out as few packs as possible to fulfill each order.

## Endpoints

The Application Packaging application provides the following endpoints:

### Health Check

- **Endpoint:** `GET http://localhost:7070/health`
- Use this endpoint to check the health status of the application.

### Add Packages

- **Endpoint:** `POST http://localhost:7070/add-packages`
- Use this endpoint to add available package sizes to the application.

#### Example CURL Request:

```bash
curl --location 'http://localhost:7070/add-packages' \
--header 'Content-Type: text/plain' \
--data '{
  "packages": [
    {"size": 250},
    {"size": 500},
    {"size": 1000},
    {"size": 2000},
    {"size": 5000}
  ]
}'
```
### Add Packages

- **Endpoint:** `POST http://localhost:7070/calculate-packages`
- Use this endpoint to calculate the optimal packaging for a given order amount.

#### Example CURL Request:
```bash
curl --location 'http://localhost:7070/calculate-packages' \
--header 'Content-Type: text/plain' \
--data '{
  "amount": 12001
}'
```

### Getting Started
To get started with the Application Packaging application, follow these steps:

1. Clone the application repository.
2. Install the required dependencies.
3. Run the application locally.

### Usage
- Once the application is running, you can use the provided CURL examples to interact with the endpoints. The calculate-packages endpoint is especially useful for optimizing the packaging of orders according to the specified rules.

## Docker
- build image `docker build -t packaging:latest .`
- run image `docker run -p 7070:7070 packaging:latest`
