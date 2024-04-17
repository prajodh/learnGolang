package main
 
import ("fmt"
		"time")

func main(){

var c = make(chan int)
go process(c)
for i := range c{
		fmt.Println(i)

 }

 var c_buffer = make(chan int, 6)
 go process_buffer(c_buffer)
 for i:= range c_buffer{
	fmt.Println(i)
	time.Sleep(time.Second*1)
 }
 fmt.Println("the main function just finished")
}

 func process(c chan int){
	defer close(c)
	for i := 0; i <5 ; i++{
		c<-i
}
 }

func process_buffer(c chan int){
	defer close(c)
	for i:=0; i < 5;i++{
		c<-i
	}
	fmt.Println("Done with the channels")

}
