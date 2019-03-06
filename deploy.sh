#!/bin/bash

echo cleaning up previous installation ...
pkill timesheet.bin
rm /opt/timesheet/client/dist/css/*
rm /opt/timesheet/client/dist/js/*
rm /opt/timesheet/data/*.csv
rm /opt/timesheet/logs/*.log
rm /opt/timesheet/timesheet.bin

echo deploing new version ...
cp -r ./Pictures/* /opt/timesheet/
chmod +x /opt/timesheet/timesheet.bin
sed -i -e 's/localhost/your.domain.com/g' /opt/timesheet/client/dist/js/app.*.js
sed -i -e 's/localhost/your.domain.com/g' /opt/timesheet/client/dist/js/app.*.js.map

echo initializing and starting timesheet ...
cd /opt/timesheet
./timesheet.bin db --clean
./timesheet.bin db --load all
nohup ./timesheet.bin server

echo new version deployed
