package main

import "fmt"

func main() {
  var counter int
  var grade int
  var total int

  var average int

  fmt.Println("Enter grade, -1 to end.")
  
  fmt.Scanln(&grade)
  
  for grade != -1 {
    
    total = total + grade
    
    counter = counter + 1
   
    fmt.Println("Enter grade, -1 to end.")
    
    fmt.Scanln(&grade)

  }
     average = total / counter
     
     fmt.Println(average)  
    
  
}