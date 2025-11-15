package lib

import "net/http"

type User struct {
	Name     string
	UserName string
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/user/1", userHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, "Rate Limiter Test")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, User{
		Name:     "Test",
		UserName: "Test",
	})
}
