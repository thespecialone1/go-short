package tools

import (
	"github.com/sirupsen/logrus"

)

// First we will define the types of data our database should return. These will look like this:

// Database Collection
type LoginDetails struct {
  AuthToken   string      // contains the AuthToken for validating the request
  Username    string
}

type CoinDetails struct {
  Coins      int64        // contains the Coin Balance
  Username  string
}
// these are our LoginDetails and CoinDetails
//----------------------------

// Now the database interface is going to define a few methods that are required for our API 
// We are using an interface here because we can swap out our databases really easily as log as we define a 
// GetUserLoginDetails method, a GetUserCoin method and SetupDatase method with these signature

type DatabaseInterface interface {
  GetUserLoginDetails(username string) *LoginDetails
  GetUserCoins(username string) *CoinDetails
  SetupDatabase() error
}

// Alright now that we deifne our interface and the return types lets create a function called NewDatabase which returns 
// this interface
// Inside the function we are going to create database variable and set it to a mockDB struct. This struct we're gonna create
// which is gonna implement our interface
func NewDatabase() (*DatabaseInterface, error) {
  var database DatabaseInterface = &mockDB{}
  // then we gonna call the setup database (SetupDatase()) method
  var err error = database.SetupDatabase()
  // Do the standard error checks
  if err != nil {
    logrus.Error(err)
    return nil, err
  }
  // and return the pointer
  return &database, nil

}
