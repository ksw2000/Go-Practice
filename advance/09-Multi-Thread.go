package main
import(
    "fmt"
    "time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")    //子進程
	say("hello")       //主進程
    /*主進程結束時，所有子進程也都會隨之結束*/
}
