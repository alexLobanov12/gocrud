package helpers

import (
	"net/http"
	"log"
	"os"
	"time"
)

var (
	LOG_FILE = "./logs/" + string(time.Now().Format("2006-01-02")) + ".log"
)

const (
	LOG_API = 1
	LOG_CUSTOM = 2
)

func Log (_ http.ResponseWriter, r *http.Request) {
    logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
        log.Panic(err)
    }
    defer logFile.Close()
	userIP := ReadUserIP(r)
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(r.URL.Path +  " ip address: " + userIP)
}

func ReadUserIP(r *http.Request) string {
    IPAddress := r.Header.Get("X-Real-Ip")
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
    }
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
    }
    return IPAddress
}