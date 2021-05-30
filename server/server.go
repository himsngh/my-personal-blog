package server

import (
	"github.com/himsngh/my-personal-blog/server/database"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Server struct {
	dbStore  database.Store
	template *template.Template
}

var CurrentWorkingDirectory, _ = os.Getwd()


func NewServer() (*Server, error) {

	tpl := template.Must(template.ParseGlob(CurrentWorkingDirectory + "/templates/*.html"))

	dbStore, err := database.NewDatabaseStore()
	if err != nil {
		log.Println("Error creating the database store: err ", err.Error())
		return nil, err
	}

	return &Server{
		dbStore: dbStore,
		template: tpl,
	}, nil
}

func (s *Server) ServeRoutes() (http.Handler, error) {

	handler := http.NewServeMux()

	fs := http.FileServer(http.Dir(CurrentWorkingDirectory + "/static"))
	handler.Handle("/static/", http.StripPrefix("/static", fs))

	// Register all the server
	handler.HandleFunc("/signup", func(writer http.ResponseWriter, request *http.Request) {
		s.handleSignUp(writer, request)
	})
	handler.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		s.handleLogin(writer, request)
	})
	handler.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		s.handleHome(writer, request)
	})
	handler.HandleFunc("/profile", func(writer http.ResponseWriter, request *http.Request) {
		s.handleProfile(writer, request)
	})

	return handler, nil
}

func (s *Server) handleSignUp(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		if err := s.template.ExecuteTemplate(w,"signup.html", nil); err != nil {
			log.Println("Error executing about template : ", err.Error())
			return
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			log.Println("Error Parsing Form Data err : ", err.Error())
			return
		}

		// TODO create accout for the user
		log.Println(r.PostForm)
		log.Println(r.Form)
		//firstName := r.FormValue("firstName")
		//lastName := r.FormValue("lastName")
		//email := r.FormValue("email")

		// logged in
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)

	default:
		http.Error(w, "wrong route", http.StatusBadRequest)
		return
	}

}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	// TODO IMPLEMENT ME
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {

	postList, err := s.dbStore.ListPost()
	if err != nil {
		return
	}

	if err := s.template.ExecuteTemplate(w, "home.html", postList); err != nil {
		log.Println("Error executing home template : ", err.Error())
		return
	}
}

func (s *Server) handleProfile(w http.ResponseWriter, r *http.Request) {

	if err := s.template.ExecuteTemplate(w,"about.html", nil); err != nil {
		log.Println("Error executing about template : ", err.Error())
		return
	}
}
