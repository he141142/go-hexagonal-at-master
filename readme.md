
# Hexagonal architecture at master

Overview:
![](https://user-images.githubusercontent.com/60600430/216277035-df882207-5fe8-4e04-a049-e5217bd42865.jpg)

![](https://user-images.githubusercontent.com/60600430/216277022-27e81a45-f1cf-41a4-b85e-db5f5fdffe02.jpg)


install:
```
go mod tidy
```

setup entgo:

```shell
go run -mod=mod entgo.io/ent/cmd/ent generate  ./internal/core/adapters/repo/sql_type/sql/ent/lib/ent
```
run the application:
```shell
go run cmd/main/main.go
```

env sample: 
```dotenv
DB_HOST=localhost
DB_PORT=9001
DB_DRIVER=POSTGRES
DB_USER=sykros
DB_PASSWORD=fqQ3nN4L
#DB_NAME=zlp-demo
DB_NAME=ZLP2
DB_SSL=disab
```

Note that `env` file name should be `dev.env`