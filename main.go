package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GDEIDevelopers/K8Sbackend/app"
	"github.com/GDEIDevelopers/K8Sbackend/config"
)

func main() {
	var configfile string
	flag.StringVar(&configfile, "conf", "", "配置文件")
	flag.Parse()

	if configfile == "" {
		log.Fatal("没有配置文件")
	}
	cfg := config.Read(configfile)
	server := app.NewServer(cfg)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	<-sigCh
	server.Close()
}
