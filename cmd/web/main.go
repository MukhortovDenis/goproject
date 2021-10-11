package main

import (
	"log"
	"net/http"
)

func main() {

	// m := &autocert.Manager{
	// 	Cache:      autocert.DirCache("golang-autocert"),
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("https://stoneshop.herokuapp.com", "stoneshop.herokuapp.com", "127.0.01.1:443"),
	// }
	// server := &http.Server{
	// 	Addr:      ":443",
	// 	Handler:   handler,
	// 	TLSConfig: m.TLSConfig(),
	// }
	router := mainHandle()
	// log.Fatal(server.ListenAndServeTLS("", ""))
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt)
	// <-c
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// server.Shutdown(ctx)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
