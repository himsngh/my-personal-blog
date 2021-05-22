package routes

import (
	"html/template"
	"net/http"
	"os"
)

func Initialize() (http.Handler, error) {

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	handler := http.NewServeMux()


	fs := http.FileServer(http.Dir(cwd + "/static"))
	handler.Handle("/static/", http.StripPrefix("/static",fs))

	// Register all the routes
	handler.HandleFunc("/home", handleHome)
	handler.HandleFunc("/about", handleAbout)

	return handler, nil
}

func handleHome(w http.ResponseWriter, r *http.Request) {

	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	tpl, err := template.ParseFiles(cwd + "/templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		return
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {

	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	tpl, err := template.ParseFiles(cwd + "/templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, nil); err != nil {
		return
	}
}