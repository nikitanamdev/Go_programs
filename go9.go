package main
import "fmt"

func swap(xptr,yptr *int){
	*xptr,*yptr=*yptr,*xptr
}
func main(){
   x,y:=0,1
   swap(&x,&y)
   fmt.Println(x,y)
}