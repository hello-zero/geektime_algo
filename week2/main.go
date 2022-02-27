package main

import (
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request)  {
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	for k, v := range r.Header{
		for _, v1 := range v{
			w.Header().Set(k, v1)
		}
	}
	log.Printf("recv a %s request from %s, the resp code is %d", r.Method, r.RemoteAddr, 200)
}

func healthz(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("health"))
}

func main()  {
	os.Setenv("VERSION", "v1.0.0")
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", mux)
	if err != nil{
		log.Fatalf("start server failed, err: %v", err)
	}
}
