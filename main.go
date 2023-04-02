package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"vizzy/bootstrap"
	"vizzy/lorca"
)

var mode string

func loadWeb() {
	if mode == "prod" {
		fs := http.FileServer(http.Dir("resources/"))
		http.Handle("/", fs)
		go func() {
			if err := http.ListenAndServe(":5173", nil); err != nil {
				logrus.Fatal(err)
			}
		}()
	}
	// 加载 React 应用程序
	ui, err := lorca.New("", "",
		480, 320,
		"--remote-allow-origins=*",
		"--kiosk",
		"–disable-automation",
		"--remote-debugging-port=9222",
	)
	if err != nil {
		logrus.Panic(err)
	}
	err = ui.Load("http://localhost:5173/")
	if err != nil {
		logrus.Panic(err)
	}
	defer ui.Close()
	// Wait for the browser window to be closed
	<-ui.Done()
}
func main() {
	go func() {
		bootstrap.Bootstrap(mode)
	}()
	// 创建一个带有缓冲的channel，长度为1
	ch := make(chan bool, 1)

	// 向channel中写入一个值
	ch <- true

	// 等待channel中的值被读取
	<-ch

	loadWeb()

}
