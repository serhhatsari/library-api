package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			return
		}
	},
	)

	srv := &http.Server{
		Addr: ":8080",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}

}
