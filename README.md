go mod init github.com/dduafa/go-server

Make a copy of .env.example

- Start postgres container
 ```docker-compose up -d --remove-orphans```

- Stop postgres container
 ```docker-compose down```

- Access docker bash
```docker exec -it postgres bash```

<!-- Connect to docker postgre and check extensions -->
psql -U postgres golang-gorm
select * from pg_available_extensions;

<!-- Install this extention (because we are using uuid_generate_v4() for or model IDs) -->
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


<!-- Used to load and validate env variables -->
go get github.com/spf13/viper

<!-- cors -->
go get github.com/gin-contrib/cors

<!-- GORM library and the Postgres driver -->
go get -u gorm.io/gorm  
go get gorm.io/driver/postgres

<!-- Run migration -->
go run migrate/migrate.go

<!-- jwt -->
github.com/golang-jwt/jwt


<!-- Create server with Gin  -->
go get github.com/gin-gonic/gin 
go install github.com/cosmtrek/air@latest -->> will help us to hot-reload the Gin server
