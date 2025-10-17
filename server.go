package main

import (
	"net/http"	
	"os"
	"fmt"
	"log"
	"io/ioutil"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, %s! You are %s years old.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	  data, err := ioutil.ReadFile("opa/opa.txt")
		if err != nil {
			log.Fatalf("Error reading file: ", err)
		}
		fmt.Fprintf(w, "opass: %s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "Hello, %s! Your password is %s.", user, password)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
  duration := time.Since(startedAt)
 
  if duration.Seconds() < 10 {
    w.WriteHeader(500)
    w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
  } else {
    w.WriteHeader(200)
    w.Write([]byte("ok"))
  }
}
