package main

import (
	"fmt"
	"main/configs"
	"main/verify"
	"net/http"
)

func main() {
	emailConfigs, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Print(emailConfigs)
	router := http.NewServeMux()
	verify.NewHandler(router)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}

}
