/*
 * @Author: wenchao
 * @Date: 2021-11-01 12:19:10
 * @LastEditors: wenchao
 * @LastEditTime: 2021-11-13 05:11:49
 * @Description:
 */
package main

import (
	"github.com/dungeonsnd/gocom/log4go"
)

// "tinygo.org/x/bluetooth"

func main() {
	// aes.Run()
	// rsa.Run()
	// thumbnail.ResizeImage("", "", 0)

	log4go.InitLog(".", "test", 0, 1000, 10485760, 100)
	log4go.SetLogLevel(5) // InfoLevel=4, DebugLevel=5, logrus.TraceLevel=6
	// log4go.I("================[%v Started]================\n\n", os.Args[0])
	log4go.D("aaa\n\n")
	log4go.D("bbb")
}
