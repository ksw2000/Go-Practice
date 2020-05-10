package main
import "fmt"
func fibonacci() func() int {
    back1, back2 := 0, 1
    return func() int {
        temp := back1
        back1, back2 = back2, (back1 + back2)
        return temp
    }
}

//another example
              // 回傳型態: [func() int]
func intSeq() func() int{
    i:=0
    return func() int{
        i+=1
        return i
    }
}

func main(){
    //demo fibonacci()
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f()," ")
	}
    //demo intSeq
    fmt.Println()

    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println("更新")

    newNextInt := intSeq()
    fmt.Println(newNextInt())
}
