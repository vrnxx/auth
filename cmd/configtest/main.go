package main

import (
	"fmt"
	infraConf "github.com/vrnxx/auth/internal/infrastructure/config"
	mainConf "github.com/vrnxx/auth/internal/presentation/config"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello, World!"))
}

func main() {
	c := mainConf.Config{}
	infraConf.LoadConfig(&c, "/home/am/Projectus/study/homework/auth/internal/configs/app/dev.toml")
	fmt.Println(c)

	http.HandleFunc("/", handler)
	port := ":8000"
	fmt.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
