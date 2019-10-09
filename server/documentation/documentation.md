# Installation

Timesheet can be installed using pre-build docker images or as a binary.

## Requirements

Application can run in [Docker containers](https://www.docker.com/resources/what-container). Server image size 31 MB, DB image size 320 MB.

If you are hosting it, supported platforms are:
- Linux, MS Windows, macOS, Raspberry Pi
- PostgreSQL DB

# Quick start

1. Update Application settings, Vacation settings and Warning limits to meet your needs. See Configuration chapter.
2. Use Backup & Restore to export demo data, update csv files and import back your projects, rates, consultants and holidays. See Backup & Restore chapter.

Enjoy. You are good to go ...

# Administration

Is accessible on Administration page:

![Administration](./statics/administration.png?raw=true "Administration")

# Configuration

Below is default and commented `timesheet.yaml` configuration file shipped with a product. 

```
### Default configuration file

######################
# Reporting settings #

dailyWorkingHours: 8 # Used to calculate weekly and monthly expected working hours, can be changed in UI
dailyWorkingHoursMin: 8 # Used to highlight if reported less, can be changed in UI
dailyWorkingHoursMax: 12 # Used to highlight if reported more, can be changed in UI

# Rate used for vacations
vacation: "Vacation"
yearlyVacationDays: 20 # Used to calculate weekly and monthly expected working hours, can be changed in UI

# Rate for additional vacations. If not used, leave blank "" and set yearlyPersonalDays: 0, can be changed in UI
vacationPersonal: "Personal Day"
yearlyPersonalDays: 3 # Used to calculate weekly and monthly expected working hours, can be changed in UI

# Rate used for additional vacation intended for sick day. If not used, leave blank "" and set yearlySickDays: 0, can be changed in UI
vacationSick: "Sick Day"
yearlySickDays: 2 # Used to calculate weekly and monthly expected working hours, can be changed in UI

# Categorize all rates into one of these types used on Reported Overview page
isWorking: "work" # when consultant works, can be changed in UI
isNonWorking: "not-work" # when consultant dows not work, examples: vacation, sick, personal day, public holiday, vacation, unpaid leave, ..., can be changed in UI

########################
# Application settings #
SSL: false # true/false, if server has SSL certificate set to true to use HTTPS, false = HTTP
GIN_MODE: "release" # "debug" or "release" - switch server app mode
url: "timesheet.simplesw.net" # URL on which application is running
# PORT: "443"                   # port on which application is running
PORT: "8080" # port on which application is running

# DB type
dbType: "postgresql" # allowed types postgresql or mysql

# Production URL - will be read from production environment config variable
# If set, Database settings section variables will be ignored
DATABASE_URL: ""

# Log folder, relative to timesheet folder
logFolder: "logs"

# # Folder for uploaded data, relative to timesheet folder
uploadFolder: "data/uploaded"
uploadFolderTemp: "data/uploaded/temp"

# csv data files which are loaded via command "timesheet db --load all"
data:
  consultants: "consultants_cloud.csv"
  rates: "rates_cloud.csv"
  projects: "projects_cloud.csv"
  reportedRecords: "reported_records_cloud.csv"
  holidays: "holidays_us_2019.csv"

export:
  location: "data/exported" # select an empty and an existing folder

#####################
# Database settings #

# DB backup settings - backup data can be imported directly by a command "timesheet db --load all"
backup:
  location: "data/backups" # select an empty and an existing folder relative to timesheet/data folder
  rotation: 14 # how many backups back will be kept
  interval: "daily" # daily or weekly - how often the DB backup should be done

# DB credentials
# used for development and testing. Ignored if DATABASE_URL is set
postgresql:
  # host: "db" #
  host: "127.0.0.1" #
  port: "5432"
  name: "timesheet"
  user: "timesheet"
  password: "timesheet"
  sslMode: "disable"

mysql:
  user: "timesheet"
  password: "timesheet"
  dbname: "timesheet"
```

## Command Line Options

`.\timesheet.exe` or `./timesheet.bin` or `./timesheet.app`
```
Web based timesheet application with DB persistence.
        
Application reads DB and server configuration from timesheet.toml, loads default data if DB is empty and launch web GUI.      

Usage:
  timesheet [command]

Available Commands:
  db          Initiate, load. backup DB og generate demo data. See timesheet help db
  help        Help about any command
  routes      Prints the list of all available routes
  server      Starts the server on URL and port defined in config.yaml

Flags:
      --config string   config file (default is ./timesheet.yaml)
  -h, --help            help for timesheet
  -v, --version         Prints application versions

Use "timesheet [command] --help" for more information about a command.
```

`.\timesheet.exe help db` or `./timesheet.bin  help db` or `./timesheet.app  help db`
```
Initiate, load, backup DB or generate demo data.

Command first tests connection to DB. If succeeds it will initiate, load, backup db or generate demo data and exit.

Usage:
  timesheet db [flags]

Flags:
  -b, --backup        Backup all DB tables in the format used by db --load command
  -c, --clean         Drop and create all required DB tables
  -g, --generate      Generate demo data and save them into ./data folder
  -h, --help          help for db
  -l, --load string   Truncate DB table/tables and load initial data from files in folder ./data. Options:
                      all - load all tables
                      rates | consultants | projects | holidays | reported_records - load selected table

Global Flags:
      --config string   config file (default is ./timesheet.yaml)
```

# API

To retrieve reported records per selected period use HTTP GET request on ```<domain>/api/reported/from/<from>/to/<to>``` where
* \<domain\> - is your domain name
* \<from\>, \<to\> - dates in format YYYY-MM-DD

Example: ```http://localhost:3000/api/reported/from/2019-04-13/to/2019-04-30```

# Backup & Restore

All data can be downloaded locally in CSV format, encoded in UTF-8.
CSV files can be modified and imported back. Import will delete and replace all existing data.

![Backup & Restore](./statics/backup_restore.png?raw=true "Backup & Restore")

Check log files using Administration / Logs. Error log should not contain any errors or warnings.

## Description of Data Files

When editing CSV files, save the file using UTF-8 encoding. Compress all into one zip archive before uploading.

### consultants.csv

Contains all consultants which can report hours.
```
created_at,name,allocation,disabled
"YYYY-MM-DD HH:MM:SS","consultant name","allocation between 0 - 1",true/false
"2019-01-14 00:00:00","Evan You","1",false
```

### holidays.csv

Contains state holidays per selected year.
```
created_at,date,description
"YYYY-MM-DD HH:MM:SS","name of the holiday"
"2019-01-01 00:00:00","2019-01-01","New Year's Day"
```

### rates.csv

Contains rates which can be selected by any consultant.
Rates are categorized into working *work* and *non-working*. Category names can be changed but update also timesheet.yaml.
```
created_at,name,type
"YYYY-MM-DD HH:MM:SS","rate name","work or non-work"
"2019-01-14 00:00:00","Standard","work"
"2019-01-14 00:00:00","Vacation","not-work"
```

### projects.csv

Contains projects on which consultants can report the work and default project rate.
```
created_at,name,rate,disabled
"YYYY-MM-DD HH:MM:SS","project name","rate name",true/false
"2019-01-14 00:00:00","Vue","Standard",false
```

### reported_records.csv

Contains all reported hours. Hours (N.N) can be a decimal number between 0 and 24. Project, rate and consultant are names of existing records from corresponding csv files.
```
created_at,date,hours,project,description,rate,consultant
YYYY-MM-DD HH:MM:SS,YYYY-MM-DD,N.N,project name,description of the work,rate name,project name
2019-01-10 10:00:00,2019-01-01,6,Vue,Updates of all Vue.js documentation examples using typescript,Off-site,Evan You
```

# Upgrade

Follow these steps to upgrade:
* Export your data as described in Backup & Restore chapter
* Stop running timesheet binary, replace `timesheet` folder with a new version and start new timesheet binary. Optionally configure and use one click `./deploy.sh` upgrade script.
* Import saved data as described in Backup & Restore chapter

# License

Free for education and non-commercial usage. Pay for the commercial usage of the application to support further development and maintenance via
<a href="https://paypal.me/StanislavValasek">PayPal</a>.

Currently I am working on a Pro Version of Timesheet for enterprise. Along with support, some of the planned features include:

* User management/login
* Permissions
* HTTPS
* Plugin for import/export 
* Reporting metrics
* Cloud version

If you or your organization would like to help beta test a Pro version of Timesheet, please get in touch with us via email: [timesheet.simplesw@gmail.com](mailto:timesheet.simplesw@gmail.com)

# Release Notes

## Version 1.4.4
Released on October 10, 2019

### New Features
* Administration - ability to add, hide/show and delete consultants and projects from UI
* Show managed data statistics on Administration and home page
* Server can generate demo data via the command line

### Usability
* Table of Contents in the documentation
* Clean up and initialize DB when deployed

### Technical
* Use Go 1.13.1 and Quasar 1.1
* Docker compose, dockerfiles and configuration for dev and production versions

## Version 1.4.3
Released on July 23, 2019

### Usability
* Styling and logo
* New breadcrumbs
* ESC on search deletes the text, prepend delete search text icon
* Overview page elements rearranged

### Fixes
* Fixed remaining weekly and monthly total remaining time

### Technical
* Use Go 1.12.7 and Quasar 1.0.5

## Version 1.4.2
Released on July 03, 2019

### Usability
* Table on previous weeks contains more space

### Fixes
* Consultants, projects, and rates selects are sorted alphabetically
* Home page graph - added filter by year
* Available working time computation fixed
* Administration / upload - parametrized url and port, uploaded file name is parsed on the fly

### Technical
* Upgrade to Quasar 1.0.2, Go 1.12.6, npm replaced with yarn

## Version 1.4.0
Released on June 24, 2019

### Features
* API to retrieve reported records in selected period
* Home page listing top 10 projects
* Footer showing app version and server name

### Usability
* Material Design 2.0
* Documentation contains API description, mentions UTF-8 for exported files
* Week unlock button moved on Reporting page
* Breadcrumbs added

### Fixes
* Axios security bug
* Overview - table pagination added, shows all records
* Record dates saved in UTC

### Technical
* UI migrated from Vuetify to Quasar framework
* Refactor component names

## Version 1.2.3
Released on April 24, 2019

### Usability
* Help page contains links to new feature and bug report forms
* Duplicate record will create new record on the next day
* Daily hours shown red in weekly overview if was reported more than 24 hours per day
* Allow to report with a warning more than 24 hours per day

### Technical:
* Update to Vuetify 1.5.13, GO 1.12.4 and latest 3rd party GO packages
* README cleanup and code refactoring

## Version 1.2.2
Released on March 30, 2019

### Usability
* Add release 1.2.1 in docs
* Import/Export text clarification 

### Fixes
* Consultant cannot report > 24 hours 

### Technical
* HTTPS support
* Update to Vuetify 1.5.8 

## Version 1.2.1
Released on March 20, 2019

### New Features
* MySQL support

### Usability
* Dark / Light mode theme switch

### Fixes
* DB not available - log human error instead of an exception
* Documentation typos

### Technical Changes
* Upgrade to Vuetify 1.5.7
* Build with Go 1.12.1
* Show server name and port on About page

## Version 1.2.0
Released on March 14, 2019

### Usability

* Documentation created
* Scrollable data table
* Faster responses on user interface
* Show message if log file is empty

### Fixes

* State holidays export contains dates

### Technical

* Replaced backend GO negroni and gorilla mux with gin
* Replaced moments.js with date-fns
* Update to Vuetify 1.5.6
* Update Travis configuration on Go 1.12
* Added deploy shell script

## Version 1.1.1
Released on March 7, 2019

### Usability

* Previous weeks are loaded faster
* Weeks and consultants can be changed directly on the Overview page
* Removed footer, main toolbar takes less space
* Disable all fields in previous weeks unless editing is enabled
* Warn if entered date is not from this week
* Validation on reported hours, allowed number: 0 - 24
* Validate maximum reported daily hours during record duplication
* Validations on entering hours and days

### Fixes

* Backup archive contains exported csv files without subfolders

## Version 1.0.0
Released on December 2018

The first publicly released version
