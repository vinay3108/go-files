package main

import (
	"fmt"
	"log"
	"net/http"
)
func formHandler(w http.ResponseWriter,r *http.Request){
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"Parseform() err: %v",err)
		return
	}
	fmt.Fprintf(w,"Post Request SuceessFull\n");
	name:=r.FormValue("name")
	email:=r.FormValue("email")
	fmt.Fprintf(w,"Name =%s and email =%s \n",name,email);
}
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"method is not supported",http.StatusNotFound);
	}
	fmt.Fprintf(w,"hello!!!!!!!!!!!")
}

func main(){
		fileServer:=http.FileServer(http.Dir("./static"))
		http.Handle("/",fileServer);
		http.HandleFunc("/form",formHandler);
		http.HandleFunc("/hello",helloHandler);
		fmt.Printf("starting sercer at port 8000\n")
		if err:=http.ListenAndServe(":8080",nil); err!=nil{
			log.Fatal(err);
		}
}