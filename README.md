# Golang API

A simple JSON API written in Go. I might expand this *a lot* in the future depending how much I feel like learning Go and Backend Development together.

## Gin & Gorm
The backend is powered by Gin Web Framework and GORM, Object-Relational Mapping library for Go.

### Migration
In order to create the database for the application, run
```
go run migrate/migrate.go
```
This is also going to be important when implementing **CI/CD** for the application.

### Environment Variables
The application currently only needs two Environment Variables, however, the `PORT` variable is optional, since Gin runs at port 8080 by default.
<br>
The only mandatory variable is `DB_STRING`, which requires these fields:
```
DB_STRING="host=<host> user=<user> password=<password> dbname=<dbname> port=<port> sslmode=require"
```

Initialising the database still happens the same way through `migrate.go`, but you have to first edit .env to include the hostname and credentials to connect to the database.
<br><br>
You can access your database easily by running the docker command below. Alternatively, you can connect to the database with a database management tool like [pgAdmin](https://www.pgadmin.org/), but I personally like to stay in the terminal.
```
docker exec -it <container_name> psql -U <password> -d <database_name>
```
