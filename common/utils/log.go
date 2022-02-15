package utils

import (
	"fmt"
	"path"
	"runtime"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

var Logger *logrus.Logger

func LoggerInit() {
	Logger = logrus.New()
	var LogFormatter = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
		Formatter: nested.Formatter{
			//HideKeys: true,
			TimestampFormat: "2006-01-02 15:04:05",
			FieldsOrder:     []string{"component", "category"},
		},
	}

	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetReportCaller(true)

	if LogFormatter != nil {
		Logger.SetFormatter(LogFormatter)
	}
}
