package api

import ("encoding/json"
        "net/http")


// Coin Balance Params
type CoinBalanceParams struct{
  Username string
}


// Coin Balance Response struct
type CoinBalanceResponse struct{
  // Success code, usually 200
  Code int

  // Account Balance
  Balance int
}


// Error Response struct
type Error struct{
  // Error code
  Code int

  // Error Message
  Message string
}

// defining it here because we gonna need it multiple times in our code for error handling
func writeError(w http.ResponseWriter, message string, code int) {  // takes in the writter, the message we want to return and status code 
  // we are gonna use the error struct and write a few things to our response writer 
  resp := Error {
    Code: code,
    Message: message,
  }
  // So here we set the content type. i/e we wanna return the JSON. 
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)

  // Finally we write the error struct out which is gonna be what the user gets back in the response
  json.NewEncoder(w).Encode(resp)
  // Now we are not gonna call this directly in our functions when we wanna return an error because we wanna have a couple of
  // types of error responses
}
// so lets make a few wrappers to this writeError functions

  var (
  // here RequestErrorHandler is going to take in the response writer and the error. We are gonna use this when 
  // we want to return a specific error in our reaponse 
  // this might guide the caller to fix a request. Forexample: 
  // if the username wasn't passed in we might wanna return a message which says username required or something like that 
  RequestErrorHandler = func(w http.ResponseWriter, err error) {
    writeError(w, err.Error(), http.StatusBadRequest)
  }
  // here we'll use the InternalErrorHandler when we want to return a generic error message. 
  // Forexample: If the error is a result of a bug in our code we don't need to return a detailed message to the user 
  // because its not all that helpful for them 
  InternalErrorHandler = func(w http.ResponseWriter) {
    writeError(w, "An Unexpected Error Occurred", http.StatusInternalServerError)
  }
)

