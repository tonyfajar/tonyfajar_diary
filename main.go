package main

import (
	"io"
	"log"
	"net/http"
	"flag"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/tonyfajar/tonyfajar_diary/app/types"
)



func initConfig() {

	if _, err := toml.DecodeFile("config.toml", &types.Config.config); err != nil {
		fmt.Println(err)
	}
	flag.StringVar(&types.Config.ConfigPath, "config", "config.toml", "Path to config file")
  	flag.Parse()
}


func main() {
	initConfig()
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
    log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}