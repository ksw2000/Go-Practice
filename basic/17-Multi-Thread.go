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
    go say("world")    // 子執行緒
    say("hello")       // 主執行緒
    // 主執行緒結束時，所有子執行緒也都會隨之結束
}
