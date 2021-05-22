package routes

import "net/http"

func Initialize() http.Handler {

	handler := http.NewServeMux()
	// Register all the routes
	handler.HandleFunc("/home", handleHome)
	handler.HandleFunc("/about", handleAbout)

	return handler
}

func handleHome(w http.ResponseWriter, r *http.Request) {




}

func handleAbout(w http.ResponseWriter, r *http.Request) {

}