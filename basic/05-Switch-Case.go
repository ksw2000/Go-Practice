package main
import(
    "fmt"
    "math"
    "runtime"
)

func pow(x,n,lim float64) float64{
    if v:=math.Pow(x,n); v<lim{    //判斷前可先做變數設置
        return v
    }
    //we can't use v here
    return lim
}

func main(){
    fmt.Println(
        pow(3, 2, 10),
        pow(3, 3, 20),
    )
    //GO lang 的 switch case 中，不需使用 break;
    fmt.Print("Go runs on ")
    switch os:=runtime.GOOS; os{
    case "darwin":
        fmt.Println("OS X.")
    case "linux":
        fmt.Println("linux.")
    default:
        fmt.Printf("%s.",os)
    }
}
