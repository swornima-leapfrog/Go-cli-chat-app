# CLI Chat App

This is a cli chat application where logged in users can chat with each other.

# Installation

## For database

- To install the postgres, run `docker-compose up`
- Then, run a command `cp .env.example .env` to create `.env` file
- Then setup the environment.
- Run migration by using `goose postgres "user=<username> password=<password> dbname=<database name> host=<host> port=<port>" up` command

## Run application

- `go run main.go`
