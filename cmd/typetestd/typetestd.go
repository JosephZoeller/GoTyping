package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/JosephZoeller/project-0/cmd/phase"
)

func main() {
	saves, er := phase.GetSave()
	if er != nil {
		log.Println(er)
	}
	hostSave(saves)
}

func hostSave(saves *phase.SaveFile) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("Displaying content")
		t, _ := template.ParseFiles("../../web/tables.html")
		t.Execute(res, *saves)
	})

	errorChan := make(chan error)
	fmt.Println("Listening on ports 8080 (http)...")
	go func() {
		errorChan <- http.ListenAndServe(":8080", nil)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT)

	for {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Fatalln(err)
			}

		case sig := <-signalChan:
			log.Println("shutting down: ", sig)
			os.Exit(0)
		}
	}

}
