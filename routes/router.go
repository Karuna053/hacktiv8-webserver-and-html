package routes

import (
	"fmt"
	"net/http"
	"webserver-and-html/config"
	// "sesi-enam/controller"
)

func Router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!")
	})

	// http.HandleFunc("/")

	// http.HandleFunc("/greet", controller.HandlerGreet)
	// http.HandleFunc("/api/employees", controller.GetEmployee)
	// http.HandleFunc("/api/employee", controller.CreateEmployee)
	// http.HandleFunc("/indexEmployee", controller.GetEmployeeUI)

	fmt.Println("Server is running on port", config.Port) // Console log
	http.ListenAndServe(config.Port, nil)                 // Actually run serve (in php)
}
