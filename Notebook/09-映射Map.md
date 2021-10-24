---
tags: Golang魔法使
---
# \#9 映射 Map 雪兔：除非是有魔法的人否則根本辦不到的 | Golang魔法使

![](https://i.imgur.com/VZZHhw8.jpg)

這次小櫻遇到一個非常兇猛的Golang牌，究竟是什麼Golang牌可以把企鵝大王給弄倒呢？

我們先前提到的陣列與切片都是以整數為鍵，而這次要介紹的複合型態，可以「不以整數為鍵」，而這麼兇猛的卡牌即為「Map」

這個功能其實是可以用 Array 加一些黑科技(Hash)實作出來，但因為自己實作過於麻煩而且該功能經常常用，所以Golang直接對Map做原生支援(相較C++, Java，Go在編譯器層面即支援 Map)

## 宣告一個 Map

先不管 Map 後面的原理到底是什麼，我們直接宣告一個來試試

```go=
package main
import "fmt"
func main(){
    // 建立一個由 string 映射到 int 的 map
    tall := make(map[string]int)
    tall["小櫻"] = 153
    tall["知世"] = 155
    tall["小狼"] = 156
    fmt.Println(tall["小櫻"])
    fmt.Println(tall["知世"])
    fmt.Println(tall["小狼"])
}
```

> 執行結果：
> 153
> 155
> 156
> 

如此一來便可以很快地建立一個以字串為鍵(key)的陣列

長得很像陣列，用起來也很像陣列，但是在Golang中稱為Map，python中稱為Dictionary(字典)、Javascript中稱為Object(物件)、PHP中稱為Associative array(關聯陣列)，總之又是一個各家自己看心情命名的複合結構。

## 更簡短的宣告

然而，剛剛所用的方法是否太過麻煩！因此有更簡短的宣告方法

```go=
package main
import "fmt"
func main(){
    // 更簡短地建立一個由 string 映射到 int 的 map
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,    // 注意這個逗號要記得加
    }
    fmt.Println(tall["小櫻"])
}
```

> 執行結果：
> 153

## 檢查該鍵是否存在

在存取時想確定該鍵是否存在該怎麼做呢？

其實可以利用兩個回傳值去存取，第一個會回傳「對應的值」，第二個會回傳「布林值」用來表示該鍵是否存在

```go=
package main
import "fmt"
func main(){
    // 更快地建立一個由 string 映射到 int 的 map
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,
    }
    val, ok := tall["小櫻"]
    fmt.Println(val, ok)
    val, ok  = tall["小可"]
    fmt.Println(val, ok)
}
```

> 執行結果：
> 153 true
> 0 false

如果該鍵並不存在，則在讀取該值時其實**並不會出錯**，而是會以**預設值充當對應值回傳**

如果只是想確認該值是否存在而不想知道該值，則可以在回傳對應值的部分設為 `_`，以免魔仗因為你沒使用宣告出來的變數而噴錯

```go=
package main
import "fmt"
func main(){
    // 更快地建立一個由 string 映射到 int 的 map
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,
    }
    _, ok := tall["小狼"]
    fmt.Printf("tall[\"小狼\"] 存在嗎？%t", ok) // 第 11 行
}
```

> 執行結果：
> tall\["小狼"\] 存在嗎？true

> **補充一點**：在第 11 行 print 的地方當我們要 print 出雙引號時，為了必免被誤認成字串用的雙引號，我們可以在雙引號前加上反斜線 `\` 去告訴編譯器這是一般的字不是識別字串起頭結尾的字

## 走訪

如同前兩天所教的課程(Array, Slice)，Map 也可以透過 `for + range` 的方式進行走訪

```go=
package main
import "fmt"
func main(){
    // 更快地建立一個由 string 映射到 int 的 map
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,
    }
    for k, v := range tall{
        fmt.Printf("%s的身高是：%dcm\n", k, v)
    }
}
```

> 執行結果：
> 小櫻的身高是：153cm
> 知世的身高是：155cm
> 小狼的身高是：156cm

> **注意**：走訪時的順序跟加入的順序無關
> 

## For迴圈走訪補充
(2020/9/13補充)
利用 for 迴圈走訪時，通常會用 `key, value` 去接 range 的回傳值，但是，其實也可以只用一個參數，而這個參數預設是 `key`，所以如果只想走訪 `key` 時就可以不用使用 `k, _ := range xxx` 這種寫法，可以直接寫成 `k := range xxx`

```go=
package main
import "fmt"
func main(){
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,
    }

    for k := range tall{
        fmt.Println(k)
    }
}
```

> 執行結果：
> 小櫻
> 知世
> 小狼

## 刪除一對鍵值

除了可以新增和修改 map 外，也可以刪除一對鍵值，做法是使用 `delete()`

```go=
package main
import "fmt"
func main(){
    // 更快地建立一個由 string 映射到 int 的 map
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,
    }
    delete(tall, "小櫻")  // 將小櫻從 map 中刪除
    for k, v := range tall{
        fmt.Printf("%s的身高是：%dcm\n", k, v)
    }
}
```

> 執行結果：
> 知世的身高是：155cm
> 小狼的身高是：156cm

## 後記

給大家看一下唬爛王山崎↓

![](https://i.imgur.com/2P54Smo.jpg)

而且整部動漫只有李小狼和小櫻會信，笑死，難道是因為同為魔法使的關係，吔\~不對，李小狼是魔法使？

**邦友的按讚是我發文的動力** ，希望各位邦友動動手按個讚！如果有不懂的地方也可以在本文底下留言互動哦！

想了解Map的運作原理可以參考
> 1. [How the Go runtime implements maps efficiently (without generics)](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics)
> 2. [Golang官方 map.go](https://golang.org/src/runtime/map.go)
> 

<!--如果想稍微了解的可以搜尋 Hash map，hash 中文稱雜湊，可以將不同字串轉成不同數字(理想上)，利用這個特性再把字串映射到 Array 上，最後為了避免碰撞，再透過 linked list 或紅黑樹等神祕資料結構去儲存。但是這一部份也許以後也不會提到。如果今天主題是 C 語言，就沒有原生的 Map 可以使用；如果是 java 則是額外包成一個物件來用；rust 也是用額外的插件去支持。

Golang 則是直接做原生支持，簡單方便又好用，雖然我覺得 rust 的 hashmap 真的是快到炸掉就是了。-->

接下來的課程會越來越難，這系列文進度會越來越快，請新來的魔法使做好心理準備。為了讓新來的魔法使能跟上進度就來出個題目好了

### 練習1
> 利用迴圈印出九九乘法表
> 提示：for迴圈內可以再加for迴圈 
> ![](https://i.imgur.com/pljfrUA.png)
> [#5 For 迴圈流程控制 | Golang魔法使](https://ithelp.ithome.com.tw/articles/10234260)


### 練習2
> 不使用現有函式庫提供的排序函式下，由小到到大排序一個整數切片
> ```go=
> arr := []int{10, 4, 5, 3, 7, 1, 9, 8}
> ```
> [#7 陣列 Array | Golang魔法使](https://ithelp.ithome.com.tw/articles/10234744)

本文多數圖片來自：
[庫洛魔法使第一季第十三集](https://ani.gamer.com.tw/animeVideo.php?sn=10781)

---

## 練習題參考解答

### 練習1

```go=
for i := 2; i < 10; i++ {
    for j := 2; j < 10; j++ {
        fmt.Printf("%dx%d = %2d ", j, i, i*j)
    }
    fmt.Printf("\n")
}
```

### 練習2
```go=
func insertionSort(list []int) []int {
    for i := 1; i < len(list); i++ {
        for j := i; j > 0; j-- {
            if list[j] < list[j-1] {
                list[j], list[j-1] = list[j-1], list[j]
            }
        }
    }    
    return list
}
```
