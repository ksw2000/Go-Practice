package main
import "fmt"

//建立一個 method
//https://tour.golang.org/methods/3
type list []int
func (arr list) bubbleSort(){
    for i:=0; i<len(arr)-1; i++{
        for j:=0; j<len(arr)-1-i; j++{
            if arr[j]>arr[j+1]{
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main(){
    fmt.Println("Hello world")
    //調用一個 method
    arr := list{12,32,43,11,54,25,37}
    arr.bubbleSort()
    fmt.Println(arr)
}
