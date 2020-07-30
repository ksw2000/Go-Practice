package main
import "fmt"

//一般型式
func add(x,y int) int{
    return x + y
}

//可以為回傳值命名
func Model(x,y int) (result int){
    return x % y
}

func Muti(x,y int) (result int){
    result = x * y
    return
}

//匿名函數 及 閉包試範
func squareSeq() func() int{
    var x int
    return func() int{
        x++
        return x*x
    }
}

//可變參數
func sum(values ...int) int{
    total := 0
    for _,val := range values{
        total+=val
    }
    return total
}

func main(){
    fmt.Println(add(1,2))
    fmt.Println(Model(9,4))
    fmt.Println(Muti(3,4))

    //函數可以做為值 而型態是 domain:int,int codomain:int 2個域都要一樣
    f := add
    fmt.Printf("f := add f(10,10) = %d\n",f(10,10))
    f = Muti
    fmt.Printf("f = Muti f(10,10) = %d\n",f(10,10))

    g := squareSeq()
    fmt.Println("匿名函數範例")
    fmt.Println(g())
    fmt.Println(g())
    fmt.Println(g())
    fmt.Println(g())
    fmt.Println(g())

    fmt.Println("可變參數範例")
    fmt.Printf("sum(8,7) = %d\n", sum(8,7))
    fmt.Printf("sum(1,2,3,4) = %d\n", sum(1,2,3,4))

    //使用可變參數時，也可以用切片作為參數，但要注意應於切片後加入...
    slice := []int{8,9,10,11}
    fmt.Printf("sum(slice) = %d\n", sum(slice...))
}
