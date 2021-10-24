---
tags: Golang魔法使
---
# \#13 介面 Interface & 排序套件實戰 sort.Interface | Golang魔法使

![](https://i.imgur.com/aWkpeVJ.jpg)

## 什麼是介面？

我第一次聽到這個詞時以為是要實作出「程式介面」，但實際上完全不是，這個詞完全不是你想的那樣。介面其實就是一種「方法的模板」，介面中提供某個型態要實作「哪些方法」比如一個`介面I`裡面規定要包含「方法A」、「方法B」和「方法C」那麼如果今天有一個型態同時有「方法A」、「方法B」和「方法C」，那麼這個型態就滿足了這個`介面`

舉個例子，比如今天有個介面叫「OO魔法使」，這個介面規定你要有「封印解除」、「收服OO牌」、「利用OO牌」，這三個方法。如果今天創建一個新型態，並且對這個新型態實作這三個方法，那麼我們就會說「這個型態」滿足了「這個介面」

```go=
package main
import "fmt"

// 介面
type geometry interface{
    area()  float64
}

// 自訂型態
type rectangle struct{
    width  float64
    height float64
}

// 新增一個滿足介面的方法
func (r *rectangle) area() float64{
    return r.width*r.height
}

func main(){
    r1 := &rectangle{width:10, height:3}

    fmt.Println(r1.area())
}
```

> 執行結果：
> 30

在範例中我們宣告了一個介面 `geometry` (幾何)，這個介面告訴我們必需實現 `area()`(面積) 這個方法，而這張Golang牌中有一個型態實作了 `area()`，這個型態是 `*rectangle` (不是 `rectangle`，請注意 `第 16 行` 的部分)

但是呢，其實在這個範例中，即使我們把介面那段全部都註解掉，程式還是能正常運作，你可能會很好奇 ── 那麼介面到底是用來幹嘛的？

## 介面的用處？

**觀察以下程式**

```go=
package main
import "fmt"

type rectangle struct{
    width  float64
    height float64
}

func (r *rectangle) area() float64{
    return r.width*r.height
}

type circle struct{
    radius float64
}

func (c *circle) area() float64{
    return c.radius * c.radius * 3.1416
}

//-----------------------------------------//
func (r *rectangle) printInfo(){
    fmt.Printf("%#v\n", r)
    fmt.Println("Its area:", r.area())
}

func (c *circle) printInfo(){
    fmt.Printf("%#v\n", c)
    fmt.Println("Its area:", c.area())
}
//-----------------------------------------//

func main(){
    r1 := &rectangle{width:10, height:3}
    c1 := &circle{radius: 5}

    r1.printInfo()
    c1.printInfo()
}
```

> 執行結果：
> &main.rectangle{width:10, height:3}
> Its area: 30
> &main.circle{radius:5}
> Its area: 78.53999999999999
> 

仔細觀察註解起來的地方，明明兩個方法 `printInfo()` 做的事「一模一樣」但是因為型態不一樣所以仍需要有兩個獨立的方法才能辦到

如果這時我們能把 `*circle` 和 `*rectangle` 都認定為「某個介面」，並以這個介面當做參數的型態會怎麼樣呢？

```go=
package main
import "fmt"

type geometry interface{
    area() float64
}

type rectangle struct{
    width  float64
    height float64
}

func (r *rectangle) area() float64{
    return r.width*r.height
}

type circle struct{
    radius float64
}

func (c *circle) area() float64{
    return c.radius * c.radius * 3.1416
}

//-----------------------------------------//
func printInfo(g geometry){
    fmt.Printf("%#v\n", g)
    fmt.Println("Its area:", g.area())
}
//-----------------------------------------//

func main(){
    r1 := &rectangle{width:10, height:3}
    c1 := &circle{radius: 5}

    printInfo(r1)
    printInfo(c1)
}
```

> 執行結果：
> &main.rectangle{width:10, height:3}
> Its area: 30
> &main.circle{radius:5}
> Its area: 78.53999999999999

因為 `*rectangle` 和 `*circle` 都滿足 `geometry` 所提出的條件 `area() float64` 所以可以將 `*rectancle` 和 `*circle` 視為 `geometry` 這個型態

---

## 中場休息 ── 物理攻擊最對味

這天小櫻遇到「鬥」，她最喜歡跟別人比武，看來是時候使用 **物理攻擊** 了。

![](https://i.imgur.com/zhaz3Gg.jpg)

![](https://i.imgur.com/cBOjsgj.jpg)

![](https://i.imgur.com/CFXmnZY.jpg)

![](https://i.imgur.com/eRKL2bl.png)

---

## 使用Go提供的某些功能時必需先實作介面 ── 以排序為例

究竟什麼時候會需要實作介面呢？舉個例子來說好了，Go已經提供了一個可以幫忙排序的咒語了，但是一些相關的參數必需由我們決定，比如長度、大小、和怎麼把兩個要排序的元素做交換

比如今天要請廠商做衣服，你必需先提供你要的圖案給廠商，廠商才能知道你的要求

要使用排序這個功能，得先使用 `import` 將該套件 `sort` 引入進來，因為同時會使用 `fmt` 所以直觀上是這種寫法：

```go=
import "fmt"
import "sort"
```

但是如果要引用很多套件得一直重打 import 很麻煩所以可以簡單記為

```go=
import(
    "fmt"
    "sort"
)
```

### STEP1: 確認要使用的函式

參考Golang文檔: [sort - The Go Programming Language](https://golang.org/pkg/sort/)

![](https://i.imgur.com/rGTYntd.png)

這次我們確定要使用 `sort.Sort()` 函式，但是裡頭的參數是 `sort.Interface` 型態

裡面寫的英文不難，翻譯如下：

> `Sort` 會排序 `data`。 它會呼叫 `data.Len` 來決定 `n`，並呼叫 `data.Less` 和 `data.Swap` O(n\*log(n)) 次。這個排序為不穩定排序

其中比較難懂的應該是 `O(n*log(n))` 和 `stable` (穩定)

> 1. [時間複雜度 - 維基百科，自由的百科全書](https://zh.wikipedia.org/wiki/%E6%97%B6%E9%97%B4%E5%A4%8D%E6%9D%82%E5%BA%A6)
> 2. [判斷各種排序演算法的穩定性 | 程式前沿](https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/548443/)
> 

### STEP2: 尋找如何實作 Interface

要使用 `sort.Sort()` 前必需先實作 `sort.Interface`

![](https://i.imgur.com/5X8mAXR.png)

> Len(): 要排序的東東的總數
> Less(i, j int) bool: 如果第 i 個東東要在第 j 個東東前面回傳 true, 否則回傳 false
> swap(i, j int): 定義如果把第 i 個東東和第 j 個東東交換
> 

> **語法補充：**
> 1. `Less(i, j int)` 與 `Less(i int, j int)` 意思相同，前者是一種簡寫，先前沒有提及，因此在此補充
> 
> 2. 使用引入的套件前面要加入套件的名稱，比如引入 `sort` 時若要使用 `sort` 裡的函式、常數、型態，必需在這些東西前方加上 `sort.` 。比如 `sort.Sort()`, `sort.Interface`, `fmt.Println()` 
> 
> 3. 如果引用的套件是 `encoding/json` 這時套件名稱為`json` ，總之就是挑「最後的那一個字串」。如果有兩個重複的套件名稱呢？比如 `foo/aaa`, `bar/aaa` ，這時只要加上別稱就行了
> 
> 
> ```go=
> imoprt(
>    a "foo/aaa"
>    b "bar/aaa"
>)
>
>func main(){
>    a.F()
>    b.F()
>}
>```
> 

### STEP3: 排序一個 int slice

**試由小到大排序：**

```go=
list := []int{1, 4, 8, 3, 5, 7, 9, 6}
```

#### 實作
實作時馬上碰到問題，因為我們無法對 `[]int` 添加方法(上一課有提到無法對Golang原生型態新增方法)

試著使用自己的型態，我們對 `[]int` 取別名

```go=
// ...略...
type myList []int

func main(){
    list := []int{1, 4, 8, 3, 5, 7, 9, 6}
    newList := myList(list) // 「轉型」成自定型態
// ...略...
}
```

對自訂義的型態新增三個方法，來滿足 `sort.Interface`

```go=
func (list myList) Len() int{
    return len(list)
}

func (list myList) Less(i, j int) bool{
    return list[i] < list[j]
    // 希望比較小的放前面 (升序)
}

func (list myList) Swap(i, j int){
    list[i], list[j] = list[j], list[i]
    // 注意如果要分開寫要寫成：
    // tmp := list[i]
    // list[i] = list[j]
    // list[j] = tmp
    // 因為 Golang 允許多對多賦值所以可以一行完成
    // 就不會使用傳統的交換方法了
}
```

最後合起來：

```go=
package main
import (
    "fmt"
    "sort"
)

type myList []int

func (list myList) Len() int{
    return len(list)
}

func (list myList) Less(i, j int) bool{
    return list[i] < list[j]
}

func (list myList) Swap(i, j int){
    list[i], list[j] = list[j], list[i]
}

func main(){
    list := []int{1, 4, 8, 3, 5, 7, 9, 6}
    newList := myList(list) // 轉型
    fmt.Println(newList)
    sort.Sort(newList)
    fmt.Println(newList)
}
```

> 執行結果：
> [1 4 8 3 5 7 9 6]
> [1 3 4 5 6 7 8 9]

**如果我想要從大到小排序呢？**

這時只需要更改 Less() 即可
```go=
func (list myList) Less(i, j int) bool{
    return list[j] < list[i]
}
```
替換後：
> 執行結果：
> [1 4 8 3 5 7 9 6]
> [9 8 7 6 5 4 3 1]

## 什麼都可以的空介面 interface{}

如果今天宣告一個什麼都沒有的介面，那麼所有的型態都會滿足 `interface{}` 因此可以利用 `interface{}` 來做為泛型使用。泛型(generic)在魔法科學中用來代表一種「不區分型態」的概念。

```go=
package main
import "fmt"

func myPrint(i interface{}){
    fmt.Println(i)
}

func main(){
    num := 10
    myPrint(num)    // int 滿足 interface{}
    str := "Golang魔法使"
    myPrint(str)    // string 滿足 interface{}
}
```

---

## 後記：

![](https://i.imgur.com/FS4Ji3h.jpg)

20年前的動畫怎麼可以把堂兄妹的愛情也這麼輕描淡寫啊

這部真的神作，後面還有師生戀

幫大家整理一下(第一季20集前的感情狀態關係圖)

> 小櫻 → 雪兔 (差了 7 歲)
> 知世 → 小櫻 (GL)
> 小狼 → 雪兔 (BL)
> 桃矢(小櫻的哥哥) ←→ 雪兔 (BL)
> 利佳 → 寺田老師 (大叔戀 + 師生戀)
> 千春 → 唬爛王山崎 (青梅竹馬)
> 莓鈴 → 小狼 (堂兄妹)
>

![](https://i.imgur.com/vKzWwa4.jpg)

本文多數圖片來自：
[庫洛魔法使第一季第廿集](https://ani.gamer.com.tw/animeVideo.php?sn=10788)
