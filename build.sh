#!/bin/bash
# Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

version="1.4.4"
# rem git describe --tags

if [ "$1" == "cloud" ]; then
    echo cloud, version $version
    cp ./server/timesheet-cloud.yaml ./server/timesheet-prod.yaml
else
    echo DataArch, version $version
    cp ./server/timesheet-dataarch.yaml ./server/timesheet-prod.yaml
fi

# ==============================================
echo Removing artifacts from the previous build ...
[ -f "./build/timesheet.exe" ] && rm ./build/timesheet.exe
[ -f "./build/timesheet.app" ] && rm build/timesheet.app
[ -f "./build/timesheet.bin" ] && rm build/timesheet.bin
[ -f "./build/timesheet_arm.bin" ] && rm build/timesheet_arm.bin
[ -f "./build/timesheet.yaml" ] && rm build/timesheet.yaml
[ -f "./build/timesheet-prod.yaml" ] && rm build/timesheet-prod.yaml
[ -f "./build/MS_Windows_64bit.zip" ] && rm build/MS_Windows_64bit.zip
[ -f "./build/Linux_64bit.zip" ] && rm build/Linux_64bit.zip
[ -f "./build/Raspberry_Pi.zip" ] && rm build/Raspberry_Pi.zip
[ -f "./build/Mac_OS_X_64bit.zip" ] && rm build/Mac_OS_X_64bit.zip
[ -f "./build/documentation/documentation.md" ] && rm build/documentation/documentation.md
[ -d "./build/documentation/statics" ] && rm -r ./build/documentation/statics/
[ -f "./build/logs/error.log" ] && rm build/logs/error.log
[ -f "./build/logs/info.log" ] && rm build/logs/info.log
[ -d "./build/client/dist" ] && rm -r  ./build/client/dist
rm ./build/data/*.csv

echo Compiling frontend ...
cd ./client || exit
if [ "$1" == "cloud" ]; then
    yarn run build-cloud
else
    yarn run build-da
fi
cp -r ./dist/spa ./../build/client/dist
cd ..

echo Compiling backend ...
cd ./server || exit
if [ "$1" == "cloud" ]; then
    cp ./timesheet-cloud.yaml ./../build/timesheet.yaml
    cp ./data/consultants_cloud.csv ./../build/data/consultants_cloud.csv
    cp ./data/holidays_us_2019.csv ./../build/data/holidays_us_2019.csv
    cp ./data/projects_cloud.csv ./../build/data/projects_cloud.csv
    cp ./data/rates_cloud.csv ./../build/data/rates_cloud.csv
    cp ./data/reported_records_cloud.csv ./../build/data/reported_records_cloud.csv
else
    cp ./timesheet-dataarch.yaml ./../build/timesheet.yaml
    cp ./data/consultants_dataarch.csv ./../build/data/consultants_dataarch.csv
    cp ./data/holidays_cz_2019.csv ./../build/data/holidays_cz_2019.csv
    cp ./data/projects_dataarch.csv ./../build/data/projects_dataarch.csv
    cp ./data/rates_dataarch.csv ./../build/data/rates_dataarch.csv
    cp ./data/reported_records_dataarch.csv ./../build/data/reported_records_dataarch.csv
fi
cp ./documentation/documentation.md ./../build/documentation/documentation.md
cp -r ./documentation/statics ./../build/documentation/statics/

# echo MS Windows, 64-bit ...
# GOOS=windows
# GOARCH=amd64
# go build -ldflags "-X github.com/valasek/timesheet/server/version.Version=$version" -o ./../build/timesheet.exe ./timesheet.go
echo Linux, 64-bit...
GOOS=linux
GOARCH=amd64
go build -ldflags "-X github.com/valasek/timesheet/server/version.Version=$version" -o ./../build/timesheet.bin ./timesheet.go
# echo Raspberry Pi, ARM 5 ...
# GOOS=linux
# GOARCH=arm
# GOARM=5
# go build -ldflags "-X github.com/valasek/timesheet/server/version.Version=$version -o ./../build/timesheet_arm.bin ./timesheet.go
# echo macOS, 64-bit ...
# GOOS=darwin
# GOARCH=amd64
# go build -ldflags "-X github.com/valasek/timesheet/server/version.Version=$version" -o ./../build/timesheet.app ./timesheet.go
cd ..

# echo =========================
# echo Compressing artifacts ...
# cd ./build
# zip -r MS_Windows_64bit.zip timesheet.exe timesheet.yaml client/ data/ logs/ documentation/
# zip -r Linux_64bit.zip ./timesheet.bin ./timesheet.yaml client/ data/ logs/ documentation/
# zip -r Raspberry_Pi.zip ./timesheet_arm.bin ./timesheet.yaml client/ data/ logs/ documentation/
# zip -r macOS_64bit.zip ./timesheet.app ./timesheet.yaml client/ data/ logs/ documentation/
# cd ..

echo Success, builds are ready in /build folder
