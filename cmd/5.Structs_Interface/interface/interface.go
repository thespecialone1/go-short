package main

import (
	"fmt"
	"math"
)

// Interfaces: You can think of interfaces like contracts for instances of method for specific type like structs. And these interfaces provides some kind of high level organization without really implementing the business logic of these methods.
// So by just defining the declarations or the definations of these methods that the structs should implement. And this king of flexibility and grouping gives
// us developers a really nice flexibility and high level organization and a struct for instance like i said before can implement these interfaces by just
// embeding or implementig the logic of these methods that were defined in the interface and as long as the struct has these specific methods in the
// interface the struct also directly implements this interface

// What i wrote above is simply this:
// Blueprint or Contract:
// Think of an interface as a blueprint or contract. It doesn't contain any actual code (or business logic); it simply states which methods a type (like a struct) must have.

// Method Declarations Only:
// The interface only declares the method names, the parameters they take, and what they return. It doesn’t tell you how these methods work—that’s up to the struct (or any type) that “implements” the interface.

type shape interface{
  area() float64

}

type measureable interface{
  perimeter() float64
}

type geometery interface{
  shape
  measureable
}

type rectArea struct{
  width, height float64
}

type circArea struct{
  radius float64
}

func (r rectArea) perimeter() float64 {
  return 2 * (r.width + r.height)
}

func (r rectArea) area() float64 {
  return r.width * r.height
}

func (c circArea) perimeter() float64 {
  return 2 * math.Pi * c.radius
}

func (c circArea) area() float64 {
  return math.Pi * c.radius * c.radius
}

func describeShape(name string, g geometery) {
  fmt.Printf("Area of %v: %f\n",name, g.area())
  fmt.Printf("Perimeter of %v: %f\n",name, g.perimeter())
}

//Error in interfaces
type calculationError struct{
  msg string
}

func(ce calculationError) Error() string {
  return ce.msg
}

func performCalculations(val float64) (float64, error) {
  if val < 0 {
    return 0, calculationError{msg: "Invalid Input"}
  }else {
    return math.Sqrt(val), nil
  }
}

func main(){
  rect:= rectArea{width: 5, height: 5}
  describeShape("Rectangle: ",rect)

  circ:= circArea{radius: 3.4}
  describeShape("Circle: ",circ)
  
  calc:= []float64{2,3,-2}
  for _, val:= range calc{
    result, err := performCalculations(val)
    if err!=nil{
      fmt.Println("Error: ", err)
    }else {
      fmt.Println("sqrt result: ", result)
    }
  }
  
  //----------------------------
  //gasEngine Example ----------
  //----------------------------
  var myGasEngine gasEngine = gasEngine{mph: 25, gallons: 15}
  fmt.Printf("MPH are: %v\nKWH are: %v\n",myGasEngine.mph, myGasEngine.gallons)
  fmt.Println("Total miles left are: \n", myGasEngine.milesLeft())
 // canMakeIt(myGasEngine, 50)

  var myElectricEngine eletricEninge = eletricEninge{mpkwh: 40, kwh: 15}
  canMakeIt2(myElectricEngine, 50)
  fmt.Printf("MPKWH are: %v\nKWH are: %v\n", myElectricEngine.mpkwh, myElectricEngine.kwh)
  fmt.Println("Total miles left are: \n", myElectricEngine.milesLeft())
}

//----------------------------------------------------------------------
// Other exaomples of interface using "gasEngine" exaomples
//----------------------------------------------------------------------
type gasEngine struct{
  mph uint8
  gallons uint8
}

// and suppose i have different type of engine: an eletric engine instead of a gasEngine
type eletricEninge struct{
  mpkwh uint8   // instead of miles per hour
  kwh uint8     // instead of gallons we have kwh which specifies how much charge is left in the battery 
}

// The eletricEninge also have a miles left method milesLeft()
func (e eletricEninge) milesLeft() uint8  {
  return e.mpkwh * e.kwh
}

func (e gasEngine) milesLeft() uint8  {
  return e.gallons * e.mph  //we will call that in the main function see under //gasEngine Example in the main function
}

//Now i can pass in this new type of function. So this function takes in gasEngine type parameter and miles parameter and check if it can drive that distance.
// func canMakeIt(e gasEngine, miles uint8){
//   if miles <= e.milesLeft() {
//     fmt.Printf("You can make it there!\n")
//   } else {
//     fmt.Printf("You can't make it!\n")
//   }
// }
//
  //our canMakeIt() function only take gasEngine type. What if we wanna change it take more of a general approach and it should take any engine type. So this
  // where Interfaces comes in.

  //INTERFACES----------------------
  //--------------------------------

  // lets define a interface{} here and see how it can help us here 
  //we use similar syntax that we used to defining a struct except we use the 'interface' keyword
  // we can see in canMakeIt() function that only method we really need here is the milesLeft() method which takes no parameters and return an uint8
  // This is called the method signature. We can specify the signature in the interface like this:
  type engine interface{
    milesLeft() uint8
  }
  //we will write canMakeIt() function again to demostrate
  // so intead of having a very specific requirment that e parameter must be a gas engine we can replace it with our interface. 
  // This function can now take anything for the parameter with only requirment being that the object has a milesLeft method. the signature that were specifies 
  // in our interface. Below is more of a general approach: -- now goto the main function
  func canMakeIt2(e engine, miles uint8){
    if miles <= e.milesLeft() {
      fmt.Printf("You can make it!\n")
    } else {
      fmt.Printf("Need to fuel up first!\n")
    }
  }
  
    













