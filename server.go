package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func applyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("r: %+v\n", r)
	fmt.Printf("header: %+v\n", r.Header)
	for k := range r.Header {
		fmt.Println("k: %v", k)
	}
	fmt.Printf("Length: %+v\n", r.ContentLength)
	fmt.Printf("Body: %+v\n", r.Body)
	fmt.Printf("body: %+v\n", r.FormValue("body"))
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("no Error ")
	}
	fmt.Printf("form: %+v\n", r.Form)
	fmt.Printf("postForm: %+v\n", r.PostForm)
	fmt.Println("Applied")
	fmt.Println("job name", r.Form.Get("jobName"), "EOL")
	jobName := r.Form.Get("jobName")
	company := r.Form.Get("company")
	contact := r.Form.Get("contact")
	var entry []byte = []byte(fmt.Sprintf("%v,%v,%v\n", company, jobName, contact))

	f, err := os.OpenFile("apply.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Does not exist?")
		f, err := os.OpenFile("apply.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}
		f.Write([]byte("company,jobName,contact\n"))
	}
	defer f.Close()

	log.Println("entry", string(entry))
	_, err = f.Write(entry)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	http.Handle("/", http.FileServer(http.Dir("")))
	fmt.Println("Listening on port 8080")
	http.HandleFunc("/apply/", applyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
