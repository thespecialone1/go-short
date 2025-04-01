package tools

import ("time")

// Now we will create a mockDB type and lets create some data

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
  // here we have some logindetails data as well as some coin details data. These just represent fake data and in the real world 
  // app this probably would just be stored in some form of database
  "alex" : {
    AuthToken: "123ABC",
    Username: "alex",
  },
  "json" : {
    AuthToken: "456DEF",
    Username: "json",
  },
  "marie" : {
    AuthToken: "789GHI",
    Username: "marie",
  },
}


var mockCoinDetails = map[string]CoinDetails{
  "alex" : {
    Coins : 100,
    Username: "alex",
  },
  "json" : {
    Coins : 200,
    Username: "json",
  },
  "marie" : {
    Coins : 175,
    Username: "marie",
  },
}
  // Remeber in order for this mockDB struct to conform to our databse interface we need to create 
  // 1. GetUserLoginDetails, 2. GetUserCoinDetails and 3. SetupDatabse methods
  // So our 2 Get methods just look up the data in the maps that we defined above and setup databse function for our 
  // mock databse jsut does nothing
func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
  // Simulate DB Call
  time.Sleep(time.Second * 1)

  var clientData = LoginDetails{}
  clientData, ok := mockLoginDetails[username]
  if !ok {
    return nil
  }

  return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
  // Simulate DB Call
  time.Sleep(time.Second * 1)

  var clientData = CoinDetails{}
  clientData, ok := mockCoinDetails[username]
  if !ok {
    return nil
  }

  return &clientData
}

func (d *mockDB) SetupDatabase() error {
  return nil
}
