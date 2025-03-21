package main

import (
	"fmt"
	"strings"
)

// We will take a deeper dive into Strings how are they related to Runes and Bytes in golang

// Strings, Runes, Bytes
// -------------------------

func main()  {
  // we will use the word résumé here as a demostration string because it has non-ASCII charectors which will help us understand string in go a bit deeper 
  var myString = "résumé"
  // First we can index a string just like arrays using same notation 
  var indexed = myString[0]
  fmt.Println(indexed)       //  output: 114 -- we get something strange
  fmt.Println(myString)      //  output: résumé

  // lets print out the type of the indexed and as well as the value using %v, %T using printf statement
  fmt.Printf("%v, %T \n",  indexed, indexed)  // output: 114, uint8  -- w=even weirder we get unsigned int8 
  //lets itterate over this string with the range keyword where i is the index and v is the value. We get an output like this: 
  for i, v := range myString{
    fmt.Println(i, v)
  }
  // output: 
            // 0 114
            // 1 233  -- skipped 2. 
            // 3 115
            // 4 117
            // 5 109
            // 6 233
  // To understand this output we need to understand the utf-8 encoding. which is how go represent strings in our computer. 
  // So remember we have to understand "strings"['01110011','01110100','01110010','01101001','01101110','01100111'] as binary numbers in computer. 
  // One such early way of doing it using ASCII encoding. This uses 7-bits to encode 128 charectors. So e.g letter a=97 ASCII charector which is this in binary
  // a-97 = 1100001
  // What do we do if we wanna represent extended set of charectors like empjis, chinese charectors forexample.
  // One way of solving this would be to use more bits. 
  // forexample: we can extend our charectors representation to use instead of ASCII(7-bits) 128 chars to UTF-8(32-bits or 4 Bytes) 1,114,112 chars
  // But: This > UTF-8(32-bits or 4 Bytes) can waste a lot of memory for many charectors.
  // Forexample:
  //            charectors            UTF-32             UTF-8
  //-----------------------------------------------------------------
  //                a               00000000 00000000    01100001   |   Here we can see in UTF-32 we don't need to assisng all these zeros to represent this(a).
  //                -               00000000 01100001               |   UTF-8 on the other hand seeks to solve this issue by allowing variable length in encoding
  //-----------------------------------------------------------------   i.e. using the apropriate number of bytes for the charectors. UTF-8 uses predefined 
  //              <smile-emoji>     00000000 00000001               |   encoding pattern which encodes information about how many bytes is particular charector
  //                -               11110110 00001010               |   uses. 
  //-----------------------------------------------------------------   A charector is uses only 1 byte is it starts with the zero. 2 bytes if it starts with 
  //              <chinese>         00000000 00000000               |   the 110 and so on...
  //                                01011011 10110110               |   The regular ASCII charectors uses 1Byte and anything beyond uses 2 or more.
  // ----------------------------------------------------------------   The acute e (é) charector has a unicode point number of é-233. Unicode charectors 
  //                                                                    numbered between 128 and 2047 and use 2 bytes and hence this pattern [110xxxxx 10xxxxxx]
  //       2-bytes- 110xxxxx 10xxxxxx
  //                    - é-233 - in binary is this:   11101001: We need to pad this number with leading zeros in order for it to fit into UTF-8 encoded 
  //                    - representation like this:    00011101001
  //                110xxxxx + 00011  -> 11000011
  //                10xxxxxx + 101001 -> 10101001
  //                  é = 11000011 10101001 --  So this is now our UTF-8 encoded value for acute e charector. 
  // Lsts jump back into the code to better understand it. 
  // String variable that we declared here: myString = résumé has an underlying array of bytes. Which represents the UTF-8 encoding of the full string which 
  // looks like this: 
  //        r             é                s            u            m                 é
  //  [01110010,  11000011, 10101001,   01110011,    01110101,    01101101,    11000011, 10101001]
  // So when we were indexing at line 17 myString[0] we get 114 (which is the value of "r".)
  // Note that if we index the string at acute e (é) index. we get 195 which is the first halp of the UTF-8 encoding of this charector. 
  var indexedR = myString[1]
  fmt.Printf("%v, %T \n", indexedR, indexedR)  // output: 195 - we wouldn't get the proper 233 we expect for this charector.
  // however when itterrating over this string using the "range" keyword we got the proper 233. 
  for i, v := range myString{
    fmt.Println(i, v)
  }
  // So range keyword is doing the extra work for us here. It knows that the second charector is a 2 byte charector[11000011,10101001] and decodes it correctly 
  // to: 00011101001 = 233
  // This is why we saw index 2 being skipped there. 
  // Take away from this is when you are dealing with strings in go, you're actually dealing with the value whose underlying representation is an array of bytes.
  // This is also why taking the length of the string is its length in the number of Bytes not the number of charectors. explained below: 
  fmt.Printf("\nThe length of 'myString' is: %v\n", len(myString))  // output: The length of 'myString' is: 8
  
  // RUNES:
  // An easier way to deal with iterrating an indexing string is to cast them to an array of runes. Rather than dealing with the underlying byte array of string
   // here's what we get if we do that with our previous string myString
  // []rune(""résumé)
  // [114, 233, 115, 117, 109, 233]
  var runeString = []rune("สวัสดี")
  for i, v := range runeString{            
    fmt.Println(i, v)
  }
  // So runes are just unicode point numbers, which represent the charector. we can individualy now we get the contineous index.
  // RUnes are just alias for int32 as string was for uint8
  
  // ANOTHER WAY OF DECLARING A RUNE:
  fmt.Printf("Another way of declarign a rune is using '' single quotes. Like this: \n")
  fmt.Println("var myRune = 'a' ")
  var myRune = 'a'
  fmt.Printf("\nmyRune = %v", myRune)
  
  // Lets discuss the string building
  // We can concatenate the string 
  var  strSlice = []string{"c", "a", "t"}
  var concateStr = ""
  for i := range strSlice{
    concateStr  += strSlice[i]
  }
  fmt.Printf("\n%v", concateStr)

  // Note that strings are immutable in Go. You cant modify the string after initilizing
  // concateStr[0] = 'a'   - This will not work
  // So concattinating a string a assigning it to a variable like this "concateStr +=" . We are actually crreating a completely new string every time. Which
  // is a very inefficient way of doing it.
  // instead we can import go's built-in package 'strings' and create a string builder
  var strBuilder  strings.Builder
  // instead of using the plus operator we call the WriteString method
  for i := range strSlice{
    strBuilder.WriteString(strSlice[i])
  }
  // and finally call the string method at the end.
  // Whats happening here is that an array is allocated internally and values are appended when calling the WriteString method. Only at the end (the line below)
  // a new string is created from the appended value much faster 
  var catStr = strBuilder.String()
  fmt.Printf("\n%v", catStr)
}
