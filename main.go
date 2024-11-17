package main

import (
	"fmt"
	"net/http"
	"log"
)

func formSubmitHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Fprintf(w,"ParsingForm() error : %v", err)
		return
	}

	fmt.Fprintf(w, "POST request sucessful \n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	email := r.FormValue("email")
	fmt.Fprintf(w,"Name = %v \n",name)
	fmt.Fprintf(w,"address = %s \n",address)
	fmt.Fprintf(w,"email = %v \n",email)
}

func formHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"./static/form.html")
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w,"Hello wanderer!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/form/submit",formSubmitHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting the server at port 8080 \n")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Fatal(err)
	}
	
}