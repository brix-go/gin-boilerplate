package inrastructure

import (
	"GinBoilerplate/config"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"sync"
)

type LogCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
}

var instance *LogCustom
var once sync.Once

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom(configServer config.ConfigModel) *LogCustom {
	var log *logrus.Logger

	configElstc := configServer.Log

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	client, err := elastic.NewClient(elastic.SetURL(
		fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(configElstc.Username, configElstc.Password))

	if err != nil {
		selfLogError(err, "config/log: elastic client", log)
	} else {
		hook, err := elogrus.NewAsyncElasticHook(
			client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
		if err != nil {
			selfLogError(err, "config/log: elastic client sync with logrus", log)
		}
		log.Hooks.Add(hook)
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: configServer.AppConfig.Name,
				Host: configServer.AppConfig.Host,
				Port: configServer.AppConfig.Port,
			},
		}
	})

	return instance
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Error(data LogData) {
	data.level = LEVEL_ERROR
	errorCause := ""
	errorString := ""
	err := data.Err

	data.packageName, data.functionName = getPackageAndFuncName()
	errorCause, errorString = getErrorStack(err)
	timeMs, timeFmt := responseTimeString(data.StartTime)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":                   l.WhoAmI,
		"error_cause":              errorCause,
		"error_message":            errorString,
		"package_name":             data.packageName,
		"function_name":            data.functionName,
		"trace_header":             data.TraceHeader,
		"level":                    data.level,
		"request":                  data.Request,
		"response":                 data.Response,
		"request_backend":          data.RequestBackend,
		"response_backend":         data.ResponseBackend,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Error(data.Description)
}

func (l *LogCustom) Success(data LogData) {
	data.level = LEVEL_SUCCESS

	timeMs, timeFmt := responseTimeString(data.StartTime)

	data.packageName, data.functionName = getPackageAndFuncName()

	l.Logrus.WithFields(logrus.Fields{
		"whoami":                   l.WhoAmI,
		"package_name":             data.packageName,
		"function_name":            data.functionName,
		"trace_header":             data.TraceHeader,
		"level":                    data.level,
		"request":                  data.Request,
		"response":                 data.Response,
		"request_backend":          data.RequestBackend,
		"response_backend":         data.ResponseBackend,
		"elapsed_time_millisecond": timeMs,
		"elapsed_time_format":      timeFmt,
	}).Info(data.Description)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Fatal(err error, description string, traceHeader map[string]string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogFatal(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}
