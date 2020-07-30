package main
import "fmt"
type Vertex struct{
    Lat, Long float64
}

func main(){
    // map在創建時必需以make()創建
    var m = make(map[string]Vertex)

    // 宣告時順便賦值
    var n = map[string]Vertex{
        "Google": Vertex{37.42202, -122.08408},
        "NCHU": { /*Vertex  可省略*/ 24.12281, 120.67616},
    }

    // 用key賦值
    m["Bell Labs"] = Vertex{40.68433, -74.39967}
    fmt.Println(m["Bell Labs"])
    fmt.Println(n["Google"])
    fmt.Println(n["NCHU"])

    /*
        插入或修改
        m[key] = elem
        獲得元素
        elem = m[key]
        刪除元素
        delete(m, key)
        雙賦值檢測某個鍵存在
        elem, ok = m[key]
    */

    //插入
    m["ZSGH"] = Vertex{24.778289, 120.988108}

    elem, ok := m["TFG"]
    if !ok{
        fmt.Println("TFG is not in the map")
    }else{
        fmt.Println("TFG is at", elem)
    }

    elem, ok = m["ZSGH"]
    if !ok{
        fmt.Println("ZSGH is not in the map")
    }else{
        fmt.Println("ZSGH is at", elem)
    }

    map1 := map[string]string{
        "牛排"    :   "Angela",
        "胡椒"    :   "阿椒",
        "大蛇丸"  :   "寶拉",
    }
    fmt.Println(map1)                   // map[大蛇丸:寶拉 牛排:Angela 胡椒:阿椒]
    fmt.Printf("%v\n", map1)            // map[大蛇丸:寶拉 牛排:Angela 胡椒:阿椒]
    fmt.Printf("%+v\n", map1["牛排"])   // Angela
    fmt.Printf("%#v\n", map1["牛排"])   // "Angela"
    fmt.Printf("%T\n", map1["牛排"])    // string
}
