package main

import (
	"net/http"
	"time"

	"github.com/beautiful-you/anniversary/str-cache"
)

func main() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	http.ListenAndServe(":5645", nil)
}
func get(rw http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	str, bool := cache.Get(key)
	if bool {
		rw.Write([]byte(str))
		return
	}
	rw.Write([]byte(""))
}
func set(rw http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	value := req.FormValue("value")
	// timeD := req.FormValue("time")
	cache.Set(key, value, 120*time.Minute)
	rw.Write([]byte(""))
}
