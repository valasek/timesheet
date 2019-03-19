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

# Documentation

Read full [documentation](./server/documentation/documentation.md).

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

# Contributing

I would love your help! Before submitting a PR, please read over the [Contributing](CONTRIBUTING.md) guide.

Here's a couple of areas that could use some love:

* **Import/Export plugins** - do you like functionality but want to integrate with your company billing SW? Create an issue and if you can submit PR.
* **Javascripts** - I'm no JS wizz, I'm sure the Javascript code in ui/src could be improved/simplified/tested.  
* **Documentation** - Typo? Does something not make sense? Could it be worded better? Please help!
* **Test Coverage** - Would love to get coverage over 80%, fontend and backend.

# Standing on the shoulders of giants

[Go](https://golang.org/), [Gin web framework](https://github.com/gin-gonic), [Vue](https://vuejs.org/), [Vuetify](https://vuetifyjs.com/en/), [PostgreSQL](https://www.postgresql.org/) and [MySQL](https://www.mysql.com/)

## Go Backend

- [Gin web framework](https://github.com/gin-gonic)
- [Gorm](https://github.com/jinzhu/gorm) (with [PostgreSQL](https://www.postgresql.org/) or [MySQL](https://www.mysql.com/) persistence)
- [Logrus](https://github.com/sirupsen/logrus), [Cobra](https://github.com/spf13/cobra), [Viper](https://github.com/spf13/viper), [Now](https://github.com/jinzhu/now), [lumberjackrus](https://github.com/orandin/lumberjackrus), [Cron](https://github.com/robfig/cron)

## JS Frontend

- [Vue.js](https://vuejs.org/) spa client with webpack
- [Vuetify](https://vuetifyjs.com/en/) - light theme
- [Axios](https://github.com/axios/axios), [date-fns](https://date-fns.org/)
