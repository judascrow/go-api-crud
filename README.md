# Golang API Example

##### Golang v.1.14

## COMMAND

Migrate & Seed :

```
go run main.go create seed
```

Run :

```
go run main.go
```

## API ROUTER URI

```
Host: localhost:8000
Base Path: /api/v1
```

### Healthcheck

| Done               | Method | URI          | authorize | comment              |
| ------------------ | ------ | ------------ | --------- | -------------------- |
| :white_check_mark: | GET    | /healthcheck | No        | ครวจสอบสถานะ Service |

### Auth

| Done               | Method | URI    | authorize | comment |
| ------------------ | ------ | ------ | --------- | ------- |
| :white_check_mark: | POST   | /login | No        | Log in  |

### Users

| Done               | Method | URI    | authorize | comment     |
| ------------------ | ------ | ------ | --------- | ----------- |
| :white_check_mark: | GET    | /users | Yes       | List Users  |
| :white_check_mark: | POST   | /users | Yes       | Create User |

### Swaggo

```
Generate command: swag init

URL: http://localhost:8000/api/v1/swagger/index.html
```

## FOR DEV

### GIT

Update

```
./git.sh "COMMENT"
```

### Project Structure

```
.

│
├───api
│   │   server.go
│   │
│   ├───controllers
│   │       users.go
│   │
│   ├───infrastructure
│   │       db.go
│   │
│   ├───middlewares
│   │   └───jwt
│   │           auth_jwt.go
│   │
│   ├───models
│   │       base.go
│   │       casbin.go
│   │       role.go
│   │       user.go
│   │
│   ├───routes
│   │       auth.go
│   │       router.go
│   │
│   ├───seeds
│   │       seeder.go
│   │
│   ├───services
│   │       shared.go
│   │       users.go
│   │
│   └───utils
│       ├───messages
│       │       errors.go
│       │       messages.go
│       │
│       └───responses
│               responses.go
│
├───docs
│       docs.go
│       swagger.json
│       swagger.yaml
│
├───rest-client
│       local.http
│
└───tests
│       setup_test.go
│
│   .env
│   .env.example
│   .gitignore
│   auth.conf
│   gin-bin.exe
│   git.sh
│   go.mod
│   go.sum
│   LICENSE
│   main.go
│   README.md
```
