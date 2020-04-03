package main
import (
	"html/template"
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/edit-ref.html", "templates/base.html"))
	_ = tmpl.Execute(w, struct{}{})
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
