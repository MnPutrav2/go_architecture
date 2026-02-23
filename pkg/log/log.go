package log

import (
	"fmt"
	"net/http"
	"time"

	ip "github.com/MnPutrav2/go_architecture/pkg/ip"
)

func Log(message, level string, r *http.Request) {
	ip := ip.ClientIP(r)
	t := time.Now()

	red := "\033[31m"
	green := "\033[32m"
	yellow := "\033[33m"
	reset := "\033[0m"
	var log string

	switch level {
	case "INFO":
		log = green + "INFO" + reset
	case "ERROR":
		log = red + "ERROR" + reset
	case "WARN":
		log = yellow + "WARN" + reset
	}

	fmt.Println("[ LOG ] ", log, "--- Client IP:[", ip, "] Time:[", t.Format("2006-01-02 15:04:05"), "] Message:[", message, "] Path:[", r.URL, "]")
}
