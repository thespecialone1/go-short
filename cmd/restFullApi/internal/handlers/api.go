package handlers

import ("github.com/go-chi/chi"
  chimiddle  "github.com/go-chi/chi/middleware"
  "github.com/thespecialone1/go-short/cmd/restFullApi/internal/middleware"
)

// so lets define our handlers function which takes in that mux type we just created 
// NOTE thaat our handler function starts with a capital H. so this just tells the compiler that the function can be imported
// in other packages. If i started with a lowercase and this would have been a private function which could only be used in this 
// handlers package

func Handler(r *chi.Mux) {

  // Now lets cover the concept of middleware. So middleware is essentially a function that gets called before the primary
  // function which handles the endpoint.
  // some examples of how we are gonna use middleware:

  // you can add middleware to a route by using the r.Use() function 
  // So first we gonna use the 'global middleware' this is gonna be the stripslashes function (this function is a pre-built()
  // function, we are grabbing from chi's middleware package

  // GLOBAL MIDDLEWARE: This middlewaremeans its applied all the time. So in other words to all endpoints we make. So this 
  // stripslashed function will make sure that trailing slashel3s (localhost:8008/account/coins-> without trailing slapshes 200 ok)
  // will always be ignored. Otherwise we geta 404 error if we included the slash like this: 
  // with trailing slash (localhost:8080/account/coins/ -> 404 - Not Founf)

  r.Use(chimiddle.StripSlashes)

  // aighr time to set a reoute. To do this we call r.route which takes in a path which is slah account for us as well as anonymou 
  // function which takes ishaker. pie router as parameter. We can use this router to define out get method 
  r.Route("/account", func(router chi.Router) {

  // Lets add another piece of middleware to this route to check if the user id authorized to access this data at first. 
  // this Authorization function is created in our middleware package later. 
  // Now every request which wants to access and endpoint starting with "/account" has to pass through this authorization function
  // first. If the authorization fails the function will return an error as a response and the rest of the code wont get 
  // executed. 
  // So this Authorization() function basically acts like a nightclub bouncer: if you dont have the proper ID you cant come in here

  router.Use(middleware.Authorization)

  // we will create a get method inside this router with slash coins- "/coins" path which will be handled by the GetCoinBalance 
  // function. We will define this function later as well
  router.Get("/coins", GetCoinBalance)
  // So overall we have created an endpoint at '/account' and '/coins'
  })

}
