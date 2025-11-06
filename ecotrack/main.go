/*
package main

import (

	"ecotrack/api"
	"fmt"
	"net/http"

)

	func main() {
		mux := http.NewServeMux()

		api.RegisterRoutes(mux)

		fs := http.FileServer(http.Dir("static"))
		mux.Handle("/static/", http.StripPrefix("/static/", fs))

		fmt.Println("🌱 EcoTrack corriendo en http://localhost:8080")
		http.ListenAndServe(":8080", mux)
	}
*/
package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("🌍 Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
