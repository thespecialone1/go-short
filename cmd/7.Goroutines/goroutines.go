package main

import (
	"fmt"
	"sync"
	"time"
)

//Goroutines:  They are basically are a way to launch multiple functions and have them execute them concurently.
// NOTE: concoruncy != parallel execution

// Concurency means i have multiple tasks running at the same time. One way i can do is jumping back and forth working on one task to another.
// lets say: task1 involves database call which takes 3 seconds to return the data. In concurent programming while i'm waiting for the data to respond
// the cpu can move on to working on the task 2 in the mean time. Then when i get a response may i move back to finish up the task 1 and then maybe move back
// again to finish up task 2.
// Other way to achieve concurency:
// Another way to achieve concoruncy is through parallel execution. Instead of having 1 cpu core working on these tasks, i can have 2 cores then execution can
// happen simultaneously. So the execution here is still concurent because we have multiple tasks in progress at the same time but also these tasks are running
// in parallel. So we can conclude that program maybe runnng concurently but may not neccesarily be executing tasks in parallel.

// In practice tho you can use some level of parallelism as long as you have multi core cpu. Which you probably do.

// lets code:
// var m = sync.Mutex{}   
var m = sync.RWMutex{}   // this has all the same functionality that we saw in before and lock() and unclock() methodds work exactly the same. But now we also have RLock() and RUnlock().
var waitGroup = sync.WaitGroup{} // waitgroups are simple they are pretty much just counters.
// When ever we spawn a new go routine we make sure to add '1' to the counter like this := wg(variable).Add(1) and then inside the function we call the done method at the endd like this '(variable).done()'

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{} // to store results from databse

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// When ever we spawn a new go routine we make sure to add '1' to the counter like this :=
		waitGroup.Add(1)
		go dbCall(i)
		// after waitGroup.Add(1) and then the waitGroup.done() we will then call the wait method. waitGroup.wait()
	}
	waitGroup.Wait() // this method is gonna wait for the counter to go back down to zero (meaning that all the tasks have completed and the rest of the code will execute).
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe Results are:%v", results)
}

// this function simulates a call from db retreiving an id from the list above [var dbData = []string id1, id2 ...
// it takes random amount of time [rand.Float32() * 2000] upto 2 seconds per call and we call this mock database [for i := 0; i < len(dbData); i++] 5 times
// Now we can call this database swquentially one by one which and take an average of 5 seconds to complete
// A better way to do this is to let these database calls run concurently.
// To do this we use the 'go' keyword infront of the function [go dbCall(i)] you wanna run concurently. Now our program wont wait for this function to complete
// rather it'll keep moving on to the next step in the loop.  {if we run this code now then nothing will happen}
// Reason: So our program spawns these tasks int he background didn't wait for them to finish and then exited the program before they were complete

// hmmmmmmmmmmmmmmmmm.....

// So we need a way for our program to wait until all these tasks have been complete and then contineu on to the rest of the code.
// This is where WAITGROUPS comes in:
// waitgroups can be imported through the sync package 'sync'. we can then create a waitgroup like this: <- see below lets code:

// Mutex:
// So what if instead of printing out the results of the console, we wanted to return them to the main function. First off lets make a slice
func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from Database is: ", dbData[i])
	// After this := wg(variable).Add(1) then inside the function we call the done method at the endd like this 'done()' this decrements the counter. And then in main we call the wait method
	// lets append the value we get back from fake database

	// m.Lock()// what happens here is that when a goroutine reaches this lock method a check is performed to see if a lock
	// is already been set by another go routine. If it has it will wait here until the lock is released and then set the lock itself. When its done tinkering
	// with results array, Lock is released again m.Unlock() with the unlock method and then other goroutines can obtain a lock is needed
	// It really matters where you put your lock in and unlock statements. e.g. If i put the lock() above time.sleep
	// it will work but then it would completely destroy the concurency of our project
	// results = append(results, dbData[i])
	// m.Unlock()
	// one drawback of this sort of mutex is that it completely locks out other gorooutines to accessing the results slice.  It can be usefull but it might not.
	// Go provides another type of mutex called read writed mutex: RWMutex
	// this has all the same functionality that we saw in before and lock() and unclock() methodds work exactly the same. But now we also have RLock() and RUnlock()
	save(dbData[i])
	log()
	waitGroup.Done()
}

// If we have multiple threads modifying the same memory location at the same time, we can get some unexpected reults. (Missing data)
// Forexample 2 processes writing to the same memory location at the same time could lead to corrupt memmory as well as whole host of other issues
// So we can't run this code [i mean we can but we really shouldn't]. Inside we can use what are called the mutex.
// to control the writing to our slice in a way that makes it safe in a concurent program like ours
// Mutex:
// Now if instead of printing out the results on the console we want them to return to the main function: For that we need to make a slice where we can store all
// the results from the database.
// How to create Mutex: var m = sync.mutex{} can create it like this using the sync package
// Mutex is short for mutual exclusion.
// Mutex Methods:
// 1. m.Lock() and 2. m.Unlock() -- and we will place them around the part of our code which accesses the result slice
// Another type of mutex is: Read Write Mutex:= RWMutex{}

// For improving our code lets make few seprate functions here to read and write from the result slice
func save(reult string)  {
	m.Lock()
	results = append(results, reult)
	m.Unlock()
}

// The part here we read the result we can add a read lock and read unlock
func log(){
	m.RLock()	// so when a goroutine reaches this line, it checks if there's a full lockon the mutex, by full lock meansthe lock type we saw before where we call the lock method
				// if this full loack exists it'll wait untill its released before contineuing. This way we are not reading while the results are potentionally being written to.
				// if no full lock is exists, the goroutine will acquire a readlock and then proceed with the rest of the code. 
// NOTE that many goroutines may hold read locks at the same time. These read locks will only block code execution up here[m.Lock() in save func]
// when a go routine hits it(m.Lock())-- In order to proceed all locks must be released and that is full locks and read locks. 
// This prevents us from accessing the slice while other goroutines are writing to or reading from the slice. Just like before.
// So in SUMMARY this patteren allows multiple goroutines to read from our slice at the same time only blocking when writes may potentionally be happening
// This is in contrast to what we saw before with the Lock() and Unlock() methods with a regular mutex. In that case even read of the data can only happen one at a time 

	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}

// Peroformance Improvements: 


