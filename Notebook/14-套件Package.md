---
tags: Golang魔法使
---
# \#14 套件 Package (2021) & 套件實戰(math, http, sort) | Golang魔法使

![](https://i.imgur.com/ttN38C9.jpg)

我懶的弄故事劇情了，反正今天就是要教套件啦

## Go 內建套件
Go 內建的套件，在安裝時就已經附了。其中最常用的就是套件 `fmt` ，`fmt` 是 format 的簡稱主要是處理格式化相關的工作，比如常見的 `fmt.Print()` `fmt.Println()`

如果要看官方的文件說明可以直接來這裡找 [https://golang.org/pkg/](https://golang.org/pkg/)

至於範例的部份，因為 [昨天的課程](https://ithelp.ithome.com.tw/articles/10236366) 已經示範使用 `package sort` 了，所以今天課程著重在教大家自己製作 package

## 自訂套件 (一) 使用 GOPATH (傳統方法)

### STEP1 架設環境

自訂套件的第一種方法是使用環境變數 `GOPATH` 來定義套件的路徑

在第一天 [環境架設教學](https://ithelp.ithome.com.tw/articles/10233333) 時有教大家設置 `GOROOT` 了，那麼什麼是 `GOPATH` 呢？

首先，請各位魔法使們開命令提示字元並輸入

```shell=
go env
```

![](https://i.imgur.com/YhD88AZ.png)

其中注意 GOPATH 那行，因為我已經有設定了，所以我的 GOPATH有值，但可能有些魔法使們還沒有。

簡單來說，只要把現在工作的目錄加進 `GOPATH` 中就行了，比如我現在在 `C:\Users\liao2\OneDrive\go-tutorial` 工作，這裡面存有大量練習的 go 的檔案

> 像這樣：
> + go-tutorial/
>    + lesson1.go
>    + lesson2.go
>    + lesson3.go
>    

**如果畫面僅有一個欄位可以填，就在現有變數值的路徑結尾加上 `;` 後，再填寫就可以了！**

簡單來說就是可以 `路徑1;路徑2;路徑3`

把 `C:\Users\liao2\OneDrive\go-tutorial` 加進 `GOPATH` 

![](https://i.imgur.com/re1uucS.png)

### STEP2 重新開機

### STEP3 建立正確的結構

在開始使用之前我們必需重新整理我們的檔案結構

> **原先的結構**
> + go-tutorial/
>    + lesson1.go
>    + lesson2.go
>    + lesson3.go
> 

> **新的結構**
> + go-tutorial/
>    + bin/
>    + pkg/
>    + src/
>        + lesson1.go
>        + lesson2.go
>        + lesson3.go
>    

如此一來就是一個正確(而且傳統)的專案了，該結構在此不做詳細說明，因為該做法已經過時了

### STEP4 開始使用自創的套件吧

就好比 `fmt`, `sort` 一樣，今天來練習創建一個自訂的套件 `calc`：

> 
> + go-tutorial/
>    + bin/
>    + pkg/
>    + src/
>        + calc/
>            + calc.go
>        + lesson14.go
>    

檔名不一定要使用 calc.go，命名時以方便維護為原則

**./src/calc/calc.go**
```go=
package calc    // 注意：不是 main 是 calc

func Add(a int, b int) int{    // 請使用大寫開頭
    return a + b
}
```

> 注意
> 1. calc.go 這個檔案不是 main package 不能直接用 `go run` 執行
> 
> 2. 函式的開頭必需大寫，否則其他檔案無法引用，這有一點類似先前提到 scope 的概念，介紹方法(二) 後會統一說明

**./src/lesson14.go**

```go=
package main
import(
    "fmt"
    "calc"
)

func main(){
    fmt.Println(calc.Add(10, 20))
}
```

> 執行結果：
> 30

因為法一是舊方法所以我們點到為止，其他相關的規定我們在方法(二)會繼續說明

## 自訂套件 (二) 使用 go mod (新方法)

然而，使用方法(一)實在太麻煩了，每次寫個新專案還要去改環境變數，windows的還算好改，每次用linux改一改，改到不小心要重灌(我太菜)。因此，Go在 v1.10 後支援一個新的方法，`go mod`

### STEP1 找出一個空的資料夾

先把先前的工作目錄內的檔案移去別的地方，我們從頭開始！

### STEP2 使用 go mod init

使用 `go mod init` 直接創件一個新的模組。這個模組名稱要叫什麼呢？就叫 `practice` (練習) 好了！

![](https://i.imgur.com/TrV54NE.png)

此時會go會自動生成一個檔案：

> 
> + go-tutorial/
>     + go.mod
>   

我們來看看 go.mod 裡寫了什麼？

```go=
module practice

go 1.12
```

看起來也不難懂，`第 1 行` 就是指我們這個 module(模組) 叫作 practice， `第 3 行` 很明顯是告訴我們現在所使用的版本

這時我們把先前的文件丟回來

> 
> + go-tutorial/
>     + go.mod
>     + lesson1.go
>     + lesson2.go
>     + lesson3.go
>     + lesson14.go
>

### STEP3 開始使用自創的套件吧

這次我們一樣以 calc 來舉例

> 
> + go-tutorial/
>     + calc/
>         + calc.go
>     + go.mod
>     + lesson14.go
>

**./calc/calc.go**
```go=
package calc    // 注意：不是 main 是 calc

func Add(a int, b int) int{    // 請使用大寫開頭
    return a + b
}
```

**calc.go 這個檔案不能直接執行**

> 注意一點，函式的開頭必需大寫，否則其他檔案無法引用，這有一點類似先前提到 scope 的概念
> 

**./lesson14.go (與方法一稍有不同)**

```go=
package main
import(
    "fmt"
    // 注意很重要所以打三遍
    // 注意很重要所以打三遍
    // 注意很重要所以打三遍
    "practice/calc"
    // 這裡要用模組名稱 + / + 套件名稱
)

func main(){
    fmt.Println(calc.Add(10, 20))
}
```

---

接下來的課程， **統一以方法(二)為主** ，因為我不喜歡過時的東西

---

## 開始來自訂套件吧！(入門)

因為我很懶所以就拿昨天的Golang牌來爆改吧！

↓ 這是昨天收服的Golang牌 ↓

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

func printInfo(g geometry){
    fmt.Printf("%#v\n", g)
    fmt.Println("Its area:", g.area())
}

func main(){
    r1 := &rectangle{width:10, height:3}
    c1 := &circle{radius: 5}

    printInfo(r1)
    printInfo(c1)
}
```

剛剛示範用的套件名稱叫 calc 好像不太有風格，不然就改叫 geo 好了


> 
> + go-tutorial/
>     + geo/
>         + geo.go
>     + go.mod
>     + lesson14.go
>

**./lesson14.go** (我們希望實現的效果)
```go=
package main
import (
    "practice/geo"
)

func main(){
    r1 := &geo.Rectangle{Width:10, Height:3}
    c1 := &geo.Circle{Radius: 5}

    geo.PrintInfo(r1)
    geo.PrintInfo(c1)
}
```

**./geo/geo.go**

**不管三七二十一，通通都弄成大寫**

```go=
package geo

import "fmt"

type Geometry interface{    // geometry 改用大寫開頭
    Area() float64          // area 改用大寫開頭
}

type Rectangle struct{      // rectangle 改用大寫開頭
    Width  float64          // width 改用大寫開頭
    Height float64          // height 改用大寫開頭
}

func (r *Rectangle) Area() float64{     // area 改用大寫開頭
    return r.Width * r.Height           // 跟著前面改用大寫開頭
}

type Circle struct{     // circle 改用大寫開頭
    Radius float64      // radius 改用大寫開頭
}

func (c *Circle) Area() float64{            // 改用大寫開頭
    return c.Radius * c.Radius * 3.1416     // 跟著前面改用大寫開頭
}

func PrintInfo(g Geometry){
    fmt.Printf("%#v\n", g)
    fmt.Println("Its area:", g.Area())
}
```

> 執行 lesson14.go：
> &geo.Rectangle{Width:10, Height:3}
> Its area: 30
> &geo.Circle{Radius:5}
> Its area: 78.53999999999999
> 

果然跟上一次教的會得到一樣的結果

**有什麼其實是不用大寫的呢？**

仔細觀察
**./lesson14.go**
```go=
package main
import (
    "practice/geo"
)

func main(){
    r1 := &geo.Rectangle{Width:10, Height:3}
    c1 := &geo.Circle{Radius: 5}

    geo.PrintInfo(r1)
    geo.PrintInfo(c1)
}
```

> 哪些是真正會使用到 geo.go 中的？
> 1. Rectangle (struct) 及其成員(Width, Height)
> 2. Circle (struct) 及其成員(Radius)
> 3. PrintInfo()
> 

這是不是代表剩下的東西其實根本可以不用給 lesson14.go 知道，這時，我們可以把不需要 `lesson14.go` 存取的函式、方法、變數、...等全部改用小寫開頭

**./geo/geo.go**
**不需要給別人知道就用小寫開頭**

```go=
package geo

import "fmt"

type geometry interface{    // geometry 不用大寫開頭
    area() float64          // area 不用大寫開頭
}

type Rectangle struct{      // rectangle 改用大寫開頭
    Width  float64          // width 改用大寫開頭
    Height float64          // height 改用大寫開頭
}

func (r *Rectangle) area() float64{     // area 不用大寫開頭
    return r.Width * r.Height           // 跟著前面改用大寫開頭
}

type Circle struct{     // circle 改用大寫開頭
    Radius float64      // radius 改用大寫開頭
}

func (c *Circle) area() float64{            // 不用大寫開頭
    return c.Radius * c.Radius * 3.1416     // 跟著前面改用大寫開頭
}

func PrintInfo(g geometry){
    fmt.Printf("%#v\n", g)
    fmt.Println("Its area:", g.area())
}
```

> 執行 lesson14.go
> &geo.Rectangle{Width:10, Height:3}
> Its area: 30
> &geo.Circle{Radius:5}
> Its area: 78.53999999999999
> 

果然還是跟上面的示範得到一樣的結果

但是，老實說，這個範例就是缺少那個 「Golang的風格 」

---

## 中場休息 ─ 百合大法好

![](https://i.imgur.com/rSjACc1.jpg)

[庫洛魔法使第1季第24集](https://ani.gamer.com.tw/animeVideo.php?sn=10792)

![](https://i.imgur.com/0LnkgNg.jpg)
[科超S第1集](https://www.youtube.com/watch?v=FSPtFkncT4s&ab_channel=Muse%E6%9C%A8%E6%A3%89%E8%8A%B1-TW)

![](https://i.imgur.com/W4HbgqZ.jpg)

[Love Live! 虹咲學園 學園偶像同好會 第02話](https://www.youtube.com/watch?v=H5QvHbAEtxE)

![](https://i.imgur.com/rEH9CcZ.jpg)
[Love Live! 二期 第11話](https://www.youtube.com/watch?v=J8MyMJ1InG4)

---

## 寫出有 Go 風格的套件前先看 Go 的套件怎麼用

為了實作有 `Go 風格` 的套件，我們先來試試幾個套件的用法！

### 第一種：一堆實用的函式和常數的 math 套件

![](https://i.imgur.com/Oq0qJGI.png)

![](https://i.imgur.com/eY2wauq.png)


首先我們來簡單的用常數來計算半徑 10 的圓型面積吧！

```go=
package main
import (
    "fmt"
    "math"
)

func main(){
    r := 10.0
    fmt.Println(r * r * math.Pi)
}
```

> 執行結果：
> 314.1592653589793

這種套件很簡單，初學者都能上手

### 第二種：使用許多已由套件定義的型態的 http 套件

```go=
package main
import (
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
    // 這裡的 ResponseWriter 是一個 interface{}
    // 但是你不需要實作，因為 HandleFunc() 已經負責實作了
    // 我們只要呼叫 Write() 方法就可以了
    // 這個 Write() 方法內的參數是 []byte
    // 直接照先前慣例對 string 轉型成 []byte

    w.Write([]byte("Hello world"));
}

func main(){
    // func NewServeMux() *ServeMux
    mux := http.NewServeMux()       // 利用 http 包提供的函式新增一個 *ServeMux 型態

    // func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
    mux.HandleFunc("/",  handler)   // 使用這個 http 包定義的型態(*ServeMux) 的方法
                                    // 第一個參數是字串
                                    // 第二個參數是一個函式(你沒聽錯，函式也可以當參數)
                                    // 我們按照 http 規定的格式創建了符合規則的函式
                                    // func (ResponseWriter, *Request)
                                    // 沒錯，http.ResponseWriter 也是他們自訂的型態

    // 除了很經典地呼叫 New 開頭的函式來新增一個變數外
    // 也時常可以用套件本身提供的 struct 型態
    server := &http.Server {
        Addr: ":8080",
        Handler: mux,
    }

    // 呼叫 http 包裡對 *Server 定義的方法 ListenAndServe()
    server.ListenAndServe()
}
```


> 執行結果：
> 
> ![](https://i.imgur.com/jIbu139.png)
> 
> 按「允許存取」
> 進入瀏覽器
> http://localhost:8080/
> 瀏覽器呈現
> > Hello world
> > 
>

### 第三種：必需按照要求實作由套件定義的 interface 的 sort 套件

引用前一課所學的實現 sort.Interface 的範例

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

---

## 後記：

原本要介紹 Error handling 的，但我今天已經花了 3 小時在弄這篇教學文了，如果繼續往下教，可能會教太多，所以今天就先停在這裡吧！

本文圖片來自：
> + [庫洛魔法使第一季第廿二集](https://ani.gamer.com.tw/animeVideo.php?sn=10790)
> + [庫洛魔法使第一季第廿三集](hhttps://ani.gamer.com.tw/animeVideo.php?sn=10791)
> + [庫洛魔法使第一季第廿四集](https://ani.gamer.com.tw/animeVideo.php?sn=10792)