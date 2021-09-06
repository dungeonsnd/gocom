package log4go

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

var logPath string
var logFilename string
var maxAge uint32
var rotationTime int64
var rotationSize int64
var rotationCount uint32

var (
	singleton *LoggerMgr
	once      sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() *LoggerMgr {
	once.Do(func() {
		singleton = &LoggerMgr{}
		singleton.init()
	})

	return singleton
}

// LOGLEVEL 日志级别
type LOGLEVEL int

const (
	_ LOGLEVEL = iota
	TraceLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func T(format string, args ...interface{}) {
	Logf("", TraceLevel, format, args...)
}
func D(format string, args ...interface{}) {
	Logf("", DebugLevel, format, args...)
}
func I(format string, args ...interface{}) {
	Logf("", InfoLevel, format, args...)
}
func W(format string, args ...interface{}) {
	Logf("", WarnLevel, format, args...)
}
func E(format string, args ...interface{}) {
	Logf("", ErrorLevel, format, args...)
}
func F(format string, args ...interface{}) {
	Logf("", FatalLevel, format, args...)
}

func InitLog(writePath, filename string, _maxAge uint32, _rotationTime int64, _rotationSize int64, _rotationCount uint32) {
	if len(writePath) > 0 {
		logPath = writePath
	}
	if len(filename) > 0 {
		logFilename = filename
	}
	maxAge = _maxAge
	rotationTime = _rotationTime
	rotationSize = _rotationSize
	rotationCount = _rotationCount

	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := os.MkdirAll(logPath, 0777); err != nil {
		log.Fatalf("conf path make failed.err：%+v", err)
	}
}

// Logger logger
type Logger struct {
	logger *logrus.Logger
	Name   string
}

// LoggerMgr log管理器
type LoggerMgr struct {
	logs map[string]*Logger
}

func (mgr *LoggerMgr) init() {
	mgr.logs = make(map[string]*Logger)
}

// GetLogger 获得一个logger
func (mgr *LoggerMgr) GetLogger(logName string) *Logger {
	// 默认文件名
	if logName == "" {
		logName = logFilename
	}
	if logger, ok := mgr.logs[logName]; ok {
		return logger
	}

	return mgr.InitLogger(logName)
}

// InitLogger 初始化访问日志
func (mgr *LoggerMgr) InitLogger(logName string) *Logger {

	var logInstance Logger
	logInstance.Name = logName

	logger := logrus.New()
	fileName := path.Join(logPath, fmt.Sprintf("%s.log", logName))
	logPath, _ := filepath.Abs(fileName)

	//禁止logrus的输出
	// src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend|0777)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// 设置日志输出的路径
	// logger.Out = src
	// logger.Out = os.Stdout

	logger.SetLevel(logrus.TraceLevel)
	logger.SetReportCaller(true)
	var logWriter *rotatelogs.RotateLogs
	if rotationTime > 0 {
		logWriter, _ = rotatelogs.New(
			logPath+".%Y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(logPath),                                   // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),             // 文件最大保存时间
			rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Hour), // 日志切割时间间隔
			rotatelogs.WithRotationCount(uint(rotationCount)),
		)
	} else {
		logWriter, _ = rotatelogs.New(
			logPath+".%Y-%m-%d-%H-%M.log",
			rotatelogs.WithLinkName(logPath),                       // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour), // 文件最大保存时间
			rotatelogs.WithRotationSize(rotationSize),
			rotatelogs.WithRotationCount(uint(rotationCount)),
		)
	}

	// writeMap := lfshook.WriterMap{
	// 	logrus.InfoLevel:  logWriter,
	// 	logrus.FatalLevel: logWriter,
	// 	logrus.DebugLevel: logWriter, // 为不同级别设置不同的输出目的
	// 	logrus.WarnLevel:  logWriter,
	// 	logrus.ErrorLevel: logWriter,
	// 	logrus.PanicLevel: logWriter,
	// }

	// _, file, line, ok := runtime.Caller(2)
	// if !ok {
	// 	file = "???"
	// 	line = 0
	// }

	formatter := &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			//处理文件名
			fileName := path.Base(frame.File)
			return frame.Function, fileName + ":" + strconv.Itoa(frame.Line)
		}}

	// lfHook := lfshook.NewHook(writeMap, formatter)
	// logger.AddHook(lfHook)

	filenameHook := filename.NewHook()
	filenameHook.Field = "file" // Customize source field name
	// filenameHook.Depth = 6
	filenameHook.SkipPrefixes = append(filenameHook.SkipPrefixes, "log4go/")
	logger.AddHook(filenameHook)

	logger.SetFormatter(formatter)
	if logInstance.Name == "" {
		logger.SetOutput(os.Stdout)
	} else {
		logger.SetOutput(logWriter)
	}
	logInstance.logger = logger
	mgr.logs[logName] = &logInstance
	return &logInstance
}

// Dir log dir
func Dir() string {
	return logPath
}

//  InfoLevel=4, DebugLevel=5, logrus.TraceLevel=6
func SetLogLevel(logLevel logrus.Level) {
	logClient := GetInstance().GetLogger("")
	logClient.logger.SetLevel(logLevel)
}

// Logf 格式化输出
func Logf(logName string, logLevel LOGLEVEL, format string, args ...interface{}) {
	logClient := GetInstance().GetLogger(logName)

	switch logLevel {
	case TraceLevel:
		logClient.logger.Tracef(format, args...)
	case DebugLevel:
		logClient.logger.Debugf(format, args...)
	case InfoLevel:
		logClient.logger.Infof(format, args...)
	case WarnLevel:
		logClient.logger.Warnf(format, args...)
	case ErrorLevel:
		logClient.logger.Errorf(format, args...)
	case PanicLevel:
		logClient.logger.Panicf(format, args...)
	case FatalLevel:
		logClient.logger.Fatalf(format, args...)
	}
}

// Logln 换行输出
func Logln(logName string, logLevel LOGLEVEL, args ...interface{}) {
	logClient := GetInstance().GetLogger(logName)

	switch logLevel {
	case TraceLevel:
		logClient.logger.Traceln(args...)
	case DebugLevel:
		logClient.logger.Debugln(args...)
	case InfoLevel:
		logClient.logger.Infoln(args...)
	case WarnLevel:
		logClient.logger.Warnln(args...)
	case ErrorLevel:
		logClient.logger.Errorln(args...)
	case PanicLevel:
		logClient.logger.Panicln(args...)
	case FatalLevel:
		logClient.logger.Fatalln(args...)
	}
}
