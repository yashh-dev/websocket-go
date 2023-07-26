package main

import "net/http"

func main(){

	initializeConsole()
	http.ListenAndServe(":3000",nil)
	
}

func initializeConsole(){
	http.Handle("/",http.FileServer(http.Dir("./console")))

}