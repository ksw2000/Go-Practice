package main
import "fmt"
func main(){
    //切片與陣列最大的差別在於切片不必指定長度
    slice1 := []string{"C","C++","GO"}
    fmt.Println(slice1[0:1])
    /*
        切片底層的資料結構
        ptr *Elem   儲存其底層陣列的指標
        int len     切片長度
        int cap     切片容量。一個切片值的容量「不是」其底層陣列的長度

        利用 append(oldslice,element,[elements...]) 擴充slice
        slice的容量(cap)是不能改的，利用append方法可以重新建立slice
        以達擴充效果
    */
    slice2 := append(slice1,"python","Java")
    fmt.Println(slice2)

    /*
        利用 make() 進行切片的宣告
        注意 make() 只能對「切片、字典、通道」三種型態進行宣告
    */
    // 利用 make 宣告 slice
    // 型態(arg[0]):[]int
    // 容量(arg[1]):20
    // 長度(arg[2]):10
    // 若忽略 arg[2] 則視 arg[2] = arg[1]
    slice3 := make([]int, 10, 20);
}
