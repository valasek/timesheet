# Timesheet for consultants

Web application to report consulting hours on projects using selected rates on a weekly bases. Supports export into csv.

# Requirements

- Linux, Windows or macOS
- DB connection to PostgreSQL
- Initial yearly data in csv format (dates should be entered in ISO format YYYY-MM-DD HH:MM:SS):
-- consultants, projects, rates, holidays, initial reported records - optional). 

## Todo

- IMPORTANT
  - create new record
  - duplicate record
- overview
  - show in days
  - edit overtime and total working time per week, month
  - compare weekly reported time against nominal total time
- Export plugin to excel and csv

# Fixes

- NEW RECORD
  - IMPORTANT fix wrong from and to dates
  - do not create new record if no consultant is selected
- Edit records
  - do not save the value id ESC is pressed
  - Do not call backend 2x e.g. for updateDescription 

## Improvements business
- Ability to lock last week
- show only available rates per project
- Paginate and sort server-side - using vuetify data table
- Consistency checks
- Add billing evidence
- Export to csv plugin

## Improvements technical

- email confirmation
- letsencrypt tls
- move remembered state from local storage to the backend once user entity available
- administration of a DB - https://getqor.com/en
https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
- refactor routes as on https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
- connect Vue with go - https://juliensalinas.com/en/golang-API-backend-vuejs-SPA-frontend-docker-modern-application/

/* eslint-disable-line no-console */

# Standing on the shoulders of giants

[Go](https://golang.org/), [Vue](https://vuejs.org/), [Vuetify](https://vuetifyjs.com/en/), [PostgreSQL](https://www.postgresql.org/)

## Backend

- Middleware: [Negroni](https://github.com/urfave/negroni)
- Router: [Gorilla](https://github.com/gorilla/mux)
- Orm: [Gorm](https://github.com/jinzhu/gorm) (with [PostgreSQL](https://www.postgresql.org/) persistence)
- Jwt authentication: [jwt-go](https://github.com/dgrijalva/jwt-go) and [go-jwt-middleware](https://github.com/auth0/go-jwt-middleware)
- User management

## Frontend

- [Vue.js](https://vuejs.org/) spa client with webpack
- [Vuetify](https://vuetifyjs.com/en/) - light theme

# To get started

``` bash
# clone repository
go get github.com/valasek/timesheet
cd $GOPATH/src/github.com/valasek/timesheet

# install Go dependencies (and make sure ports 3000/8080 are open)
go get -u ./... 
go run timesheet.go

# open a new terminal and change to the client dir
cd client

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run serve

# build for production with minification
npm run build
```

# License

[MIT license](./LICENSE.md) - see LICENSE for more details