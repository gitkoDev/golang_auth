# Medods Authentication Service

JWT Golang Authentication service built with `Clean Architecure` and `Dependency Injection` principles

## Endpoints

- 0.0.0.0:8080/v1/health &ensp; `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Check API connection

  
- 0.0.0.0:8080/v1/auth?id=*{user id}* &ensp; `=>`  &ensp; **GET** &ensp;  `=>` &ensp; Get token pair

`Response example:`

```json
{
    "access_token": "access_token",
    "refresh_token": "refresh_token",
    "user_ip": "172.24.0.1"
}
```
- 0.0.0.0:8080/v1/auth/refresh &ensp; `=>`  &ensp; **POST** &ensp;  `=>` &ensp; Refresh token pair
  
`Response example:`
```json
{
    "access_token": "access_token",
    "refresh_token": "refresh_token",
}
```

## Running the app

```bash
# start the app
make run
```


## Tools used

- `Logging` &nbsp; **=>**  &nbsp; [Logrus](https://github.com/sirupsen/logrus)
- `Routing` &nbsp; **=>**  &nbsp;  [Gin](https://github.com/gin-gonic/gin)
- `Database` &nbsp; **=>**  &nbsp;  Postgres + [pgx](github.com/jackc/pgx/v5)
- `Database migrations` &nbsp; **=>**  &nbsp; [Goose](https://github.com/pressly/goose#sql-migrations)
- `Containerization` &nbsp; **=>**  &nbsp; [Docker](http://docker.com/) + Docker Compose
- `Authentification and middleware` &nbsp; **=>**  &nbsp;  [JWT Go](https://github.com/golang-jwt/jwt)
