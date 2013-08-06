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
  
  id := r.FormValue("id")
  if id != "3db31ee9cf28" {
		w.WriteHeader(500)
		fmt.Fprint(w, "bad client")
		return  
  }

 	rm := map[string]string{}
	rm["clientid"] = id;
	rm["topic"] = r.FormValue("tp");
  rm["update"] = "http://i.imgur.com/1b5X8Vl.jpg"
  
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
