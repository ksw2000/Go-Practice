package main
import "fmt"
type vertex struct{
    x int
    y int
}

type Circle struct{
    vertex // 可以直接繼承
    radius float64
}

func main(){
    v := vertex{1,2}
    v.x=4
    fmt.Println(v.x,v.y)

    var c Circle
    c.x = 0
    c.y = 0
    c.radius = 10
    fmt.Printf("(%d, %d) 半徑：%.1f\n", c.x, c.y, c.radius)
    /*
        Printf format
        %v      預設格式
        %+v     印出結構體時會印出結構體名
        %#v
    */
}
