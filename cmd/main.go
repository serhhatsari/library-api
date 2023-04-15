package main

import (
	"github.com/serhhatsari/library-api/repository"
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

	repo := repository.New()
	defer repo.Close()

	srv := &http.Server{
		Addr: ":3000",
	}

	err := srv.ListenAndServe()
	if err != nil {
		return
	}

}
