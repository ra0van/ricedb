package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/ra0van/ricedb/config"
)

func readCommand(c net.Conn) (string, error){
    // TODO : Max read 512 bytes in one go
    // To allow input > 512 bytes, then repeated read until
    // we get EOF or delimiter
    var buf []byte = make([]byte, 512)
    n, err := c.Read(buf[:])
    if err != nil{
        return "", err
    }
    return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
    if _, err := c.Write([]byte(cmd)); err != nil {
        return err
    }
    return nil
}

func RunSyncTcpServer() {
    log.Println("starting synchronous TCP server on", config.Host, config.Port)

    var con_clients int = 0

    // listening to the configured host:port
    lsnr, err := net.Listen("tcp", config.Host + ":" + strconv.Itoa(config.Port))
    if err != nil {
        log.Println("Error listening:", err.Error())
        panic(err)
    }

    // close the listener when the app closes
    defer lsnr.Close()

    for {

        // blocking call : waiting for the new client to connect
        c, err := lsnr.Accept()
        if err != nil {
            panic(err)
        }

        con_clients += 1
        log.Println("client connected with address: ", c.RemoteAddr(), "concurrent clients ", con_clients)

        for {
            cmd, err := readCommand(c)
            if err != nil {
                c.Close()
                con_clients -= 1
                log.Println("client disconnected", c.RemoteAddr(), "concurrent clients", con_clients)
                if err == io.EOF {
                    break
                }
                log.Println("err", err)
            }
            log.Println("command", cmd)
            if err = respond(cmd, c); err != nil {
                log.Print("err write:", err)
            }
        }
    }

}
