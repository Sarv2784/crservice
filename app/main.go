package main

import (
	"crservice/app/config"
	"crservice/app/router"
	"crservice/grpc"
	"flag"
	"fmt"
	"log"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
	go func() {
		if err := grpc.NewGRPCServer(&cfg); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	r := router.InitializeRouter(&cfg)
	host := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	err = r.Run(host)
	if err != nil {
		return
	}
}
