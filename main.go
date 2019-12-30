package main

import (
	"fmt"
	"log"
	"net/http"
)

func Redicect301(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("cache get")
		http.Redirect(w, r, "/foo/get", http.StatusMovedPermanently)
	case "POST":
		fmt.Println("cache post")
		http.Redirect(w, r, "/foo/post", http.StatusMovedPermanently)
	default:
		fmt.Println("cache other")
		http.Redirect(w, r, "/foo/rest", http.StatusMovedPermanently)
	}
}

func FooGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + "\n"))
}

func FooPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + "\n"))
}

func FooRest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + "\n"))
}

func Redicect302(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Println("no cache get")
		http.Redirect(w, r, "/bar/get", http.StatusFound)
	case "POST":
		fmt.Println("no cache post")
		http.Redirect(w, r, "/bar/post", http.StatusFound)
	default:
		fmt.Println("no cache other")
		http.Redirect(w, r, "/bar/rest", http.StatusFound)
	}
}

func BarGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + "\n"))
}

func BarPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + "\n"))
}

func BarRest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.Method + "\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/foo", Redicect301)
	mux.HandleFunc("/foo/get", FooGet)
	mux.HandleFunc("/foo/post", FooPost)
	mux.HandleFunc("/foo/rest", FooRest)

	mux.HandleFunc("/bar", Redicect302)
	mux.HandleFunc("/bar/get", BarGet)
	mux.HandleFunc("/bar/post", BarPost)
	mux.HandleFunc("/bar/rest", BarRest)

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Println(err)
	}
}
