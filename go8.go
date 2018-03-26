package main
import "fmt"

func main(){
   var fname string
   fmt.Scanf("%s",&fname)
   fmt.Printf("You entered %s\n", &fname)
}