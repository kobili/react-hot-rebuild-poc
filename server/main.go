package main

import (
	"log"
	"net/http"
	"regexp"
)

/*
Serves a static React Router app. Handles the case where the client refreshes on a URL
set client side by the React Router Javascript.
*/
func serveReactRouterApp(fs http.Handler) http.HandlerFunc {
	// Regex to check if a string ends in a file extension (.js, .css, etc)
	fileMatcher := regexp.MustCompile(`\.[a-zA-Z]*$`)

	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if !fileMatcher.MatchString(req.URL.Path) {
			// if the request is not for a specific file, it is probably a reload on a url set by react router
			// in this case, just serve the entry point into the react app and let the client side js handle the redirect
			http.ServeFile(w, req, "./static/index.html")
		} else {
			// This is still necessary as the frontend will still expect javascript and css to be served
			fs.ServeHTTP(w, req)
		}
	}

	return http.HandlerFunc(fn)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", serveReactRouterApp(fs))

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		if method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("hello world"))
	})

	log.Print("Server listening on localhost:2345")

	err := http.ListenAndServe(":2345", nil)
	if err != nil {
		log.Fatal(err)
	}
}
