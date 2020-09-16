package web

import (
	"DGraph-Example/db"
	"encoding/json"
	"log"
	"net/http"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

func NewApp(d db.DB, prod bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	searchHandler := app.Search
	if !prod {
		searchHandler = disableCors(searchHandler)
	}
	app.handlers["/api/search"] = searchHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 3000")
	return http.ListenAndServe(":3000", nil)
}

func (a *App) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	searchType := r.URL.Query()["searchtype"]
	phrase := r.URL.Query()["phrase"]
	if searchType == nil || phrase == nil {
		sendErr(w, http.StatusBadRequest, "searchtype and phrase are required")
		return
	}
	var response interface{}
	var err error
	switch searchType[0] {
	case "actors":
		response, err = a.d.SearchActors(phrase[0])
		break
	case "movies":
		response, err = a.d.SearchMovies(phrase[0])
		break
	default:
		sendErr(w, http.StatusBadRequest, "Invalid searchtype")
		return
	}
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
