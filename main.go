package main

import (
    "flag"
    "log"

    "github.com/ra0van/ricedb/config"
    "github.com/ra0van/ricedb/server"
)

func setupFlags () {
    flag.StringVar(&config.Host, "host", "0.0.0.0", "host for ricedb server")
    flag.IntVar(&config.Port, "port", 7379, "port for dicedb server")
    flag.Parse()
}

func main() {
    setupFlags()
    log.Println("Starting the rice db 🍚")
    server.RunSyncTcpServer()
}
