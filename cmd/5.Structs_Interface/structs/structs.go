// Structs and Interfaces
package main

import "fmt"

// Structs: Structs are basically defining your own types. To create a struct we can use the 'type' keyword. As we are defining a type here, name of our type
// followed by the struct keyword and curly braces
// structs can hold mix types in the form of fields which we can define by name

type budgetExpense struct{
  income uint8
  salary uint8
  expense uint8
  ownerInfo owner    // we create an ownerInfor field here of type owner and it can be accessed using ownerInfo.name
  //we also have another option which is to directly put the owner type without ownerInfor like this:
  //owner
  //we can also just type int our type and now we have a field called int which is also of type int 
  //int 
}
//Fields can be anything yu want even another struct
// a struct of type owner who owns these assets 
type owner struct{
  name string
}

func main(){
  // we can create a new variable like this:
  var myBudget budgetExpense
  fmt.Println(myBudget.income, myBudget.salary)// output: 0 0 because we haven't defined our income and salary fields yet. So its a zero value
                                               //          struct meaning default values are set for the fields.
  // One way to set these values is to use the struct literal syntax like this: 
  fmt.Printf("To assign values: \nvar myBudget2 budgetExpense = budgetExpense{income: 55, salary: 50}\n")
  // another way is remove the field names and then values will be assigned in order
  fmt.Println("Another way to assign values in order: \nvar myBudget2 budgetExpense = budgetExpense{55, 50}")

  var myBudget2 budgetExpense = budgetExpense{income: 55, salary: 50, ownerInfo: owner{"Alex"}, expense: 10}
  var myBudget3 budgetExpense = budgetExpense{85, 75, 20, owner{"Adam"}} // here we are using the int field which is also of type int
  fmt.Printf("My Income: %v, My Salary: %v, Owner: %v , Expanses: %v\n",myBudget2.income, myBudget2.salary, myBudget2.ownerInfo.name, myBudget2.expense) // so after putting the only type 'owner' 
  // in the budgetExpense struct we can omit the myBudget.ownerInfo.name to > myBudget.name(syntax) we can add the sub field directly and use the .name syntax
  // Actually we can do this with any type 
  
  fmt.Printf("My Income: %v, My Salary: %v, Owner: %v, Expanses: %v\n",myBudget3.income, myBudget3.salary, myBudget3.ownerInfo.name, myBudget3.expense)
  
  // we can set the values directly like this:
  myBudget.income = 20
  myBudget.salary = 15
  
  // ANONYMOUS Structs; Here we are not defining the name type like we did above. With anonymous struct we have to define and initilize it in the same location 
  // main reason is that this is not resuable. forexample if i wanna create another struct then i have to re-write the defination
  var anonStruct = struct {
    income uint
    salary uint
  }{25, 30}

  fmt.Println(anonStruct.salary, anonStruct.income)

  // Method:
  fmt.Printf("Total profit made by %v is:  %v\n", myBudget2.ownerInfo.name, myBudget2.profit())
  fmt.Printf("Total Printf made by %v is:  %v\n",   myBudget3.ownerInfo.name, myBudget3.profit())

  whoMadeMore(myBudget2, myBudget3)
}

  // Structs also have a concept of method which we can use aswell. These are functions which are directly tied to the struct and have access to the struct 
  // instance itself
  // lets say we wanna make a method which calculates the profit of the owner
  func (b budgetExpense) profit() uint8 {   // except for this (b budgetExpense) part here. A method is just like a function. In this case we would return an 
                                            //  uint8. now what we were doing in this (b budgetExpense) first part here was we were assigning this profit() 
                                              // function to the budgetExpense type. This (profit()) function or method now has access to the fields and even 
                                            // other methods that were signed to this budgetExpense type. So now we go to our main function
    return b.income-b.expense
  }
  
  func whoMadeMore(b1, b2 budgetExpense)  {
    profit1 := b1.profit()
    profit2 := b2.profit()
    
  if profit1 > profit2 {
    fmt.Printf("%v has made more profit (%v vs %v)\n", b1.ownerInfo.name, profit1, profit2)
  }else if profit2 > profit1{
      fmt.Printf("%v has made more profit (%v vs %v)\n", b2.ownerInfo.name, profit2, profit1)
    }else{
      fmt.Printf("Both %v and %v has made same profits", b1.ownerInfo.name, b2.ownerInfo.name)
  }
}







