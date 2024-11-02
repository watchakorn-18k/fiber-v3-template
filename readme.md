# Go Fiber v3 Template

## Getting Started
I have developed a web API structure tailored for microservices using Fiber v3. This version introduces various enhancements and modifications, making it even more effective. I hope this resource proves valuable to anyone exploring the updated features and methods.

### Prerequisites

- Go 1.23+

### Installation and Run

1. Clone the repository
```bash
git clone https://github.com/gofiber/fiber.git
```

2. Change directory
```bash
cd fiber
```

3. Run the server
```bash
go run .
```

## Docker Build and Run
**Note:** You can use `go mod vendor` before building the image to reduce the time it takes to download modules to virual machine

1. Build the image
```bash
docker build -t fiber_v3 .
```

2. Run the image
```bash
docker run -it --rm -p 1818:1818 --env-file .env fiber_v3
```

## Podman Build and Run
**Note:** You can use `go mod vendor` before building the image to reduce the time it takes to download modules to virual machine

1. Build the image
```bash
podman build -t fiber_v3 .
```

2. Run the image
```bash
podman run -it --rm -p 1818:1818 --env-file .env fiber_v3
```

# API

## GET http://127.0.0.1:1818/api/number/

### Response

```json
{
  "message": "Hello, World!"
}
```

## GET http://127.0.0.1:1818/api/number/random_number

### Response

```json
{
  "message": "Random number generated successfully",
  "data": 123456789
}
```

## POST http://127.0.0.1:1818/api/number/calculate

### Payload

```json
{
  "a": 12,
  "b": 22
}
```

### Response

```json
{
  "message": "Calculated successfully",
  "data": 24
}
```

## GET http://127.0.0.1:1818/api/carbon/check_carbon?url=https://www.google.com

### Response

```json
{
  "message": "Carbon calculated successfully",
  "data": {
    "url": "https://www.google.com/",
    "green": true,
    "bytes": 1024887,
    "cleanerThan": 0.76,
    "rating": "B"
  }
}
```
