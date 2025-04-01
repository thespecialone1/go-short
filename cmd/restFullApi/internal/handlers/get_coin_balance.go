package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"github.com/thespecialone1/go-short/cmd/restFullApi/api"
	"github.com/thespecialone1/go-short/cmd/restFullApi/internal/tools"
)

// Here we're assuming the call is already ran through the authorization middleware. So we just needs to grab the username
// from the parameters passed in
// The easiest and most go like way to do this is to decode the parameters into our CoinBalaneParams struct whuch we made earlier

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
  var params = api.CoinBalanceParams{}

  // we can use the gorilla session scheema package for this 
  // calling the NewDecoder function 
  var decoder *schema.Decoder = schema.NewDecoder()
  var err error

  err = decoder.Decode(&params, r.URL.Query())   // so this line just basically gonna grab the parameters in the URL and 
  // set them to the values in the struct 
  // now in this case it'll just grab the ussername from the URL and put it into the username field in the struct 

  if err != nil {
    logrus.Error(err)
    api.InternalErrorHandler(w)
    return
  }
  
  // the rest of the function is very stright forward 
  // so we gonna instantiate the databse interface
  var database *tools.DatabaseInterface
  database, err = tools.NewDatabase()
  if err != nil {
    api.InternalErrorHandler(w)
    return
  }
  
  // Call the GetUserCoins method 
  var tokenDetails *tools.CoinDetails
  tokenDetails = (*database).GetUserCoins(params.Username)
  if tokenDetails == nil {
    logrus.Error(err)
    api.InternalErrorHandler(w)
    return
  }

  // Set the value to our response struct 
  var response = api.CoinBalanceResponse{
    Balance: int((*tokenDetails).Coins),
    Code: http.StatusOK,  
  }
  // and write it to the ResponseWriter
  w.Header().Set("Content-Type", "application/json")
  err = json.NewEncoder(w).Encode(response)
  if err != nil {
    logrus.Error(err)
    api.InternalErrorHandler(w)
    return
  }
    // and thats it! §
}
