package main

import (
	"fmt"
	"go-tcp-kvs/server"
	"go-tcp-kvs/server/logger"
	"go-tcp-kvs/store"
	"net"
	"os"
	"strconv"
)

var done chan bool

func main() {
	logFile := logger.SetLogs()
	defer func(logFile *os.File) {
		_ = logFile.Close()
	}(logFile)

	done = make(chan bool)

	args := os.Args

	port := setPort(args)

	store.InitStore()

	go startTCP(port)

	<-done
	os.Exit(0)
}

func setPort(args []string) string {
	if len(args) < 3 || args[1] != "--port" {
		logger.ErrorLogger.Fatal("exit code -1: format should be './store --port <port>'")
	}
	port, err := strconv.Atoi(args[2])
	if err != nil {
		logger.ErrorLogger.Fatalf("exit code -1: failure to parse %s into an integer port", args[2])
	}
	return fmt.Sprintf(":%d", port)
}

func startTCP(port string) {
	l, err := net.Listen("tcp4", port)
	if err != nil {
		logger.ErrorLogger.Println(err.Error())
	}
	defer func() { _ = l.Close() }()

	for {
		c, err2 := l.Accept()
		if err2 != nil {
			logger.ErrorLogger.Println(err2)
			return
		}

		go server.HandleConnection(c, done)
	}
}
