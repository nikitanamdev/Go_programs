package main
import (
"bufio"
"fmt"
"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	name,_ := reader.ReadString('\n') //_ is a place holder variable
	fmt.Println("You entered "+name)
}