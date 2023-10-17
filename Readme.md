# Project Title

TEST CRUD USERS + AUTH GOLANG GIN

- Import File Postman API Documentation from folder

  ```
  docs/Test API Auth.postman_collection,json
  ```

## Installation

Please copy file .env.example to .env

```
  cp .env.example to env
```

Create Database Postgres and rename in file .env

```
  DB_USER=postgres
  DB_PASS=Your_DATABASE_PASS
  DB_NAME=Your_DATABASE_NAME
  DB_HOST=localhost
  DB_PORT=5432
```

Add JWT Config in env file.

```
  JWT_SECRET=test_crud_api // example value
  JWT_EXPIRED_TOKEN=60
```

And Finally Run :

```bash
  go mod download
  go run main.go
```

## API Documentation

#### Login

```http
  POST /api/login
```

| Body Json  | Type     | Description   |
| :--------- | :------- | :------------ |
| `email`    | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### Register

```http
  POST /api/Register
```

| Body Json          | Type     | Description   |
| :----------------- | :------- | :------------ |
| `name`             | `string` | **Required**. |
| `email`            | `string` | **Required**. |
| `password`         | `string` | **Required**. |
| `password_confirm` | `string` | **Required**. |

#### Get All User

```http
  GET /api/users
```

#### Detail User

```http
  GET /api/users/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of user to fetch |

#### Update User

```http
  PATCH /api/users/:id
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

```
Payload Body Update User
```

| Body Json          | Type     | Description   |
| :----------------- | :------- | :------------ |
| `name`             | `string` | **Required**. |
| `password`         | `string` | **Required**. |
| `password_confirm` | `string` | **Required**. |

#### Delete User

```http
  DELETE /api/users/:id
```

| Parameter | Type     | Description              |
| :-------- | :------- | :----------------------- |
| `id`      | `string` | **Required**. Id of user |

## Authors

- [@lailiseptiand](https://www.github.com/lailiseptiandi)

## Features

- Login/Register Auth
- Middleware
- CRUD Users

## Tech Stack

**API:** Go, Postgress
