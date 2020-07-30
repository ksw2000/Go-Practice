package main
import "fmt"
const boilingF=212.0
func main(){
    var f=boilingF
    var c=(f-32)*5/9
    fmt.Printf("boiling point = %g F or %g C\n",f,c)

    var(
        old=15
        str="Yuna我婆"
    )
    fmt.Println(str,"今年",old,"歲")

    const(
        e=2.71828182846 //常數不能使用 := 語法定義。
        pi=3.14159
    )
    fmt.Println("e=",e)
    fmt.Println("pi=",pi)
}
