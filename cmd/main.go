package main

import (
	"cuckoo/internal/adapter/interface"
	"cuckoo/internal/conf"
	"cuckoo/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			printStack(err)
		}
	}()

	conf.Init()

	logger.Init(conf.GetLoggerConfig())

	go _interface.Run()

	// 退出信号监听
	qc := make(chan os.Signal, 1)
	signal.Notify(qc, syscall.SIGINT, syscall.SIGTERM)
	<-qc
}

func printStack(err any) {
	const size = 16 << 10
	buf := make([]byte, size)
	buf = buf[:runtime.Stack(buf, false)]
	logger.Error("SystemCrash!!!",
		"stack", fmt.Sprintf("\n%s", buf),
		"panic", true,
		"error", err,
	)
}
