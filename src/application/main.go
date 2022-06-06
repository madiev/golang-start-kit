package main

import (
    "start_kit/src/application/handler"
    "start_kit/src/application/config"
    "encoding/json"
    "net/http"
    "strconv"
    "flag"
    "time"
    "fmt"
    "log"
    "os"
)

var c *config.Config

func init() {
    conf := flag.String("conf", "conf/conf.json", "a string")

    flag.Parse()

    log.Printf("read conf => %s", *conf)

    file, err := os.Open(*conf)
    if err != nil {
        log.Println(err)
    }
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&c)
    if err != nil {
        log.Println(err)
    }
}

func main() {
    server := &http.Server{
        Addr: ":" + strconv.Itoa(c.Server.Addr),
        Handler: &handler.Handler{
            Config: c,
        },
        ReadTimeout:  time.Duration(c.Server.ReadTimeout) * time.Millisecond,
        WriteTimeout: time.Duration(c.Server.WriteTimeout) * time.Millisecond,
    }

    server.SetKeepAlivesEnabled(false)
    fmt.Printf("--> listen :%v\n", c.Server.Addr)
    log.Fatal(server.ListenAndServe())
}
