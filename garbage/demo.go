/*package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	file , err := os.Open("temp.txt")
	if err != nil{
		return
	}
	scan := bufio.NewScanner(file)
	scan.Scan()
	RW := bufio.NewReadWriter(bufio.NewReader(file),bufio.NewWriter(file))
	num , err := RW.Reader.Read([]byte(scan.Text()))
	fmt.Println(num)

	scan.Scan()
	num , err = RW.Reader.Read([]byte(scan.Text()))
	fmt.Println(num)
	if num == 0{
		_ , err = RW.Writer.Write([]byte("/*")) // need to figure out how to get this to write to the file
	}

}

*/

package main

import(
	"fmt"
	"time"
)

func main(){
	var a chan int
	go func(){
		time.Sleep(3 * time.Second)
		a<-0
	}()
	fmt.Println("Waiting...")
	select{
		case <-a:
		fmt.Println("You just waited two seconds")
	}

}
