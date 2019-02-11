// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package main

import (
	"github.com/valasek/timesheet/server/cmd"
	"github.com/valasek/timesheet/server/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.Log = logrus.New()
	cmd.Execute()
}
