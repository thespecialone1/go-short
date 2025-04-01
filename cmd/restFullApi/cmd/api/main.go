package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/thespecialone1/go-short/cmd/restFullApi/internal/handlers"
)

// in our main function lets first setup our logger so when we get an error we get the file and the line number
// to do this we call the SetReportCaller(true)
func main() {
  logrus.SetReportCaller(true)

  // Now we create a new chi mux variable using the NewRouter function which returns a pointer to a mux type
  var r *chi.Mux = chi.NewRouter()   // it is really just a struct which we will use to setup our api
  
  // we gonna pass this onto our handler function 
  handlers.Handler(r)   // we have defined it in the internal/handlers (we have imported this package up ther) this will setup 
  // our router. i.e add the endpoint definations thar we want 

  // lets add print statement and then the cooler print statement lol

  fmt.Println("Starting Go API Server...")
  // more cooler print statement
  fmt.Println(`
   ██████╗  ██████╗      █████╗ ██████╗ ██╗
  ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║
  ██║  ███╗██║   ██║    ███████║██████╔╝██║
  ██║   ██║██║   ██║    ██╔══██║██╔═══╝ ██║
  ╚██████╔╝╚██████╔╝    ██║  ██║██║     ██║
   ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝
   `)
  // and then start the server with http package
  err := http.ListenAndServe("localhost:8080", r)  // this takes the base location of your server which in our case is just localhost on port 8080 as well as handler which are muxtype satisfues and then offcourse lets log in our err
  if err != nil{
    logrus.Error(err)
  }

}
