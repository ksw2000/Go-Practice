package main
import "fmt"
func main(){
    //值標記法
    map1 := map[string]string{
        "牛排" : "Angela" ,
        "胡椒" : "阿椒" ,
        "大蛇丸" : "寶拉" ,
    }
    fmt.Println(map1)
    fmt.Printf("%v\n",map1)
    fmt.Printf("%+v\n",map1["牛排"])
    fmt.Printf("%#v\n",map1["牛排"])
    fmt.Printf("%T\n",map1["牛排"])

    //類型標記法
    var map2=make(map[int]string)
    map2[1]="甲"
    map2[2]="乙"
    map2[3]="丙"
    map2[4]="丁"
    map2[5]="戊"
    fmt.Println(map2)
}
