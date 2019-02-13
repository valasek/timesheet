#!/bin/sh
# ls -la /
# ls -la /dist/
./timesheet.bin -v
./timesheet.bin db --clean
./timesheet.bin db --load all
./timesheet.bin server