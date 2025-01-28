# React and Golang JWT Authentication

https://youtu.be/d4Y2DkKbxM0?si=axM0D7aLAdvh9Dn8

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
