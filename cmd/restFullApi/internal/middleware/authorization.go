package middleware

import (
	"errors"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/thespecialone1/go-short/cmd/restFullApi/internal/tools"
	"github.com/thespecialone1/go-short/cmd/restFullApi/api"
)

// First lets create a custome unauthorized error which we're gonna use in this function

var UnAuthorizedError = errors.New("Invalid username or token.")

// So because we are using the authorization function as middleware this needs to follow a certain signature. i.e.
// it need to take in and return an http handler interface 
func Authorization(next http.Handler) http.Handler {
  // now we can make this function return in HTTP Handler like this: we are using this HandlerFunc in the http package
  // this takes in a function which itself takes in a response writer and a pointer to the request so the response writter 
  // is what you use to contruct a response to the caller. e.g. Set the response body, the headers and the status code.
  // The request is what contains all the information about the incoming request. Forexample: headers, payloads and other
  // information about the http request which came in.

  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// Now in here we will define our logic for authorizing the HTTP request
    // grab the username parameter from the request pointer by calling r.URL.Query.Get("name of the parameter")
    var username string = r.URL.Query().Get("username")

    // grab the auth token from the header like this
    var token = r.Header.Get("Authorization")
    // now if either of these are empty we can return an error 

    if username == "" || token == ""{
      // to do this lets go back to our api in api.go and we can create a function which we can use to write an 
      // error response to the HTTP response writer in other words this function is gonna return an error response
      // to the person who called the endpoint. 
      // Point of making a function there in api/api.go is because we're gonna reuse it in multiple places in our code
      // when we wanna return an error

      // Now we can call the error to the console and then call our new request error handler passing unauthorized error 
      // and then exiting the function with return 
      logrus.Error(UnAuthorizedError) 
      api.RequestErrorHandler(w, UnAuthorizedError)
      return
    }
    // if we have both of these things we can proceed to getting the data from our database and checking the username and 
    // Authorization token is correct. 
    // we are gonna isntatiate a pointer to our databse like this using an interface type and call the new database method
    // *dont worry we are going to define all these functions types later 
    var database *tools.DatabaseInterface
    database, err := tools.NewDatabase()
    // if we get an error back we'll return an internal error in this case
    if err != nil{
      api.InternalErrorHandler(w)
      return
    }
    
    // Now lets actually querry the databse using the "GetUserLoginDetails" method
    var loginDetails *tools.LoginDetails
    loginDetails = (*database).GetUserLoginDetails(username)

    // Finally if we didn't find a client with a username or the token doesn't match what we got back from our database then again
    // we're gonna return a request error. End of this function we need to call the next.serveHTTP method(explaination below)
    if loginDetails == nil || (token != (*loginDetails).AuthToken){
      logrus.Error(UnAuthorizedError)
      api.RequestErrorHandler(w, UnAuthorizedError)
      return
    }
    // what this next.ServeHTTP method does?
    // It calls the next middleware in line (Middleware --> next.ServeHTTP --> Middleware2 --> ... --> HandlerFunc) or the Handler
    // Function for the endpoint if  there's no middleware left 

    // So in our case this will call our GetCoinBalance Function at the end like this: 
    // Authorization --> next.ServeHTTP --> GetCoinBalance
    next.ServeHTTP(w, r)
  })
}
