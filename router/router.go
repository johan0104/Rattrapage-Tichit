package router

import (
	"TICHIT/controller"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Serveur() {
	http.HandleFunc("/accueil", controller.AccueilHandler)
	http.HandleFunc("/ajout", controller.AjoutHandler)
	http.HandleFunc("/detail/", controller.DetailHandler)


	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	// Lance le serveur
	runServer()
}

// Fonction qui permet de lancer le serveur
func runServer() {
	port := "localhost:8080"
	url := "http://" + port + "/accueil"
	go http.ListenAndServe(port, nil)
	fmt.Println("Server is running...")
	time.Sleep(time.Second * 3)
	fmt.Println("If the navigator didn't open on its own, just go to ", url, " on your browser.")
	isRunning := true
	for isRunning {
		fmt.Println("If you want to end the server, type 'stop' here :")
		var command string
		fmt.Scanln(&command)
		if command == "stop" {
			isRunning = false
		}
	}
}
