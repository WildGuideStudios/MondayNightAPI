package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

const cert string = "self.crt"      // constant for https tls certificate
const key string = "self.key"       // constant for https tls key

/* This setup will pull the certificate and key from the environment.
const cert string = os.Getenv("TLS_CERT") // constant for https tls certificate
const key  string = os.Getenv("TLS_KEY")  // constant for https tls key
*/

// Context is a type that defines what a context can take
type Context struct {
	Title  string
	Static string
}

// get is a function to handle get requests
func get(w http.ResponseWriter, req *http.Request) {
	response := ""
	switch req.RequestURI {
		case "/get/accounts":
			mapD := map[string]int{"apple": 5, "lettuce": 7}
			mapB, _ := json.Marshal(mapD)
			response = string(mapB)
			break
		default:
			r, _ := json.Marshal("Request not found")
			response = string(r)
			break
	}

	context := Context{Title: response}
	render(w, "api", context)
}

// set is a function to handle get requests
func set(w http.ResponseWriter, req *http.Request) {
	response := ""
	switch req.RequestURI {
	case "/get/accounts":
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		response = string(mapB)
		break
	default:
		r, _ := json.Marshal("Request not found")
		response = string(r)
		break
	}

	context := Context{Title: response}
	render(w, "api", context)
}

// update is a function to handle get requests
func update(w http.ResponseWriter, req *http.Request) {
	response := ""
	switch req.RequestURI {
	case "/get/accounts":
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		response = string(mapB)
		break
	default:
		r, _ := json.Marshal("Request not found")
		response = string(r)
		break
	}

	context := Context{Title: response}
	render(w, "api", context)
}

// create is a function to handle get requests
func create(w http.ResponseWriter, req *http.Request) {
	response := ""
	switch req.RequestURI {
	case "/get/accounts":
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		response = string(mapB)
		break
	default:
		r, _ := json.Marshal("Request not found")
		response = string(r)
		break
	}

	context := Context{Title: response}
	render(w, "api", context)
}

// markDelete is a function to handle get requests
func markDelete(w http.ResponseWriter, req *http.Request) {
	response := ""
	switch req.RequestURI {
	case "/get/accounts":
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		response = string(mapB)
		break
	default:
		r, _ := json.Marshal("Request not found")
		response = string(r)
		break
	}

	context := Context{Title: response}
	render(w, "api", context)
}

// restore is a function to handle get requests
func restore(w http.ResponseWriter, req *http.Request) {
	response := ""
	switch req.RequestURI {
	case "/get/accounts":
		mapD := map[string]int{"apple": 5, "lettuce": 7}
		mapB, _ := json.Marshal(mapD)
		response = string(mapB)
		break
	default:
		r, _ := json.Marshal("Request not found")
		response = string(r)
		break
	}

	context := Context{Title: response}
	render(w, "api", context)
}

// home is a function to display the homepage
func home(w http.ResponseWriter, req *http.Request) {
	rootPath := "views/"
	file := "home.html"
	filePath := rootPath + file
	if (len(filePath) != 0) && (req.RequestURI == "/") {
		f, err := http.Dir(rootPath).Open(file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, filePath, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}

// render is a function to handle rendering
func render(w http.ResponseWriter, tmpl string, context Context) {
	tmplList := []string{"views/base.html",
		fmt.Sprintf("views/%s.html", tmpl)}
	t, err := template.ParseFiles(tmplList...)
	checkErr(err)
	checkErr(t.Execute(w, context))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/get/", get)
	http.HandleFunc("/set/", set)
	http.HandleFunc("/update/", update)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/delete/", markDelete)
	http.HandleFunc("/restore/", restore)
	if (len(cert) > 0) && (len(key) > 0) {
		fmt.Print("Listening on https/8443")
		err := http.ListenAndServeTLS(":8443", cert, key, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		fmt.Print("Listening on http/8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
