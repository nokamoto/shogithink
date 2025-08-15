// usi-bridge is a simple USI bridge that listens to stdin and outputs to stdout.
// It also provides an HTTP server to view logs.
package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	logs   []string
	logsMu sync.Mutex
)

func appendLog(s string) {
	logsMu.Lock()
	defer logsMu.Unlock()
	logs = append(logs, s)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	logsMu.Lock()
	defer logsMu.Unlock()
	for _, log := range logs {
		fmt.Fprintln(w, log)
	}
}

func startHTTPServer() {
	http.HandleFunc("/", logsHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	go startHTTPServer()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		appendLog("recv: " + line)
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "usi":
			fmt.Println("id name usi-bridge")
			fmt.Println("id author github.com/nokamoto/shogithink")
			fmt.Println("usiok")
			appendLog("send: id name usi-bridge")
			appendLog("send: id author github.com/nokamoto/shogithink")
			appendLog("send: usiok")
		case "isready":
			fmt.Println("readyok")
			appendLog("send: readyok")
		case "position":
			appendLog("position command received")
		case "go":
			fmt.Println("info score cp 0")
			appendLog("send: info score cp 0")
			fmt.Println("bestmove resign")
			appendLog("send: bestmove resign")
		case "quit":
			appendLog("quit received, exiting")
			return
		default:
			appendLog("unknown command: " + line)
		}
	}
}
