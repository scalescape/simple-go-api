# Me
Dineshkumar
Senior software engineer @ elasticsearch
@devdineshkumar

# Go - Database Integration

## HighLevel Agenda
* Integrating with postgres (few concepts)
* Live code walkthrough

## Detailed
* Context on base http ping code
* We'll use postgres
    * DB config
* DB connections / connection pool
* Migrations
    * create
    * running it along with deployment
* Query
    * Using context
    * Transaction
* Gotchas
    * connpool != threads
    * ssl/md5 auth
    * never overwrite migration
    * rollback should be tested
    * [migrate installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
    * env var sourcing

## Points


# Libs:
 * [golang-migrate](https://github.com/golang-migrate/migrate)
 * [jmoiron/sqlx](https://github.com/jmoiron/sqlx)
 * [devdinu/simple-go-api](https://github.com/devdinu/simple-go-api)

# Questions:
 
