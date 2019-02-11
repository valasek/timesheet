// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package main

import (
	"github.com/sirupsen/logrus"
	"github.com/valasek/timesheet/server/cmd"
	"github.com/valasek/timesheet/server/logger"
)

func main() {
	logger.Log = logrus.New()
	cmd.Execute()
}
