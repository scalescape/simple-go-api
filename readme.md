# Simple API
Simple REST API with database (postgres) integration

## HighLevel Agenda
* Integrating with postgres (few concepts)
* Live code walkthrough

### Detailed
Check [screencast video](https://youtu.be/jzfTSNrHPjk) for more explanation which covers
* Context on base http ping code
* DB config
* DB connections / connection pool
* Migrations
    * creating basic one
    * running it along with deployment
* Query
    * Transaction
    * Using context

### Gotchas

* connpool != threads
    * 5-30 would suffice for 10k/s transactions responding in few ms
* ssl/md5 auth for connecting with DB
* never overwrite/edit older migration
* rollback should be tested
* [migrate installation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* ensure env var sourcing is local to command
* Use transaction when you've more than one modification

# Dependencies
 * [golang-migrate](https://github.com/golang-migrate/migrate)
 * [jmoiron/sqlx](https://github.com/jmoiron/sqlx)

## Running
* create env file `.env` with following config and modify as necessary

```
export HOST="localhost"
export PORT=8080

export DB_HOST="localhost"
export DB_USER="postgres"
export DB_PASSWORD="<db_password>"
export DB_NAME="simple_api"
export DB_SSL_MODE="disable"
export DB_PORT="5432"
```

* Ensure you've postgres up and running and run the following to setup the DB
 
```
make db.create db.migrate
```

* Run the service by running the following command

```
source local.env && make run
```

## Disclaimer

This code was drafted quickly for a meetup session, and **this is not production ready**.
You could use this as a base but still need to improvise.
