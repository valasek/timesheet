# Simple working time evidence with export plugins
Web application

# Requirements
PostgreSQL DB VERSION

## Todo
- Edit item
- update week when month is changed
- lock last week
- overview
-- edit overtime and total working time per week, month
-- compare weekly reported time against nominanl total time
- Remember my settings (which month, week, consultant)
- Export plugin to excel and csv

## Improvements business
- consistency checks
- initial set up
-- add environment configuration - for cors on backend and frontend)
-- seed DB, all and per table
- Add billing evidence and export plugin

## Improvements technical
- read https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
- refactor routes as on https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
- connect Vue with go - https://juliensalinas.com/en/golang-API-backend-vuejs-SPA-frontend-docker-modern-application/

/* eslint-disable-line no-console */

### Features:
- Middleware: [Negroni](https://github.com/urfave/negroni)

- Router: [Gorilla](https://github.com/gorilla/mux)

- Orm: [Gorm](https://github.com/jinzhu/gorm) (sqlite or postgres)

- Jwt authentication: [jwt-go](https://github.com/dgrijalva/jwt-go) and [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)

- [Vue.js](https://vuejs.org/) spa client with webpack

- User management

### TODO:
- config from file

- email confirmation

- logrus

- letsencrypt tls

### To get started:

``` bash
# clone repository
go get github.com/valasek/time-sheet
cd $GOPATH/src/github.com/valasek/time-sheet

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run server.go

# open a new terminal and change to the client dir
cd client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build
```

### License

MIT License  - see LICENSE for more details