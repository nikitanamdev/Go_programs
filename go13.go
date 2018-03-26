package main
import (
 "fmt"
 "time"
 )
 func fact(n int) int {
 	if n==0{
 	return 0
    } else {
 	return (n*fact(n-1))
 	} 
 }
 func f2(i int,str string) {
 
   fmt.Println(i,": ",fact(i)%107)
   time.Sleep(time.Millisecond * 100)
 }
 
 func f(str string) int {
   for i:=10000;i<50000;i++{
   go f2(i,str) //f2(i,str)
 }
 return 1
 }

 func main(){
 	fmt.Println("starting")
 go f("direct")

 func (msg string){
  fmt.Println(msg)

 }("going")

 var input string
 fmt.Scanln(&input)
 fmt.Println("done")
 }