package main

import (
	"github.com/dungeonsnd/gocom/test/encrypt/aes"
	"github.com/dungeonsnd/gocom/test/encrypt/rsa"
)

func main() {
	aes.Run()
	rsa.Run()
	// thumbnail.ResizeImage("", "", 0)

	// log4go.InitLog(".", "test", 0, 1000, 10485760, 100)
	// log4go.SetLogLevel(5) // InfoLevel=4, DebugLevel=5, logrus.TraceLevel=6
	// log4go.I("================[%v Started]================\n", os.Args[0])

}
