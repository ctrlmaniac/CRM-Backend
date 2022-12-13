# CRM-Backend

a simple CRM backend written in Go for the Udacity Golang course

## RUN

To run this application you have to type into your terminal `go run main/main.go`

## BUILD

To build this application run in your terminal `go build main/main.go`

## Project directory structure explained

In this repository you will find different directories, each of them has a purpose

**controllers** holds all the business logic, for example all methods to create, get, update, delete a customer for the fake database.

**main** holds the main module of this application.

**persistance** holds the logic to access the fake-database and the initial values of the fake-database.

**static** holds all html files that are needed, in this case index.html.
