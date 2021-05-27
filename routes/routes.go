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
}

var post = []*Post {
	{
		Title: "Hello",
		Data: "World",
	},
	{
		Title: "Test",
		Data: "Testing-Data",
	},
}

func Initialize() (http.Handler, error) {

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	handler := http.NewServeMux()

	fs := http.FileServer(http.Dir(cwd + "/static"))
	handler.Handle("/static/", http.StripPrefix("/static", fs))

	tpl, err := template.ParseGlob(cwd + "/templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	// Register all the routes
	handler.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		handleHome(writer, request, tpl)
	})

	handler.HandleFunc("/about", handleAbout)

	return handler, nil
}

func handleHome(w http.ResponseWriter, r *http.Request, tpl *template.Template) {

	//cwd, err := os.Getwd()
	//if err != nil {
	//	return
	//}

	//tpl, err := template.ParseFiles(cwd + "/templates/home.html")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	if err := tpl.ExecuteTemplate(w, "home.html", post); err != nil {
		log.Fatal(err)
		return
	}
	//if err := tpl.Execute(w, post); err != nil {
	//	return
	//}
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
