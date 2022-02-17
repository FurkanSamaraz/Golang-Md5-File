package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var Version string
var satir string

var h = md5.New()
var c string

func main() {

	http.HandleFunc("/", httpF)
	http.ListenAndServe(":8080", nil)
}

func httpF(w http.ResponseWriter, r *http.Request) {

	//file, _ := os.Create("dosya.txt") // dosya oluştur.
	f, err := os.Open("dosya.txt")
	checkError(err)

	defer f.Close()
	h = md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	c = string(h.Sum(nil))
	fmt.Printf("%x", h.Sum(nil))

	fmt.Fprintf(w, c)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
