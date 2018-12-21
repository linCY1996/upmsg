package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
	"upimg/util"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, `application/json`)
	w.Header().Set(`Access-Control-Allow-Origin`, `*`)
	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	r.ParseMultipartForm(1 << 32)
	_, h, _ := r.FormFile("file")
	// fmt.Println(err)
	ext := path.Ext(h.Filename)

	dir := `res/img/` + time.Now().Format("2006-01-02") + "/" //yyyy-MM-rr
	os.MkdirAll(dir, 0666)
	f, _ := h.Open()
	name := util.RandStr() + ext
	f1, _ := os.Create(dir + name)
	io.Copy(f1, f)
	f.Close()
	f1.Close()
	w.Write([]byte("/" + dir + name))

}

func views(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile(`demo.html`)
	w.Write(buf)
}

func main() {
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
	http.HandleFunc(`/view`, views)
	http.HandleFunc(`/upload`, Upload)
	http.ListenAndServe(":80", nil)
	// fmt.Println(err)
}
