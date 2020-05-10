package main
import "fmt"
func main(){
    /*
        bool            true,false
        byte            0,1
        rune            Unicode (32bits)
        int/uint
        int8/unit8      8bits
        int16/uint16    16bits
        int32/uint32    32bits (long int in C)
        int64/uint64    64bits (long long int in C)
        float32         32bits (float in C)
        float64         64bits (double in C)
        complex64       float32(real part)+float32(imaginary part)
        complex128      float64(real part)+float64(imaginary part)
        string          Once string is declared, it will be unchangeable
    */
    var a complex64
    var b rune
    a = 1+5i
    b = '\u0022'
    fmt.Println(a)
    fmt.Printf("%c",b)
    /*
        Printf format
        %v      預設格式
        %+v     印出結構體時會印出結構體名
        %#v
    */
}
