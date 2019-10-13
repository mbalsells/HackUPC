package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func sampleEnpoint(w http.ResponseWriter, req *http.Request) {
	body, _ := json.Marshal(map[string]interface{}{
		"missatge": "el endpoint funciona",
	})
	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Write(buf.Bytes())
}

// SetupHandlers initiates the servers HTTP endpoints
func SetupHandlers() {
	http.HandleFunc("/sampleEndpoint", sampleEnpoint)
	http.ListenAndServe(":8080", nil)
}

func main() {
	SetupHandlers()

}
