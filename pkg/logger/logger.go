package logger

import (
	"os"
  "github.com/charmbracelet/log"
)

var Logger = log.NewWithOptions(os.Stdout, log.Options{
  ReportCaller: true,
  Level: log.DebugLevel,
})
