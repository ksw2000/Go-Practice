package main
import (
    "fmt"
    "os"
)

// interface 是由一組方法定義的集合
// interface 是一種類型，interface是一組具有多一組方法的類型
// 在 GO 中，並不像 C++ 或 Java 一樣，要由特定的 class 才能實作 interface
// GO 只要類型 T 的方法符合 介面 I 就可以實作

type operate interface{
    BubbleSort() []int
    Reverse()
    Max() int
}

type list struct{
    val []int
}

func (arr *list) BubbleSort() []int{
    for i:=0; i<len(arr.val)-1; i++{
        for j:=0; j<len(arr.val)-1-i; j++{
            if arr.val[j]>arr.val[j+1]{
                arr.val[j], arr.val[j+1] = arr.val[j+1], arr.val[j]
            }
        }
    }
    return arr.val
}

func (arr *list) Reverse(){
    for i:=0; i<len(arr.val)/2; i++{
        arr.val[i], arr.val[len(arr.val)-1-i] = arr.val[len(arr.val)-1-i], arr.val[i]
    }
}

func (arr list) Max() int{
    if(len(arr.val)>0){
        tmp := arr.val[0]
        for _,v := range(arr.val){
            if v > tmp{
                tmp = v
            }
        }
        return tmp
    }
    fmt.Fprintf(os.Stderr,"Empty array\n")
    os.Exit(1)
    return 0
}

type listFloat struct{
    val []float64
}

func (arr *listFloat) BubbleSort(){
    for i:=0; i<len(arr.val)-1; i++{
        for j:=0; j<len(arr.val)-1-i; j++{
            if arr.val[j]>arr.val[j+1]{
                arr.val[j], arr.val[j+1] = arr.val[j+1], arr.val[j]
            }
        }
    }
}

func main(){
    fmt.Println("Hello world")
    var arr list
    arr.val = []int{12,32,43,11,54,25,37,88}
    (&arr).BubbleSort()
    fmt.Println(arr.val)
    (&arr).Reverse()
    fmt.Println(arr.val)
    fmt.Println(arr.Max())

    var arrf listFloat
    arrf.val = []float64{1.2, 8.31, 12.89, 0.98, -12.3}
    (&arrf).BubbleSort()
    fmt.Println(arrf.val)

    //用interface型態
    var l list
    var o operate
    l.val = []int{1,4,9,23,64,13,63}
    o = &l
    fmt.Println(o.BubbleSort())
}
