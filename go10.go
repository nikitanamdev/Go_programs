package main
import "fmt"

func main(){
	xptr:=new(int)
	*xptr=3
   fmt.Println(*xptr)
}