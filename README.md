# RoadGuard User Service

This service provides user authentication and registration APIs for the RoadGuard platform, leveraging AWS Cognito for secure user management. Built with Go and the Gin web framework.

---

## Features

- **User Signup:** Register new users with email and password.
- **User Login:** Authenticate users and return JWT tokens (access and ID tokens).
- **AWS Cognito Integration:** Secure, scalable user management.

---

## Project Structure

```
user_service/
  ├── .env
  ├── main.go
  ├── auth/
  │   ├── cognito.go
  │   └── config.go
  ├── handlers/
  │   └── user.go
  └── models/
      └── user.go
```

---

## Getting Started

### Prerequisites

- Go 1.18 or newer
- AWS account with Cognito User Pool
- AWS credentials configured locally
- [Gin](https://github.com/gin-gonic/gin) and [godotenv](https://github.com/joho/godotenv) Go modules

### Setup

1. **Clone the repository.**
2. **Configure environment variables:**  
   Create a `.env` file in `user_service/` with:
   ```
   AWS_REGION=ap-south-1
   COGNITO_USER_POOL_ID=ap-south-1_XXXXXXX
   COGNITO_APP_CLIENT_ID=YYYYYYYYYYYY
   ```
3. **Install dependencies:**
   ```sh
   go mod tidy
   ```
4. **Run the service:**
   ```sh
   go run user_service/main.go
   ```
   The service will start on `localhost:8080`.

---

## API Endpoints

### `POST /signup`

Registers a new user.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Response:**
- `200 OK` on success
- `400 Bad Request` on error

---

### `POST /login`

Authenticates a user and returns JWT tokens.

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```

**Response:**
- `200 OK` with `access_token` and `id_token`
- `401 Unauthorized` on error

---

## Roadmap / Next Steps

- [ ] Add email confirmation endpoint
- [ ] Implement password reset functionality
- [ ] Add user profile management APIs
- [ ] Write unit and integration tests
- [ ] Add API documentation (Swagger/OpenAPI)
- [ ] Dockerize the service for deployment
- [ ] Add CI/CD pipeline

---

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements and bug fixes.

---

## License

This project is licensed under the