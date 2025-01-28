# Golang JWT Authentication Server

## To run the server and db:

```
docker-compose up --build
```

## Technologies

- Go
- Postgres
- Docker

## Postman Collection

- **GET /auth** - Check if the server is running.
- **GET /auth/authenticated-user** - Retrieve details of the authenticated user.
- **POST /auth/register** - Register a new user.
- **POST /auth/login** - Log in with credentials.
- **POST /auth/logout** - Log out the current user.
