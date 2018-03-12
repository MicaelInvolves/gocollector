[![Build Status](https://travis-ci.org/gesiel/gocollector.svg?branch=master)](https://travis-ci.org/gesiel/gocollector)
[![Coverage Status](https://coveralls.io/repos/github/gesiel/gocollector/badge.svg)](https://coveralls.io/github/gesiel/gocollector)

# Go Collector

A Go (Rest API) + Vue (Web Client) experiment.

## This project uses
- [Clean Architecture](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Go](https://golang.org/)
- [GoVendor](https://github.com/kardianos/govendor)
- [Echo](https://echo.labstack.com/)
- [Mgo](https://labix.org/mgo)
- [Vue](https://vuejs.org/)
- [Axios](https://github.com/axios/axios)
- [Bulma](https://bulma.io/)
- and more...

## How it works

This project collects users web surfing data from web sites configured with [its js client](https://github.com/gesiel/gocollector/tree/master/jsclient). A Go Rest API receives all the user navigation data and a Vue web client shows all collected data.

### Packages
1. Both `access` and `subscriber` are domain specific logic. These packages depend only on abstractions. All dependencies are injected with DI and IC.
2. Package `controllers` stands for the web api classes.
3. `databases` stores all database gateways implementations. It's basic Mongo queries.
4. All Go dependencies are in the `vendor` folder.
5. The `webclient` folder stores the Vue web client. Running `npm run build` in this folder updates the `static` folder content with a minified version of the web client.
5. The JS lib stay on the `jsclient`folder. It collects and sends information to the backend.

## Running

**Atentition: You need a Mongo DB up and running to proceed.**

1. Run the tests with 
```
go test ./...
```
2. Run the app with
``` 
PORT=8080 MONGODB_URI=localhost go run main.go
```
3. Open your browser on `http://localhost:8080/` to list all collected data. If was no collected data, surf a little bit on `http://localhost:8080/examples` and fill up the contact page.

A version of this tool is up and running on heroku. Just visit `https://gocollector.herokuapp.com/examples` and `https://gocollector.herokuapp.com/`.

To run the JS lib tests, see [here](https://github.com/gesiel/gocollector/tree/master/jsclient).
