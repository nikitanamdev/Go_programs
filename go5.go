package main
import "fmt"



func main() {
            var a =[]int{1,2,3,4}; //slice
            var b =[]int{1,2,3,4};
            fmt.Printf("Slice:%v \nArary:%v\n",a,b);
            
           a=append(a,'a','b','c')
           fmt.Printf("Append to Slice?:%v\n",a)
            
            
}