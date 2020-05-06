package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	LISTEN_ADDRESS = ":8080"
	STORAGR_ROOT   = "D:\\code\\go\\src\\ObjectStorage\\data"
)

type object struct {
}

func main() {
	http.HandleFunc("/object/", Handler)
	http.ListenAndServe(LISTEN_ADDRESS, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}

	if m == http.MethodGet {
		get(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

///object/<object_name>
func put(w http.ResponseWriter, r *http.Request) {
	f, err := os.Create(STORAGR_ROOT + "\\object\\" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}

func get(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(STORAGR_ROOT + "\\object\\" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
