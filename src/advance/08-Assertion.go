package main
import "fmt"

func funcName(a interface{}) {
    value, ok := a.(string)
    if !ok {
        fmt.Println("It's not ok for type string")
        return
    }
    fmt.Println("The value is ", value)
}

func main(){
    /*
        interface 可以作為泛型使用
        在做為泛型時用時若需要確定其真實型別時
        可以用斷言的方法來確定
    */
    var t interface{}
    t = "test"
    funcName(t)

    var u interface{}
    u = 20
    funcName(u)

    var v interface{}
    v = true
    switch v := v.(type) {
        default:
            fmt.Printf("unexpected type %T", v)       // %T prints whatever type t has
        case bool:
            fmt.Printf("boolean %t\n", v)             // t has type bool
        case int:
            fmt.Printf("integer %d\n", v)             // t has type int
        case *bool:
            fmt.Printf("pointer to boolean %t\n", *v) // t has type *bool
        case *int:
            fmt.Printf("pointer to integer %d\n", *v) // t has type *int
    }
}
