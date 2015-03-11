package main

import (
	"fmt"
	"github.com/spudfkc/gotify"
	"log"
	"net/http"
)

type handler struct{}
type authHandler struct{}
type authCallbackHandler struct{}
type staticHandler struct{}

var auth *gotify.Oauth2Authenticator

func (h *staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth = gotify.CreateAuthenticator(gotify.PLAYLIST_READ_PRIVATE, gotify.PLAYLIST_MODIFY_PRIVATE, gotify.PLAYLIST_MODIFY_PUBLIC)
	authUrl := auth.GetAuthURL()
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
}

func (h *authCallbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth.ParseToken(r)
	gotify.AuthorizeClient(auth)
	w.Header().Set("AmIAuthorized", "Yes")

	user := gotify.User{"spudfkc", "lol", "derp"}
	p, err := gotify.CreatePlaylist("goshuffleplaylist", true, user)
	if err != nil {
		log.Println("Failed to create playlist", err)
	}
	log.Println("playlist", p)

	w.WriteHeader(http.StatusOK)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	h := handler{}
	a := authHandler{}
	c := authCallbackHandler{}

	http.HandleFunc("/", h.ServeHTTP)
	http.HandleFunc("/auth", a.ServeHTTP)
	http.HandleFunc("/auth/cb", c.ServeHTTP)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8081", nil)
	fmt.Println("Listening on 8081...")
}
