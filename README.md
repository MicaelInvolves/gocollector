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

## Running

**Atentition: You need a Mongo DB up and running to proceed.**

Edit `.env` with your mongo url and app port. Then, execute `go run main.go`. You can open your browser on `http://localhost:PORT/` to list all collected data. If was no collected data, surf a little bit on `http://localhost:PORT/examples` and fill up the contact page.

A version of this tool is up and running on heroku. Just visit `https://gocollector.herokuapp.com/examples` and `https://gocollector.herokuapp.com/`.