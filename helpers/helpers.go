package helpers

import (
	"net/http"
	"log"
	"os"
)

const (
	LOG_FILE = "./logs/myapp_log.txt"
)

func Log (_ http.ResponseWriter, r *http.Request) {
    logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
        log.Panic(err)
    }
    defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println(r.URL.Path + ": Logging to custom file")
}