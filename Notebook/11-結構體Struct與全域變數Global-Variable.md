---
tags: Golang魔法使
---
# \#11 結構體 Struct & 全域變數 Global Variable | Golang魔法使

![](https://i.imgur.com/9lo5ldL.jpg)

戶外教學的第二天晚上老師們準備了試膽遊戲，讓小孩子們拿著蠟燭進入一個洞窟在返回到洞口。可惜這個試膽遊戲沒有接下來的課程精彩

![](https://i.imgur.com/WndtLTV.jpg)

↑ 小櫻到是在怕什麼啦，都見過大風大浪了\~，以後還有資料結構、演算法、作業系統、計算機組織等著你吔


![](https://i.imgur.com/SxQIpKy.jpg)

![](https://i.imgur.com/kUNxSN5.jpg)

![](https://i.imgur.com/xGli2IK.jpg)

```javascript=
$("#奈緒子").fadeOut(1000, function(){
    $("#利佳").fadeOut(1000, function(){
         $("#千春").fadeOut(1000);
    });
});
```

↑ jQuery 梗

為什麼大家會消(淡)失(出)呢？一定又是 Golang 牌搞得鬼

各位魔法使們，我們一起跟著小櫻和小狼，**收服「struct」並把被消失的同學們淡入回來吧！！**

## type 定義型態別名

收服 struct 之前，我們必需先搞懂什麼是 `type` ，如果是精通 C 的魔法使們，Go中的`type` 其實可以理解成 C 中的 `typedef` (但使用順序剛好相反)

```go=
package main
import "fmt"

type myint int

func main(){
    var num myint
    num = 10
    fmt.Println(num)
}
```

> 執行結果：
> 10

利用 `type` 可以實現「型態別名」什麼是型態別名呢？就是同個型態有不同的名稱。常見的型態別名比如 `rune`， `rune` 其實就是 `int32` 的別名，在第三課就有提到： [#3 字串 ─ 知世：反正我們是屬於正義的一方 | Golang魔法使](https://ithelp.ithome.com.tw/articles/10233799)

## struct 定義複合型態

struct 是一種複合型態，可以在一個 struct 內存放不同型態的變數，這一節也是相當複雜的單元

透過 struct 可以存放各種不同型態的變數，比如我想在一個變數中同時存有 `var name string` 和 `var age uint` 。

直覺的做法是分別宣告兩個變數，但實際上可以利用 `struct` 可以把這兩個型態包在一起：

```go=
struct{
    name string
    age uint
}
```

這個 struct 並**不是變數** 而是一個自創的型態，我們可以把這個自創型態命名成：`person` 

```go
type person struct{
    name string
    age uint
}
```

### 宣告 struct 型態的變數

宣告一個 `person` 型態的變數

```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func main(){
    var sakura person    // 型態：person
    sakura.name = "小櫻"
    sakura.age  = 10
    fmt.Println(sakura)
    fmt.Printf("%T", sakura)    // 利用 %T 印出型態
}
```

> 執行結果：
> {小櫻 10}
> main.person

型態為 person 可以理解但為什麼會印出 `main.person` 呢？這個會牽涉到後面的課程，留到以後再做說明

### 更簡短的宣告

```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func main(){
    sakura := person{
        name : "小櫻",
        age  : 10,
    }
    fmt.Println(sakura)
    fmt.Printf("%T", sakura)    // 利用 %T 印出型態
}
```

> 執行結果：
> {小櫻 10}
> main.person

### 不使用 type 直接使用 struct 宣告

```go=
package main
import "fmt"

func main(){
    sakura := struct{
        name string
        age uint
    }{
        name : "小櫻",
        age  : 10,
    }
    fmt.Println(sakura)
    fmt.Printf("%T", sakura)    // 利用 %T 印出型態
}
```

> 執行結果：
> {小櫻 10}
> struct { name string; age uint }
> 

可能一時之間沒辦法理解，請大家對照前一支程式把 `person` 的地方用 `struct{name string age uint}` 取代掉也許就能理解了

![](https://i.imgur.com/1XwjaCV.png)


## 在函式中使用 struct

以下為**錯誤範例** 試著將 person 中的 age 加一

```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func plusOne(p person){
    p.age = p.age + 1
}

func main(){
    sakura := person{
        name : "小櫻",
        age  : 10,
    }
    fmt.Println(sakura)
    plusOne(sakura)
    fmt.Println(sakura)
}
```

> 執行結果：
> {小櫻 10}
> {小櫻 10}

如同上一課所說，`sakura` 是一個 `person 型態` 的變數，並不是「指標變數」，若想要能在函式中更改 `sakura` 那麼我們需要取得 `sakura` 的指標才行

## 在函式中使用指向某個 struct{} 的指標


```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func plusOne(p *person){    // 改用 *person
    (*p).age = (*p).age + 1 // 第 10 行
}

func main(){
    sakura := person{
        name : "小櫻",
        age  : 10,
    }
    fmt.Println(sakura)
    plusOne(&sakura)        // 利用 & 取得 sakura 的位址
    fmt.Println(sakura)
}
```

> 執行結果：
> {小櫻 10}
> {小櫻 11}

`第 10 行`中這個寫法看似有些麻煩有沒有更簡單的寫法呢？其實是有的！在 Golang 中，如果是一個 struct{} 出來的變數是用 `.` 去取得值；而如果是一個指向 struct{} 的變數仍然可以直接透過 `.` 去取值，也就是說

```go=
(*p).age = (*p).age + 1
// 完全可以改成
p.age = p.age + 1
// 即使 p 的型態是 *person 不是 person
// 這個特性與 C, C++ 不太相同
```

```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func plusOne(p *person){
    p.age = p.age + 1  // 等效於 (*p).age = (*p).age + 1
}

func main(){
    sakura := person{
        name : "小櫻",
        age  : 10,
    }
    fmt.Println(sakura)
    plusOne(&sakura)
    fmt.Println(sakura)
}
```

> 執行結果：
> {小櫻 10}
> {小櫻 11}


## 利用 new() 宣告一個指向 struct{} 的指標變數

除了先前提到的宣告方式，在 Golang 中可以直接用 new() 宣告指向 struct{} 的指標變數，因為比起 struct{} 我們更喜歡使用 *struct{}

要注意的一點是，如果 new 裡所使用的是型態 T 那麼其產生的會是 *T

請觀察以下程式

```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func main(){
    // new() 出來的變數型態會是 new() 裡頭所放的指標
    // 即 sukura 的型態是 *person
    sakura := new(person)
    sakura.name = "小櫻"
    sakura.age  = 10

    fmt.Println(sakura)    // sakura 的型態是 *person
    fmt.Println(*sakura)   // *sakura 的型態是 person
}
```

> 執行結果：
> &{小櫻 10}
> {小櫻 10}

這個執行結果還滿有趣的，透過 `new(person)` 產生的 `sakura` 型態是 `*person`，所以在我們呼叫 `fmt.Println(sakura)` 時理論上是可以印出一串指向 sakura 位址的數字，但 Golang 反而是更人性化的印出記憶體的內容並在前面標示 `&` 來代表是取該值的位址


這個 new 的用法，跟 C 語言中 `stdlib.h` 套件的 `calloc()` 相似

## 全域變數

前幾篇文有提到 scope 的概念，然而如果將變數在某個地方宣告，那麼他的 scope 會相當之大，可以讓該Golang牌的所有函式都能直接使用，該變數稱為全域變數

```go=
package main
import "fmt"

var n int

func double(){
    n = n*2
}

func main(){
    n = 5
    fmt.Println(n)
    double()
    fmt.Println(n)
}
```

> 執行結果：
> 5
> 10

我們可以在 `func main()` 之外直接宣告「全域變數」這個變數不但可以在 `main()` 裡面使用，也能在其他函式使用，不需要傳遞

**或者你也可以這樣寫**

```go=
package main
import "fmt"

var n = 5

func double(){
    n = n*2
}

func main(){
    fmt.Println(n)
    double()
    fmt.Println(n)
}
```

> 執行結果：
> 5
> 10

先前並沒有提及 `var n = 5` 這種寫法，這種寫法其實就跟 `n := 5` 有異曲同功之妙。另一方面除了這種寫法，寫成 `var n int = 5` 也是可以的

但是**不可以**這樣寫

```go
package main
import "fmt"

n := 5

func double(){
    n = n*2
}

func main(){
    fmt.Println(n)
    double()
    fmt.Println(n)
}
```

> \# command-line-arguments
> 
> .\lesson11.go:4:1: syntax error: non-declaration statement outside function body

先前沒有特別提到，這種 `:=` 稱為短變數宣告，Golang 中規定這種寫法不可以在函式外使用

## 補充 ── 從 C 到 Go

在 C 魔法中 struct 搭配 `->` 取值，而指標則搭配 `.` 取值，會令人相當不習慣

我們嘗試用 C 魔法改寫以下 Golang 魔法

```go=
package main
import "fmt"

type person struct{
    name string
    age uint
}

func plusOne(p *person){
    p.age = p.age + 1  // 等效於 (*p).age = (*p).age + 1
}

func main(){
    sakura := person{
        name : "小櫻",
        age  : 10,
    }
    fmt.Println(sakura)
    plusOne(&sakura)
    fmt.Println(sakura)
}
```

> 執行結果：
> {小櫻 10}
> {小櫻 11}
> 

**以 C 改寫：**

```c=
#include<stdio.h>

// golang 的 type 和 c 的 typedef 順序相反
typedef struct p{
    char* name;
    unsigned int age;
}person;

// golang 的 *person 在 c 中是 person*
void plusOne(person* p){
    // 指標取值用 ->
    p->age = p->age + 1;
}

int main(){
    person sakura;
    // 一般的變數則用 .
    sakura.name = "小櫻";
    sakura.age  = 10;
    printf("{%s %d}\n", sakura.name, sakura.age);
    plusOne(&sakura);
    printf("{%s %d}\n", sakura.name, sakura.age);

    return 0;
}
```

> 執行結果：
> {小櫻 10}
> {小櫻 11}

本文多數圖片來自：
[庫洛魔法使第一季第十七集](https://ani.gamer.com.tw/animeVideo.php?sn=10785)