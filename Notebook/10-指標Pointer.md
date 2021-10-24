---
tags: Golang魔法使
---
# \#10 Pointer 指標 | Golang魔法使

昨天跟各位魔法使預告過了，今天將會是一個很困難的課程，如果你是小櫻的初級工具人，沒有寫過 C, C++, Rust 這類的語言，今天的庫洛牌可能在收服上會吃力許多


![](https://i.imgur.com/NterryL.jpg)

↑ 沒帶賴打跟小狼借個火


這個單元，我是覺得沒有很難啦(畢竟我是從 C 轉過來的人)，但是對初學者來說可能真的稍微難懂一些。所以要用圖海戰術來迎造這個單元很簡單的樣子

![](https://i.imgur.com/vwzRPH0.jpg)

## 什麼是指標？

我們先前提到的變數都是直接把右邊的值存進左邊變數名稱中，這些變數其實就在記憶體中。而記憶體中都有「位址」，先前 9 天的課中我們完全沒有提到。

而今天我們想要透過神奇魔法將這些變數所在的位址給提取出來

那麼要怎麼樣才能把這個位址抓出來呢？其實很簡單只要在變數前面加上 `&` 就能取得這個變數的位址了，而這個位址其實就是一串數字而已

```go=
package main
import "fmt"
func main(){
    // 宣告一個 int 變數
    num := 10
    // 取得 int 變數的位址
    p := &num    // 第 7 行
    // 將 num 的位址印出來
    fmt.Println(p)
}
```

> 執行結果：
> 0xc0000160a0
> 

這張Golang牌使用時結果不一定每次都一樣，由你的作業系統決定。

> 也許你會好奇，不是說是地址是一串數字為什麼會出來數字和英文？在魔法科學中以 `0x` 開頭的「數字」代表是 `16進位` 由 `0~f` 組成，跟一般我們常用的 `10進位` 不太一樣。我們的鑰匙(電腦)是由 0 和 1 的2進位系統組成，因為 **2進位與16進位之間比較好轉換** ，所以比起10進位 **魔法使們更常使用 16進位**

`第7行` 的 p 即為指標變數


至於「指標」這個東西到底什麼用呢？我們繼續看下去\~

![](https://i.imgur.com/iY1y6ca.png)

<!--
## 中場休息一下 ── 請大家愛惜鯖魚罐頭

這天是小櫻哥哥桃矢和雪兔哥的高中園遊會。雪兔哥隨便地就幹掉了籃球隊。還被籃球隊的邀請加入

+ 吃貨 (大食客一個人要吃三個人份的便當)
+ 籃球超強
+ 射箭超強 (後面集數)

這天雪兔和桃矢他們班在演戲，桃矢男扮女裝(女裝大佬)飾演灰姑娘，而雪兔則飾演被人家丟掉的鯖魚罐頭

![](https://i.imgur.com/W7tFdiM.jpg)

![](https://i.imgur.com/aEsp9bC.jpg)

![](https://i.imgur.com/M0JVszD.jpg)

我也很好奇為什麼日本動漫那麼常出現鯖魚罐頭？

而且芙蘭達也超愛，蛤？你說你不知道芙蘭達是誰？

![](https://i.imgur.com/7TwCOUJ.jpg)

![](https://i.imgur.com/DUOqiy0.jpg)
↑ 芙蘭達想用娃(ㄓㄚˋ)娃(ㄧㄠˋ)跟淚子換鯖魚罐頭

![](https://i.imgur.com/HMNL1rH.jpg)
↑ 芙蘭達從四次元胖次中拿出更多炸藥想跟淚子換鯖魚罐頭

![](https://i.imgur.com/yNqqjkh.jpg)
↑ 終於凹到一罐鯖魚罐頭想用炸藥打開來吃的芙蘭達

![](https://i.imgur.com/vTM1Y01.jpg)
↑ 火藥不小心用太兇

![](https://i.imgur.com/9cgj1jh.jpg)
↑ 知世：我想雪兔哥這個角色的目的是要讓大家愛惜東西

> **結論：芙蘭達我婆**

-->

## 利用指標更改其指向變數的值

### 用 var 宣告指標變數

如果指向的變數原先是 `int` 那麼這個指標變數的型態則是 `*int` (跟C語言的`int*`剛好相反) 

```go=
package main
import "fmt"
func main(){
    num := 10
    var p *int
    p = &num
    fmt.Println(p)
}
```

> 執行結果：
> 0xc0000160a0
> 

這張Golang牌使用時結果不一定每次都一樣，由你的作業系統決定。

### 利用指標變數取得指向的變數的值
直接印出指標變數時是印出地址出來，要怎麼樣才能透過地址把值給取出來呢？這時你只需要將指標變數前面加上 `*` 就可以把值取回來了哦！

```go
package main
import "fmt"
func main(){
    num := 10
    p := &num
    fmt.Println("num 的位址", p)
    fmt.Println("num 的值", *p)
}
```

> 執行結果：
> num 的位址 0xc0000160a0
> num 的值 10
>

num 的位址不一定每次都一樣，由你的作業系統決定。

### 利用指標變數更改指向變數的值

我們可以利用指標的特性直接更改其指向的變數的值

```go=
package main
import "fmt"
func main(){
    num := 10
    p := &num
    *p = 20 // 將 p 所址向的變數的值改為 20
    fmt.Println("*p =", *p)
    fmt.Println("num =", num)
}
```

> 執行結果：
> \*p = 20
> num = 20
> 

### 如何利用非指標變數更改另一個變數的值

不能。

```go=
package main
import "fmt"
func main(){
    num := 10
    p2 := num // 嘗試複製 num (第5行)
    p2 = 20
    fmt.Println("p2 =", p2)
    fmt.Println("num =", num)
}
```

> 執行結果：
> p2 = 20
> num = 10

`第 5 行` 宣告 `p2` 只是又找了一塊新的記憶體把 num 的值放進去而已，放進去後，`num`和`p2`就「田無溝水無流」了


## 所以指標可以做什麼？

### 以指標變數為參數，在函式中更動變數值

一般而言，我們傳遞至函式中的變數，都是複製一份值進去，無法直接更改傳進函式的變數值。

舉個例子

```go=
package main
import "fmt"
func double(num int){
    num = num * 2
}

func main(){
    num := 10
    double(num)
    fmt.Println("num =", num)
}
```

> 執行結果：
> num = 10
> 

在執行 double() 時只是把 10 這個數傳進去而已，並不會更改 main() 裡面的 num，雖然你還是可以換個方法去實現：

```go=
package main
import "fmt"
func double(num int) int{
    num = num * 2
    return num
}

func main(){
    num := 10
    num = double(num)
    fmt.Println("num =", num)
}
```

> 執行結果：
> num = 20
>

指標的作法：

```go=
package main
import "fmt"
func double(p *int){
    *p = (*p) * 2
}

func main(){
    num := 10
    double(&num)
    fmt.Println("num =", num)
}
```

> 執行結果：
> num = 20
>

透過指標，許多事情都能輕鬆很多。指標玩得最透徹的，非 Rust 莫屬，Golang 只是小菜一疊。

## 切片其實是一種指標

切片的本值上類似一個指標，所以傳遞切片時可以直接更改切片中的值

```go=
package main
import "fmt"

func double(nums []int){
    for i := 0; i < len(nums); i = i+1{
        nums[i] = nums[i] * 2
    }
}

func main(){
    nums := []int{1, 3, 4, 7, 9}
    double(nums)
    for _, v := range nums{
        fmt.Printf("%d ", v)
    }
}
```

> 執行結果：
> 2 6 8 14 18
> 

## Map 也是一種指標

```go=
package main
import "fmt"

func plusOne(m map[string]int){
    for k, v := range m{
        m[k] = v + 1
    }
}

func main(){
    tall := map[string]int{
        "小櫻" : 153,
        "知世" : 155,
        "小狼" : 156,
    }
    plusOne(tall)
    for _, v := range tall{
        fmt.Printf("%d ", v)
    }
}
```

> 執行結果
> 154 156 157

## 陣列不是指標

這個部分跟 C 不太一樣，基本上 C 中的陣列就是單純的指標，但在 Go 中陣列其實加了一些料，所以不能當指標來用。也因此，比起 C 語言的指標，Go 中的指標算是好學許多

```go=
package main
import "fmt"

func double(nums [5]int){    // 第 4 行
    for i := 0; i < len(nums); i = i+1{
        nums[i] = nums[i] * 2
    }
}

func main(){
    nums := [5]int{1, 3, 4, 7, 9}
    double(nums)
    for _, v := range nums{
        fmt.Printf("%d ", v)
    }
}
```

> 執行結果
> 1 3 4 7 9
> 

雖然在函式`double()`中更動了`nums`，但因為陣列`nums`傳進函式`double()`是傳送「值」而不是「址」所以並不會影響原先的陣列

有一點要注意的是，`第4行`中以陣列當參數時型態中要指定長度，要是沒有指定長度，Golang會把它當成切片而不是陣列

## 陣列透過取址也是可以當指標傳遞

然而，如果你很執著想要透過函式更改陣列的值也不是不行，透過取址的方式仍然可以實現

```go=
package main
import "fmt"

func double(nums *[5]int){
    for i := 0; i < len(nums); i = i+1{
        (*nums)[i] = (*nums)[i] * 2
    }
}

func main(){
    nums := [5]int{1, 3, 4, 7, 9}
    double(&nums)
    for _, v := range nums{
        fmt.Printf("%d ", v)
    }
}
```

> 執行結果：
> 2 6 8 14 18
> 

---

本文多數圖片來自：
1. [庫洛魔法使第一季第十四集](https://ani.gamer.com.tw/animeVideo.php?sn=10782)
2. [庫洛魔法使第一季第十五集](https://ani.gamer.com.tw/animeVideo.php?sn=10783)