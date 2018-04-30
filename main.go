package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func main() {
	log.SetPrefix("gcloudroku ")
	port := os.Getenv("PORT")
	_, err := strconv.Atoi(port)
	if err != nil {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, _ := httputil.DumpRequest(r, true)
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		log.Printf("%s %s %s %s %s", r.Method, r.URL, r.Proto, r.UserAgent(), ip)
		w.Write(dump)
	})
	log.Println("Listen port " + port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}
