package log4go

import (
	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestBindHooks(t *testing.T) {
	hook, err := logrus_sentry.NewSentryHook("", []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		t.Error(err)
	}
	BindHooks(hook)
	E("test BindHooks error")
}
