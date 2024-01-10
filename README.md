# Golang Grove

A simple JSON API written in Go. I might expand this *a lot* in the future depending how much I feel like learning Go and Backend Development together.

## CI/CD Status
[![Go Build](https://circleci.com/gh/santtuniskanen/golang-grove.svg?style=svg)](https://app.circleci.com/pipelines/github/santtuniskanen/golang-grove)
At this moment, the pipeline only contains a Build check to make sure the application builds properly.

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