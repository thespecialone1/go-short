## Things that will be discussed here: 

RestFull API in GO: 
- Golang Code Structure
- API Authorization
- Middleware
- Package Management

We create the project with:

```go mod init <name_of_the_module>```
in my case i'll use my github repo for it: 
```go mod init github.com/thespecialone1/go-short/cmd/restFullApi```

First we start with the structure of our project. Many Go projects are adhere to a standard layout which i will vaguely follow
here: 

NOTE: You don't have to follow this structure for your future projects. Do what works for you:
You can go to this repo and study more about this detailing what code should go in what folder.

1. We gonna ```mkdir api``` here we gonna have specs things like: parameters, response types for our endpoint. This is also where you can put your yaml spec file.

2. Next we gonna have ```mkdir cmd/api``` this will contain our main.go file.

3. Aaand we gonna have our internal folder ```mkdir internal``` this will contain most of the code for this api.

So we start with our api at ```api/api.go```
- we will make a few structs which gonna represent our parameters for our endpoint and responses.
-   first we have a CoinBalanceParams

Now we will define our main package in ```cmd/api/main.go```
- our api will use a chi ```go-chi/chi``` package "its good for we development"
- import our module here ```internal/handlers```
- use logrus with alias log to log errors for debugging ```sirupsen/logrus```
- when we import the packages like chi or logrus we can install them by using ```go mod tidy``` 

Now lets create this handler function  ```  handlers.Handler(r)   ``` in our internal/handlers folder which will setup our router
- we created endpoints at '/account', '/coins'
- Authorization function {which we will make it later}
- GetCoinBalance function {which also will be made later}

Now we got to define our Authorization and GetCoinBalance functions:
1. Lets start with Authorization:
- we will put this in our ```internal/middleware/authorization.go```


Create our database interface in ```internal/tools/database.go```
- create a mock database in ```internal/tools/mockdb.go```

If we go back to our ```internal/handlers/api.go``` api handler we still needs to define GetCoinBalance function 
- in order to use this function in this get method we need to define it such that it takes in a response writer and a pointer to the request as parameters
- So lets do that as our final step in the same Handlers package 
