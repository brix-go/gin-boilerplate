package inrastructure

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const LEVEL_SUCCESS = "success"
const LEVEL_ERROR = "error"
const LEVEL_INFO = "info"
const LEVEL_FATAL = "fatal"

type ConfigLog struct {
	Name string `env:"NAME_SERVER"`
	Port string `env:"PORT_SERVER"`
	Host string `env:"HOST_SERVER"`

	HookElasicEnabled bool `env:"HOOK_ELASTIC_ENABLED"`
	ElasticConfig     ElasticConfig
}

type ElasticConfig struct {
	Host     string `env:"HOST_ELASTICSEARCH,required"`
	Port     string `env:"PORT_ELASTICSEARCH,required"`
	User     string `env:"USER_ELASTICSEARCH"`
	Password string `env:"PASS_ELASTICSEARCH"`
	Index    string `env:"INDEX_ELASTICSEARCH,required"`
}

type LogData struct {
	Err             error             `json:"err"`
	Description     string            `json:"description"`
	StartTime       time.Time         `json:"startTime"`
	TraceHeader     map[string]string `json:"traceHeader"`
	Request         HttpData          `json:"request"`
	Response        HttpData          `json:"response"`
	RequestBackend  HttpData          `json:"requestBackend"`
	ResponseBackend HttpData          `json:"responseBackend"`

	level        string
	packageName  string
	functionName string
}

type HttpData struct {
	Path    string      `json:"path"`
	Verb    string      `json:"verb"`
	Headers interface{} `json:"headers"`
	Body    interface{} `json:"body"`
}

type Logs struct {
	ID              uint      `json:"id" gorm:"column:id"`
	Level           string    `json:"level" gorm:"column:level"`
	Message         string    `json:"message" gorm:"column:message"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:created_at"`
	Request         string    `json:"request" gorm:"type:JSONB; default:'{}'"`
	Response        string    `json:"response" gorm:"type:JSONB; default:'{}'"`
	RequestBackend  string    `json:"request_backend" gorm:"type:JSONB; default:'{}'"`
	ResponseBackend string    `json:"response_backend" gorm:"type:JSONB; default:'{}'"`
	PathError       string    `json:"path_error"`
	ResponseTime    string    `json:"response_time"`
	TraceHeader     string    `json:"trace_header" gorm:"type:JSONB; default:'{}'"`
}

type iAm struct {
	Name string
	Host string
	Port string
}

func responseTimeString(startTime time.Time) (timeInt int64, timeFmt string) {
	if startTime.IsZero() {
		return 0, ""
	}
	timeDelta := time.Since(startTime)
	return timeDelta.Milliseconds(), fmt.Sprintf("%s", timeDelta)
}

func getPackageAndFuncName() (pkgName string, fncName string) {
	// Get package name and function name
	pc, _, _, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash

	return funcName[:lastDot], funcName[lastDot+1:]
}

func getErrorStack(err error) (errorCause string, errorString string) {
	if err == nil {
		return "", ""
	}

	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	errorCause = fmt.Sprintf("%+v", st[2:3])
	errorString = err.Error()

	return errorCause, errorString
}
