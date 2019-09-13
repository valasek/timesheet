#!/bin/bash

# CHANGE HERE - set based on your environment
# DOMAIN=your.domain.com
SOURCE_FOLDER=./timesheet
TARGET_FOLDER=/opt/www/timesheet.simplesw.net-app

# DO NOT CHANGE - deploy script
echo cleaning up previous installation ...
pkill timesheet.bin
rm $TARGET_FOLDER/client/dist/css/*
rm $TARGET_FOLDER/client/dist/js/*
rm $TARGET_FOLDER/data/*.csv
rm $TARGET_FOLDER/logs/*.log
rm $TARGET_FOLDER/timesheet.bin

echo deploing new version ...
cp -r $SOURCE_FOLDER/* $TARGET_FOLDER/
chmod +x $TARGET_FOLDER/timesheet.bin
# sed -i -e 's/localhost/$DOMAIN/g' $TARGET_FOLDER/client/dist/js/app.*.js
# sed -i -e 's/localhost/$DOMAIN/g' $TARGET_FOLDER/client/dist/js/app.*.js.map

echo initializing and starting timesheet ...
cd $TARGET_FOLDER
./timesheet.bin db --clean
./timesheet.bin db --load all
nohup ./timesheet.bin server

echo new version deployed
