#!/bin/sh
# Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

./timesheet.bin -v
./timesheet.bin db --clean
./timesheet.bin db --load all
./timesheet.bin server