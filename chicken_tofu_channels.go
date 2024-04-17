package main
import (
	"fmt"
	"math/rand"
	// "sync"
	"time"
)

var thresholdPrice float32 = 10

func main(){
	chicken_buffer := make(chan string)
	tofu_buffer := make(chan string)
	websites := []string{"hellofresh.com","walmart.com","target.com"}
	for website := range websites{
		go dealOnChicken(chicken_buffer, websites[website])
		go dealOnTofu(tofu_buffer, websites[website])
	}
	sendMessage(chicken_buffer, tofu_buffer)
}


func dealOnTofu(t chan string, web string){
	// defer close(t)
	for{
		time.Sleep(time.Second*1)
		price:= rand.Float32()*20
		if price < thresholdPrice{
			t<-web
			break
		}
	}
}

func dealOnChicken(c chan string, web string){
	// defer close(c)
	for{
		time.Sleep(time.Second*1)
		price := rand.Float32()*20
		if price <= thresholdPrice{
			c<-web
			break
		}
	}


}



func sendMessage(c chan string, t chan string){
	// select i slike an if statement for channels
	select{
	case website:=<-c:
		fmt.Printf("the was a deal on chicken in %v\n", website)
	
	case website:=<-t:
		fmt.Printf("the was a deal on tofu in %v\n", website)

}
}