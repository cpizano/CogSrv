package cogsrv

import (
    "encoding/json"
    "fmt"
    "net/http"

    "appengine"
)

func init() {
    http.HandleFunc("/", top_handler)
    http.HandleFunc("/reg", reg_handler)
}

func top_handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")

    fmt.Fprintf(w, "cogsrv  (c) 2013 vortex code\n")
    fmt.Fprintf(w, "appengine app id = %q\n", appengine.AppID(c))
    fmt.Fprintf(w, "appengine version = %q\n", appengine.VersionID(c))

    name, index := appengine.BackendInstance(c)
    fmt.Fprintf(w, "appengine backendinstance = %q, %d\n", name, index)
}

func reg_handler(w http.ResponseWriter, r *http.Request) {  
  
  type Message struct {
    ServerId string
    Update string
    Version int32
    Path string
  }
  
  rm := Message{"773377", "http://localhost:8080/file1.txt", 2, "data/file1.txt"}
	writeJSON(w, rm)
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	buf, err := json.Marshal(i)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, "json.Marshal failed: %v", err)
		return
	}
  w.Header().Set("Content-Type", "text/json; charset=utf-8")
	w.Write(buf)
}
