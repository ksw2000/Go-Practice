package main
import (
    "fmt"
)
func main(){
    // 1. 聲明陣列時，方括號内寫明了陣列的長度或者，聲明slice時候，方括號内為空
    // 2. 作為函數參數時，陣列傳遞的是陣列的副本，而slice傳遞的是指針。

    //宣告陣列
    var array [2]string
    array[0] = "是在"
    array[1] = "哈囉"
    fmt.Println(array[0] + array[1])

    array2 := [2]int{94, 87}
    fmt.Printf("%d %d\n", array2[0], array2[1])

    //利用 [...] 來自動宣告不確定的長度
    array3 := [...]int{8, 8, 64}
    fmt.Printf("%d %d %d\n", array3[0], array3[1], array3[2])

    //宣告slice
    N:=[]int{0,1,2,3,4,5,6,7,8,9}
    for i:=0; i<len(N);i++ {    //PHP count() = GO len()
        fmt.Printf("N[%d]=%d\n",i,N[i])
    }
    // [LO:HI] 表示 從 LO ~ HI-1 的值
    fmt.Println("N[0:3]=",N[0:3]) //0:3 [0] [1] [2]
    fmt.Println("N[:5]=",N[:5])   //:5  [0] [1] [2] [3] [4]
    fmt.Println("N[5:]=",N[5:])   //5:  [5] [6] [7] [8] [9] [10]

    // 用 make() 宣告 slice
    // make(型態, 長度 [,容量])
    a := make([]int,5)
	printSlice("a",a)
	b := make([]int,0,5)
	printSlice("b",b)
	c := b[:2]
	printSlice("c",c)
	d := c[2:5]
	printSlice("d",d)

    /*
    len() 可用來返回陣列或slice長度
    cap() 可用來返回陣列或slice容量
    由於陣列長度不可變 len 永遠等於 cap
    另外對於slice，len()返回可見元素個素，而cap()返回所有元素個素
    */

    //nil 表示空slice()
    var z []int
    printSlice("z",z)
    if z==nil {
        fmt.Println("z is nil!")
    }

    /*
    for迴圈遍歷 slice 或 map
    (1)
        for foo,bar := range sliceName
        foo is the position of elements, bar is the contents of elements
    (2)
        for foo := range sliceName
        foo is the position of elements
    (3)
        for _ , bar := range sliceName
        bar is the contents of elements
    */

    var slice = []int{10,11,12,13,14,15,16,17,18,19,20}
    for i,v := range slice {
        fmt.Printf("slice[%d]=%d\n",i,v)
    }

    for i := range slice {
        slice[i]=1<<uint(i)
    }

    for _ , v := range slice {
        fmt.Print(v,",")
    }
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
