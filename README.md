[![GitHub release](https://img.shields.io/github/release-pre/valasek/timesheet.svg)](https://github.com/valasek/timesheet/releases)
[![GitHub issues](https://img.shields.io/github/issues/valasek/timesheet.svg)](https://github.com/valasek/timesheet/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/valasek/timesheet)](https://goreportcard.com/report/github.com/valasek/timesheet)


| **Linux & Mac & Windows** |
| :-----------------------: |
| [![Build Status](https://travis-ci.org/valasek/timesheet.svg?branch=master)](https://travis-ci.org/valasek/timesheet) |

# Simple Timesheet

Self-hosted web application for weekly reporting. Report your consulting hours on projects using selected rates.

![Screencast](screenshots/screencast.gif?raw=true "Screencast")

# Supporting Simple Timesheet

Simple Timesheet is project with its ongoing development made possible entirely by the support of these awesome backers. If you'd like to join them, please consider:

- [Become a backer or sponsor on Patreon](https://www.patreon.com/valasek)
- [One-time donation via PayPal](https://paypal.me/StanislavValasek)

# Rationale

Reporting and billing process includes three separated steps with well-defined data which flows in between:
* **Time reporting** - all consultants
  * covered by this app - data are safe, eliminates manual errors, shows important data for consultants, data entry is as easy as possible
* **Billing preparation** (confirming/editing reported hours for billing) - project manager or administrative person
  * covered partially by this app. If you bill 1:1 with your reporting then it has all you need
  * data are available in the DB which can be exported by this app read or exported by any DB tool
* **Exporting reported and billed hours** in the internal accounting system for the invoicing to the clients and consultants - salaries
  * not covered by this app
  * data are available in the DB which can be exported by this app read or exported by any DB tool

# Requirements

- Linux, Windows or MacOS
- PostgreSQL DB
- Initialize the DB via `timesheet db --clean`
- Import data from csv via `timesheet db --load all`
  - mandatory: consultants, projects, rates, holidays
  - optional: reported_records
- start the server `timesheet server`

Simple timesheet can be also deployed using Docker (server image size 24.9 MB, DB image size 312 MB)

# How to start

`docker-compose -f "docker-compose.yml" up -d --build`
- Update Application settings , Vacation settings and Warning limits to meet your needs
- Use Backup & Restore to export demo data. Update csv files and import back your projects, rates, consultants and holidays.
- You are good to go 


# Contributing

I would love your help! Before submitting a PR, please read over the [Contributing](CONTRIBUTING.md) guide.

Here's a couple of areas that could use some love:

* **Import/Export plugins** - do you like functionality but want to integrate with your company billing SW? Create an issue and if you can submit PR.
* **Javascripts** - I'm no JS wizz, I'm sure the Javascript code in ui/src could be improved/simplified/tested.  
* **Documentation** - Typo? Does something not make sense? Could it be worded better? Please help!
* **Test Coverage** - Would love to get coverage over 80%, fontend and backend.


# Usage

Free for education and non-commertial usage.

Pay for the commertial usage of the application to support further development and maintenance via
<a href="https://www.patreon.com/valasek">Patreon</a> or <a href="https://paypal.me/StanislavValasek">PayPal</a>.

My plan is to soon start working on a Pro Version of Timesheet for enterprise. Along with support, some of the planned features include:

* User management/login
* Permissions
* HTTPS
* Plugin for import/export 
* Reporting metrics
* Cloud version

If you or your organization would like to help beta test a Pro version of Timesheet, please get in touch with me:

    Twitter: @valasek
    Email: valasek at gmail.com

# Standing on the shoulders of giants

[Go](https://golang.org/), [Gin web framework](https://github.com/gin-gonic), [Vue](https://vuejs.org/), [Vuetify](https://vuetifyjs.com/en/), [PostgreSQL](https://www.postgresql.org/)

## Go Backend

- [Gin web framework](https://github.com/gin-gonic)
- [Gorm](https://github.com/jinzhu/gorm) (with [PostgreSQL](https://www.postgresql.org/) persistence)
- [Logrus](https://github.com/sirupsen/logrus), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), [Now](https://github.com/jinzhu/now), [lumberjackrus](https://github.com/orandin/lumberjackrus), [Cron](https://github.com/robfig/cron)

## JS Frontend

- [Vue.js](https://vuejs.org/) spa client with webpack
- [Vuetify](https://vuetifyjs.com/en/) - light theme
- [Axios](https://github.com/axios/axios), [moment.js](http://momentjs.com/), [date-fns](https://date-fns.org/)
