package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]
        fmt.Fprintf(w, "Title: %s, Page: %s", title, page)
	})
	h := mux.NewRouter()
	h.HandleFunc("/cal/{cal}/plus/{plus}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		// Extract raw values from URL parameters
		cal, err := strconv.Atoi(vars["cal"])
		if err != nil {
			http.Error(w, "Invalid cal parameter", http.StatusBadRequest)
			return
		}

		plus, err := strconv.Atoi(vars["plus"])
		if err != nil {
			http.Error(w, "Invalid plus parameter", http.StatusBadRequest)
			return	
		}
        fmt.Fprintf(w, " this is : %d + %d = %d\n",cal, plus,cal+plus)
        fmt.Fprintf(w, " this is : %d - %d = %d\n",cal, plus,cal-plus)
        fmt.Fprintf(w, " this is : %d * %d = %d\n",cal, plus,cal*plus)
        fmt.Fprintf(w, " this is : %d / %d = %d\n",cal, plus,cal/plus)
	})


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello world")
    })
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "My Name is Anun")
    })

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))


    http.ListenAndServe(":8080", h)
}