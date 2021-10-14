---
tags: Golang魔法使
---
# \#6 函式 Function | Golang魔法使

![](https://i.imgur.com/7BwrtoP.jpg)

**圖文不符，謝謝**

當我們 code 越變越多時，我們可以將常用的 code 包成「函式」(function)，這樣一來，可以讓程式變得更簡潔。另外，也可以使用 Golang 本身提供的函式，或者其他魔法使創造出來的函式

## 函式 func

函式什麼呢？簡單來說，就拿數學中的函數來舉例，比如有一個函式 `f(x) = 2x + 3` 當我給定 `x` 時 `f(x)` 就會有我要的答案

數學上的函式主要都是數字和數字在玩，而魔法語言則更多元，可以是數字、字串、布林...等

直接舉個例子來看

> 實作一個函式，可以將浮點數取絕對值
> 

```go=
package main
import "fmt"

// func [函式名](輸入值名稱 輸入值型態)輸出值型態

func abs(n float64) float64{
    if n >= 0 {
        return n    // 回傳 n
    }
    return -n       // 回傳 -n
}

func main(){
    // abs() 將回傳值傳給 fmt.Println
    fmt.Println(abs(10.0))
    fmt.Println(abs(-10.0))
}
```

> 執行結果
> 10
> 10
> 

在一張基本的Golang牌中，一定會有 `func main()` 這個函式最先執行，Golang的`main()`函式沒有接收任何值，所以 `()` 內是空的，又因為沒有回傳任何值所以 `()` 後也是空的。第一次接觸函式可能會覺得有點抽象，因此以下多舉幾個函式使用的例子

### 沒有回傳值的函式
函式不一定要有回傳值，函式可以只是一個動作，比如我們實作一個函式

> 實作一個函式輸入字串 someone，印出「someone啊，你快變成懲戒的鎖鏈」

```go=
package main
import "fmt"

func beChain(someone string){
    fmt.Println(someone + "啊，你快變成懲戒的鎖鏈")
}

func main(){
    beChain("風")
    beChain("水")
}
```

> 執行結果：
> 風啊，你快變成懲戒的鎖鏈！
> 水啊，你快變成懲戒的鎖鏈！
> 

![](https://i.imgur.com/rP0T5KB.jpg)
[圖源：庫洛魔法使第一季第四集](https://ani.gamer.com.tw/animeVideo.php?sn=10772)

> 初學者可能會對「有回傳值」和「沒回傳值」搞的很混亂。在沒有回傳值的函式中，可以想象成只是執行了某些行咒語，而有回傳值的函式則是除了執行這些咒語外還會把值帶回來給其他地方使用，比如給 `fmt.Println()` 使用，當然也可以把值帶給變數使用

### 沒有輸入值的函式

沒有規定函式一定要有輸入值

```go=
package main
import "fmt"

func release() string{
    return "隱藏著黑暗力量的鑰匙啊，在我面前顯示你真正的力量，跟你定下約定的小櫻命令你，封印解除！"
}

func main(){
    // release() 將回傳值傳給 fmt.Println
    fmt.Println(release())
}
```
> 執行結果：
> 隱藏著黑暗力量的鑰匙啊，在我面前顯示你真正的力量，跟你定下約定的小櫻命令你，封印解除！
> 

### 2個以上輸入值的函式
當然函式也沒有規定你只能有一個輸入值

```go=
package main
import "fmt"

func release(power string, name string) string{
    return "隱藏著" + power + "力量的鑰匙啊，在我面前顯示你真正的力量，跟你定下約定的" + name + "命令你，封印解除！"
}

func main(){
    // release() 將回傳值傳給 fmt.Println
    fmt.Println(release("星星", "小櫻"))
}
```

> 執行結果
> 隱藏著星星力量的鑰匙啊，在我面前顯示你真正的力量，跟你定下約定的小櫻命令你，封印解除！

> 補充：星星鑰匙是第二季的內容

### 2個以上的回傳值

Golang 支援多回傳值。這個特性也可以用來實作錯誤處理。

```go=
package main
import "fmt"

func calc(a int, b int) (int , int){
    return a+b, a-b
}

func main(){
    sum, diff := calc(5, 2)
    fmt.Println(sum)
    fmt.Println(diff)
}
```

> 執行結果：
> 3
> 7
> 

在等號左側放入接收回傳的變數，並用 `,` 做分隔即可實現多回傳值。舉個例子比較好懂：

```go=
package main
import "fmt"

func main(){
    a, b := 3, 10    // Golang 允許一次替兩個變數賦值
    fmt.Println(a, b)
}
```
> 執行結果：
> 3 10
> 

**有兩個回傳值但一個值暫時用不到該怎麼處理？**
前面的章節有提到說在 Go 的世界中，如果宣告沒使用的變數編譯器會報錯，那如果我只想使用其中一個回傳值該怎麼辦？此時只要將不使用的值的變數設成 `_` 即可

```go=
package main
import "fmt"

func calc(a int, b int) (int , int){
    return a+b, a-b
}

func main(){
    sum, _ := calc(5, 2)
    fmt.Println(sum)
}
```

> 執行結果：
> 7
> 

## 為回傳值命名

在Golang中，你甚至可以替回傳值命名。

```go=
func calc(a int, b int) (sum int , diff int){
    // 注意 sum 和 diff
    // 視為已經宣告，所以是用 = 不是 :=
    sum, diff = a+b, a-b
    return
}
```

要這樣寫也是沒問題的
```go=
func calc(a int, b int) (sum int , diff int){
    sum, diff = a+b, a-b
    return sum, diff
}
```

## 注意事項

1. 在使用函式時，只要一觸發 return ，那麼後方的程式碼就不會繼續執行
2. 有回傳值的函式不一定要去接收回傳值
3. 函式中所宣告的變數有所謂的 scope，不會影響其他函數中所設的變數
4. 有規定回傳值型態的函式，在函數最尾一定要使用 return
5. 宣告函式時擺放的順序不重要

**範例 1 & 2：reutrn 後的程式碼並不會執行, 有回傳值的函式不一定要去接收回傳值**
```go=
package main
import "fmt"

func test(num int) int{
    if num > 10{
        fmt.Println("num > 10")
        return num                // 第 7 行
    }
    fmt.Println("num < 10")
    return num
}

func main(){
    test(100)    // 第 14 行
}
```

> 執行結果：
> num > 10
> 

很明顯可以看出，當程式執行到`第 7 行`後，就會跳回`第 14 行`，後方的程式碼就不會再執行了。只要在函式中遇到 return ，程式就不會繼續將後方的咒語解析。

另一方面，有回傳值(不管有幾個回傳值)的函數並沒有強制一定要有接收他的變數。

**範例3：函式中所宣告的變數有所謂的 scope**
在上一次教的迴圈中有提到 scope，在迴圈內的變數並不可以在迴圈外使用，這件事同樣套用在 function 上

```go=
package main
import "fmt"

func test(num int){
    fmt.Println("in test()", num)
}

func main(){
    num := 10
    test(100)
    fmt.Println("in main()", num)
}
```

> 執行結果：
> in test() 100
> in main() 10

雖然 `main()` 和 `test()` 中都是使用 `num` 但是兩者完全不會受到影響，這就是所謂的 `scope`，如此一來可以方便魔法使們更方便使用變數，使用時只要顧慮自己的 `scope` 即可

**範例4：有規定回傳值型態的函式，在函數最尾一定要使用 return**
許多魔法卡不會檢查這個錯誤，但是Golang會。首先我們先借用改造一下範例1的函式

![](https://i.imgur.com/qWjFtqy.png)

![](https://i.imgur.com/feCGG3k.png)

大家懂了嗎？

**範例5：函式擺放的順序不重要**

在一些魔法咒語中(如C)，一定要先宣告才能呼叫，沒有宣告的函式是不可以呼叫的，但在 Golang 中則沒有這個限制

顛倒範例二的 `func test()` 和 `func main()`

```go=
package main
import "fmt"

func main(){
    num := 10
    test(100)                      // 先呼叫使用 func test
    fmt.Println("in main()", num)
}

func test(num int){                // 後宣告 func test
    fmt.Println("in test()", num)
}
```
執行結果與範例二相同無異

## 後記

在某些魔法中如 C, Java, Dart 使用 function 時不使用 function 這個關鍵字，而是直接給定回傳型態和傳入型態

而某些魔法則使用 `function` 來作為函式的關鍵字，但因為 `function` 有足足 8 個字母，所以各家魔法門派開始縮簡這個關鍵字，不過有時大家縮簡邏輯不太一樣，真的令人感到混亂

+ PHP, Javascript `function` (1995年)面市
+ Go `func` (2009年面市)
+ Kotiln `fun` (2011年面市)
+ Rust `fn` (2010年面市)
+ Python `def` (1991年面市)

如果以字數來看的話 Rust 大勝只用了僅僅 2 個字母，我先來預言一下，以後的魔法咒語說不定會直接以 `f` 來代替 `function`

本文多數圖片來自：
+ [庫洛魔法使第一季第八集](https://ani.gamer.com.tw/animeVideo.php?sn=10776)
+ [庫洛魔法使第一季第九集](hhttps://ani.gamer.com.tw/animeVideo.php?sn=10777)