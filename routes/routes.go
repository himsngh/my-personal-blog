package routes

import (
	"html/template"
	"log"
	"net/http"
	"os"
)


type Post struct {
	Title string
	Data string
	Date string
	Author Author
}

type Author struct {
	 Name string
}

var post = []*Post {
	{
		Title: "Hello",
		Data: "World",
		Date: "Today",
		Author : Author {
			Name: "Him",
		},
	},
	{
		Title: "Test",
		Data: "Testing-Data",
		Date: "Today",
		Author : Author {
			Name: "Him-2",
		},
	},
}

func Initialize() (http.Handler, error) {

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	handler := http.NewServeMux()

	tpl := template.Must(template.ParseGlob(cwd + "/templates/*.html"))

	fs := http.FileServer(http.Dir(cwd + "/static"))
	handler.Handle("/static/", http.StripPrefix("/static", fs))

	// Register all the routes
	handler.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		handleHome(writer, request, tpl)
	})
	handler.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		handleAbout(writer, request, tpl)
	})

	return handler, nil
}

func handleHome(w http.ResponseWriter, r *http.Request, tpl *template.Template) {

	if err := tpl.ExecuteTemplate(w, "home.html", post); err != nil {
		log.Fatal(err)
		return
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request, tpl *template.Template) {

	if err := tpl.ExecuteTemplate(w,"about.html", nil); err != nil {
		return
	}
}
