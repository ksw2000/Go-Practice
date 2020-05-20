package main
import "fmt"

func sender1(c chan string){
    c <- "sender1"
}

func sender2(c chan string){
    c <- "sender2"
}

func main(){
    //https://colobu.com/2016/04/14/Golang-Channels/
    ch1 := make(chan string)
    ch2 := make(chan string)

    //select
    go sender1(ch1)
    go sender2(ch2)

    s,t := <- ch1, <- ch2
    //s,t 接收到 channel 後才會往下走

    fmt.Println(s, t)

    //利用 range 接收不確定數量的 channel
    n := make(chan int)
    go func(end int, n chan int){
        for i:=0; i<end; i++{
            n <- i
        }
        close(n)
    }(10,n)

    fmt.Println("接收")
    //range 會一直接收 n 直到 n 被關閉為止
    for v := range n{
        fmt.Printf("%d ",v)
    }

    //Buffer
    //https://blog.wu-boy.com/2019/04/understand-unbuffered-vs-buffered-channel-in-five-minutes/
    //unbuffered channel 就是代表在主程式內，需要等到讀或寫都完成
    //main 才可以完整結束 (讀跟寫 buffered channel 需要在不同的 goroutine 才不會被 block)

    //buffered channel 相反，你可以一直丟資料進去 Channel 內
    //不需要讀出來
    //（前提是 buffered channel 空間夠大不會爆掉）
    //所以 main 才提前結束。
}
