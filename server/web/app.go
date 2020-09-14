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

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}
	usersHandler := app.GetUsers
	if !cors {
		usersHandler = disableCors(usersHandler)
	}
	app.handlers["/api/users"] = usersHandler
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

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	technologies, err := a.d.GetUsers()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(technologies)
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
