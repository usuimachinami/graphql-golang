# graphQl Project

## setup

Docker containers up

```
cd /path/to/project/env

docker-compose up -d
```

Install dependent packages in the golang container.

```
go get github.com/labstack/echo
go get github.com/dgrijalva/jwt-go
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
go get github.com/go-sql-driver/mysql
go get github.com/graphql-go/graphql
```

Run app.

```
go run src/app/main.go
```