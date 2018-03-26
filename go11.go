package main
import (
    "fmt"
    "bufio"
    "os"
)    

func main() {
	var test int=3;
    reader := bufio.NewReader(os.Stdin)
    name,_ := reader.ReadString('\n')
    var count =[5]int 
    a:=1
    b:=2
    for i:=0;i<3;i++ {
    	for i:=0;i<len(name);i{
    		if name[i]=='f'{
    		   count[0]++;
    		} else if name[i]=='o'{
    			count[1]++;
    		} else if name[i]=='b' {
    			count[2]++;
    		} else if name[i]=='a' {
    			count[3]++;
    		} else if name[i]=='r' {
    			count[4]++;
    		}
    	}
    	for i:=0;i<5;i++ {
    		if count[i]==0{
    			fmt.Println(0)
    			break
    		} ekse {
    			
    		}
    	}

    	 
    }


}