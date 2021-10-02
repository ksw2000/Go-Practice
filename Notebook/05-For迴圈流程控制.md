---
tags: Golang魔法使
---
# \#5 For 迴圈流程控制 | Golang魔法使

這天小櫻和朋友們參加友枝國小所舉辦的馬拉松，卻在終點前不斷的繞圈圈，這次是「環」Golang牌搞的鬼，各位小櫻的工具人們，是否能收服他嗎？

![](https://i.imgur.com/oiXBR15.jpg)

![](https://i.imgur.com/EgwyYtx.jpg)

## 前言

![](https://i.imgur.com/wqqu9ux.jpg)
↑本作男主終於出現啦(不是寺田老師！！~~寺田老師是利佳的~~)

## 迴圈

在控制流程時除了上次所講的 `if`-`else if`-`else` 外，再來最重要的就是迴圈，什麼是迴圈呢？就是指定一個範圍的程式要重覆跑幾次，而其實這個「跑幾次」也不是一定要是個明確的數字，主要是要設定「迴圈到什麼時候為止」

一般的魔法中常見的迴圈有 `for` 迴圈和 `while` 迴圈，但在 Golang 魔法中，沒有 `while` 這個關鍵字，而是全都用 `for` 來表示。如果你是老工具人，你可以這樣理解：當有三個參數時 Go 的 for 就是一般常用的 `for` ，當只有一個參數時，Go 的 for 就是一般常用的 `while`。

### 只有一個參數的 for

> 實作一支程式，該支程式能印出 3 次「李君」(一開始不熟時小櫻都這樣叫李小狼，後面，大概第二季吧，都一直小狼君\~小狼君\~)

```go=
package main
import "fmt"
func main(){
    fmt.Println("李君")
    fmt.Println("李君")
    fmt.Println("李君")
}
```

這次我們用彈性一點的做法

```go=
package main
import "fmt"
func main(){
    i := 1           // i預設為 1
    for i <= 3 {     // i<=3 為真就執行大括號內的咒語，否則不執行
        fmt.Println("李君")
        i = i + 1    // 每次執行時 i 加 1
    }                // 回到 for 那行 (第5行)
}
```

> 執行結果：
> 李君
> 李君
> 李君
> 

1. 這一種作法是只要 `for` 右方的條件成立，程式就會往花括號`{` 內跑。
2. 如果一開始條件就不符就不會往花括號內做執行。
3. 可以把 `for` 看成「在離開 `}` 時會再一次回到最開頭的 `if`」(進階版的 `if`)

所以如果一開始條件就不合，則 `for` 裡的程式也不會執行。

```go=
package main
import "fmt"
func main(){
    i := 4
    for i <= 3 {     // 對於 i<=3 就執行大括號內的咒語
        fmt.Println("李君")
        i = i + 1 
    }
}
```

> 執行結果：
> (啥都沒有)
>

**再來一個範例**
> 試實作一支程式，可以印出 1~10
> 

```go=
package main
import "fmt"
func main(){
    i := 1          // i預設為 1
    for i <= 10 {   // i<=10 為真就執行 {} 內的咒語，否則不執行
        fmt.Printf("%d ", i)
        i = i + 1
    }               // 回到 for 那行 (第5行)
}
```

> 執行結果：
> 1 2 3 4 5 6 7 8 9 10
>

當然不只有這種寫法，也可以倒過來印出 10 ~ 1

```go=
package main
import "fmt"
func main(){
    i := 10        // i預設為 10
    for i >= 1 {   // i>=1 為真就執行 {} 內的咒語，否則不執行
        fmt.Printf("%d ", i)
        i = i - 1
    }              // 回到 for 那行 (第5行)
}
```

> 執行結果：
> 10 9 8 7 6 5 4 3 2 1
> 

### 簡化寫法：三個參數的 for

大家有沒有發現迴圈起手式，首先一定要有一個初始值，比如 `i := 1`，再來一定要有一個條件 `i <= 10`，最後，再回到最開始前一定會做一個遞增、遞減...的動作。我們可以把這三個動作同時放入 `for` 中。

**簡化寫法**
```go=
package main
import "fmt"
func main(){
    // 三個參數中以 ; 為分隔
    for i := 1; i <= 10; i = i+1 {
        fmt.Printf("%d ", i)
    }
}
```

**該寫法與以下的寫法幾乎相同**
```go=
package main
import "fmt"
func main(){
    i := 1          // i預設為 1
    for i <= 10{    // i<=10 為真就執行 {} 內的程式，否則不執行
        fmt.Printf("%d ", i)
        i = i + 1
    }               // 回到 for 那行 (第5行)
}
```

> 執行結果：
> 1 2 3 4 5 6 7 8 9 10
>

簡化寫法需要一段時間的熟悉，因為實際上使用時並不會這麼簡單，也許是用在走訪鏈結串列、走訪樹、走訪圖...，一定要確時熟悉簡化寫法的流程

<!--一般程式語言(python除外)會把剛剛教的只有一個參數的 `for` 叫作 `while`，但在 Golang 中，統一都叫 `for`，分辨的方法是：當作`while` 的 `for` 只有一個參數(沒有 `;`)，而「有兩個 `;` 做分隔的」則是一般常見 `for` 的用法-->

### 簡化寫法：三個參數的 for 不用每個都給滿

當然，在簡化時其實不用每個參數都給滿，只是「兩個 `;` 一定要有」，這樣魔仗才能理解該 Golang 牌的意思，否則噴錯給你看

**比如這張魔法牌**
```go=
package main
import "fmt"
func main(){
    i := 1
    for ;i <= 10; i = i+1 {
        fmt.Printf("%d ", i)
    }
}
```
**幾乎等於這張**
```go=
package main
import "fmt"
func main(){
    for i := 1; i <= 10; i = i+1 {
        fmt.Printf("%d ", i)
    }
}
```

**也幾乎等於這張**
```go=
package main
import "fmt"
func main(){
    i := 1
    for ;i <= 10; {
        fmt.Printf("%d ", i)
        i = i + 1
    }
}
```

使用上，就看各位小櫻的工具人如何應用，建議現階段初學工具人還是先把三個參數填好填滿

### 無窮迴圈
![](https://i.imgur.com/EEnxbVI.png)

如果不小心寫到永遠為真的參數，那這支程式就會永遠停不下來

```go=
package main
import "fmt"
func main(){
    i := 1
    for i == 1{
        fmt.Printf("%d ", i)
    }
}
```

> 執行結果：
> 1 1 1 1 1 1 1 1 ...(無限跑下去)
>

你會想說，電腦這麼笨不知道你寫了個無窮迴圈？電腦還真的這麼笨，有興趣的朋友可以搜尋「停機問題」

> [停機問題──維基百科](https://zh.wikipedia.org/wiki/%E5%81%9C%E6%9C%BA%E9%97%AE%E9%A2%98)


如果我現在真的需要一支無窮迴圈的程式，那麼你可以把 for 裡的條件設成 `true` 那他就會不斷執行

```go=
package main
import "fmt"
func main(){
    for true {
        fmt.Print("李君")
    }
}
```

**但其實Golang提供更簡單的寫法**

```go=
package main
import "fmt"
func main(){
    for{
        fmt.Print("李君")
    }
}
```

只要在 `for` 右方不給定任何參數，就會是一個無窮迴圈

### 強制跳離，break
![](https://i.imgur.com/xkbJ4UG.jpg)

在某些情況下，你會想要立刻跳離迴圈，(比如被困在迴圈中想要回到學校的小櫻)，這時，你可以使用 `break` 來直接跳離迴圈。什麼時候會需要使用 `break` 呢？通常是「如果寫在for的參數裡」會變太肥時才使用，不然一般我們靠第二個參數就可以處裡不需要 `break`

舉個例子
> 檢查 1 ~ 100 中是否有乘以 3 會變 201 的數？

```go=
package main
import "fmt"
func main(){
    find := false
    for i := 1; i <= 100; i = i+1 {
        if i * 3 == 201 {
            find = true
            break    // 直接跳出 for 的魔爪 (前往第11行)
        }
    }
    if find { // 第 11 行
        fmt.Println("存在！")
    }else{
        fmt.Println("不存在！")
    }
}
```

> 執行結果：
> 存在
>

因為`i` 一旦到了 67，就會通過條件，所以 68 ~ 100，這整趟都不用檢查，檢查了也是浪費時間，所以我們在找到答案時就直接讓他跳出

### 略過這一次，continue

除了有強制跳出的 break 外，還有一個神祕的關鍵字叫 continue，continue 會略過後面的程式碼，直接回到 for 迴圈那行

> 實作一支程式印出 1~50 內不是 5 的倍數(即5的倍數略過不印)

```go=
package main
import "fmt"
func main(){
    for i := 1; i <= 50; i = i+1 { // 第 4 行
        if i % 5 == 0 {
            continue    // 直接回到第 4 行
        }
        fmt.Printf("%d ", i)
    }
}
```

> 執行結果：
> 1 2 3 4 6 7 8 9 11 12 13 14 16 17 18 19 21 22 23 24 26 27 28 29 31 32 33 34 36 37 38 39 41 42 43 44 46 47 48 49
> 

<!--
#### 碎碎念

我是滿好奇為什麼要叫 continue(繼續)，不叫 skip(略過)，skip 感覺更有這個用法的韻味而且 skip 只有 4 個字母 continue 有 8 個，不過一般 continue 不是會縮寫成 cont 嗎，不然也縮寫成 cont 嘛！

像很多長單字都在逐漸縮短中，尤其 `rust` 很明顯 `function` 變 `fn`，`implement(s)` 變 `impl`，`string` 用 `str`。還有像在有些語言中 `synchronize(d)` 變 `sync`, `asynchronize(d)` `async` 這種，到底誰沒事拼的出 `synchronized` 啦，就是在說你Java魔法使(而且還要記得用過去分詞哦)
-->
## 變數作用範圍

不知道大家有沒有遇到類似的問題，當我在 for 迴圈外(第 7 行)要取用 for 迴圈的變數(第 4 行)時就無法使用，這是很正常的事，在許多的程式中，幾乎都會使用這個規則，在 `{ }` 中，我們稱為 `scope` (中文：範圍，其實還包含右大括號`{`左邊那些) ，在這個 `scope` 中(4 ~ 6行)宣告的變數不能在這外面使用

![](https://i.imgur.com/uI2NuGC.png)

也就是說 `i` 所能使用的範圍只能這樣：

![](https://i.imgur.com/fu7lftz.png)

若要突破這個狀況，可以把 i 宣告在 scope 之外

```go=
package main
import "fmt"
func main(){
    i := 1    // 在 for 迴圈外先做宣告
    for ; i <= 10; i = i+1 {
        fmt.Printf("%d ", i)
    }
    fmt.Printf("\n\n%d ", i)
}
```
> 執行結果：
> 1 2 3 4 5 6 7 8 9 10
>
> 11

這是這個 `i` 所能作用的範圍

![](https://i.imgur.com/LUtsLwj.png)

所以其實

```go=
i := 1
for ; i <= 10; i = i+1
```

**和**
```go=
for i := 1 ; i <= 10; i = i+1
```

還是有微小的差異的

```go=
package main
import "fmt"
func main(){
    for i := 1; i <= 10; i = i+1 {
        fmt.Printf("%d ", i)
    }
    // i 在這裡已經不可再被存取了
    // TODO
}
```

```go=
package main
import "fmt"
func main(){
    i := 1
    for i <= 10 {
        fmt.Printf("%d ", i)
        i = i+1
    }
    // i 仍然可以被存取
    // TODO    
}
```

## 後記

其實「環」這張庫洛牌在第1季第21集才有出現，但劇情準備帶第8集，可是第8集劇情一次又太多，看來只能移花接木一下了。而且第8集的圖也太經典了吧

![](https://i.imgur.com/TDvPN5v.jpg)
↑ 利佳你這樣不行吔，20年前的動畫就能這麼開放真的很猛

![](https://i.imgur.com/vVjSGVi.jpg)

本文多數圖片來自：
+ [庫洛魔法使第一季第八集](https://ani.gamer.com.tw/animeVideo.php?sn=10776)
+ [庫洛魔法使第一季第廿一集](hhttps://ani.gamer.com.tw/animeVideo.php?sn=10789)