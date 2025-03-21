package main

import "fmt"

// Pointers:
// So what are pointers and how to use them in golang
// Pointers are actually special type, these are variables which store memory location [pointer variables =(var p=0x140000b2020)]
// rather than values like intergers [normal boring variables = (var i=123) or floats (var f=1.23)] forexample.

//To create a pointer we use a star(*) syntax
 func main()  {

  // this line(var p *int32) states that variable p will hold the memory addr of int32 value. Initial value of this                  // pointer is nil because it hasn't been initilized yet. Meaning this pointer doesn't point to anything. Yet another               // way to say: This pointer does not have an address assign to it in which to store int32 value yet.
  // As we discussed pointers are types. p still have a memory address himself and and it also stores the value at that address
  // in this case of a pointer this is another memory address

  // lets say: p is at memory location = 0x1b00 it points to this location 0x1b0c
  var p *int32 = new(int32)  //To give this pointer an address we use the built-in function called new(type)
  var i int32                 // a regular int32 variable
  // if we wanna get the value store at this memory location, we can use the * symbol like this (*p). Its called dereferencing the
  // the pointer.
  // when you initilize a pointer with memory location it zeros out the memory location. In other words it sets a value at the 
  // memory location to the zero value of that type. forexample: 0 for int32 and empty string "" for strings and false for boolean
  fmt.Printf("the value p points to is:  %v \n", *p)  // output: the value p points to is:  0 | default for int32
  
  fmt.Printf("the value of i points to is: %v\n", i)  // output: the value of i points to is: 0
  
  //To change the value stored at the memory location of a poimter we use (*p = 10) star notation again
  *p = 10    // this lines states: Set the value at the memory location p is pointing (which we discussed was = 0x1b0c) to 10.
  // Note that * notation does double duty which may be a little confusing. So lets break it down.

  // Here (var p *int32 = new(int32) ) we use it (*) to tell the compiler that we wanna initilize a pointer
  // But on these 2 lines 
  // 1. fmt.Printf("the value p points to is: %v \n", *p)
  // 2. *p = 10
  // we use it (*) to tell the compliler that we want to reference a value of the pointer.
  // So these are two seprate roles that the star syntax has.

  // Another common source of headache is that to get or set the value of a nil pointer. So if we don't call the new function,
  // new function is this that we difned above var p *int32 = new(int32) so if we remove the ('new(type)) and don't call the 
  // new function we don't see any compiler errors. But if we run this then we get this error(nil pointer exception in other 
  // words we get a runtime error):
  // ERROR: -------------------------
  // panic: runtime error: invalid memory address or nil pointer dereference
  // [signal SIGSEGV: segmentation violation code=0x2 addr=0x0 pc=0x104689cf0]

  // goroutine 1 [running]:
  // main.main()
  // --------------------------------
  // Explaination of the error: 
  //         What happening here is that we didnt assign a memory address to a pointer so obviouslt we can't get the value at
  //  (fmt.Printf("the value p points to is:  %v \n", *p)) a memory address (*p) that doesnt exist. 
  // So make sure your pointer is at nil before you assign values to it. 
  
  // Next:
  // we can also create a pointer from the address of another variable using "&" symbol like this: 
  p = &i  // the & symbol here mean is that we want the memory address of the variable not its value. So p now refers to the 
          // memory address of i. In other words both p and i reference the same int32 value in memory.
  // if I use the * notation again to change the value of p the value of i is now also changed. Like this: 
  *p = 1 
  fmt.Printf("the value of p is: %v \n", *p)
  fmt.Printf("the value of i is: %v \n", i)
  
  // This is different from when you're using a regular variable
  // forexample: If you create a new variable K and set i to K so what your program will do here is copy the value of k into
  // i's memory location. Like this:
  var k int32 = 2 
  i = k  
  // the main exception of this copy behavior of non-pointer variables is when working with slices
  // Lets say I copy a slice in the regular way without using a pointer
  var slice = []int32{1,2,3}
  var sliceCopy = slice
  // So now lets try modifying the sliceCopy variable and printing out the values of our two slices again
  sliceCopy[2] = 4
  fmt.Println(slice)         // output: [1 2 4]
  fmt.Println(sliceCopy)     // output: [1 2 4]
  // We can see the values of our orignal slice has also changed. This is because under the hood slices contain pointers to an 
  // underlying array. Basically with slices we're just copying pointers when we do this. So both variables refer to the same
  // data. Remember that while debuggung.


  // Now Pointers and Funtions:
  // these two go really nicely together. Lets see why by looking at the example of func square:
  var thing1 = [5]float64{1,2,3,4,5}
  fmt.Printf("\nThe memory location of thing1 array is: %p", &thing1) //output: The memory location of thing1 array is: 0x1400012a030
  var result [5]float64 = square(&thing1)
  fmt.Printf("\nThe result is: %v", result)  //output: The result is: [1 4 9 16 25]
  // as we can see that their memory addresses are different meaning that these are 2 different arrays. Therefore we can modify 
  // the values of thing2 without affecting the values of thing1 in our main fucntion. Like This:
  fmt.Printf("\nThe value of thing1 is: %v", thing1)
  // but we are also doubling our memory usage of the variables passed in because we are creating a copy for use in our function.
  // so potentially we are using way more memory than we need.
  /*
   WELL NOT ANYMORE!!!!!!!!!!!!!
  */
  // instead lets use pointers we can make our function(square) taking a pointer to an array instead so this:
  // func square(thing2 [5]float64) [5]float64 {  this will changed into -> func square(thing2 [5]float64) [5]float64 {
  // fmt.Printf("\nThe memory location of thing2 array is: %p", &thing2) -> remove & , thing2
  //  return thing2  --> return *thing2
  // in main function we change this:
  // var result [5]float64 = square(thing1)  ---> var result [5]float64 = square(&thing1) - add the & with thing1
  // OUTUT:
  //      The memory location of thing1 array is: 0x140000ba030
  //      The memory location of thing2 array is: 0x140000ba030
  //      The result is: [1 4 9 16 25]
  //      The value of thing1 is: [1 4 9 16 25]%  | NOTE: THE VALUE OF THING1 ALSO CHANGED.
  // Now we can see the memory locations of these two arrays are the same 

  // Pointers are really usefull when passing in large parameters so you don't have to create the copies of the data every time 
  // you call a function wasting time and memory. Instead you can just pass in a pointer to the data. 
  // The only thing to be mindfull of now is that since thing1 and thing2 refer to the same array changing the value of thing2 
  // means the value of thing1 also changed



}

// func square: It take a float64 array of size 5 and squares all the values

func square(thing2 *[5]float64) [5]float64 {        // takes float64 array of 5
  fmt.Printf("\nThe memory location of thing2 array is: %p", thing2) //output: The memory location of thing1 array is: 0x1400012a060
  for i := range thing2{
    thing2[i] = thing2[i] * thing2[i]         // squares all the values
  }
  return *thing2

}
