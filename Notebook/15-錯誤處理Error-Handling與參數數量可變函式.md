---
tags: Golang魔法使
---
# \#15 錯誤處理 Error Handling & 參數數量可變函式 | Golang魔法使

這天狡滑的Golang牌變成了小櫻的樣子想害死小櫻的哥哥，螢幕前的鍵盤Golang魔法使們，你們能幫忙小櫻收服這張「誤入歧途的`Error`Golang牌嗎？」

![](https://i.imgur.com/NiEeqRI.jpg)

---

## fmt.Scanf()
為了更快理解 Error handling 及參數數量可變的函式，我們以 `fmt.Scanf()` 來舉例

![](https://i.imgur.com/JVO1pJv.png)

```go=
func Scanf(format string, a ...interface{}) (n int, err error)
```

其中會讓人困感的有兩個地方
1. ...interface{} 是什麼？
2. error 是什麼

## error

```go=
type error interface {
    Error() string
}
```

`error` 也是一個原生型態，但他是一個由 interface 所定義的型態，通常設為 nil 時當做沒有錯誤，反之如果為非 nil 則代表有錯誤發生，利用這個方法就能簡單地偵測是否有錯誤發生

**觀察以下程式**
```go=
package main
import "fmt"
func main(){
    var num int
    fmt.Print("請輸入一個整數：")
    fmt.Scanf("%d", &num)    // 注意：參數是使用指標變數
    fmt.Println("你輸入的是", num)
}
```

這是一張簡單的Golang牌，允許使用者從命令提示字元中輸入

> 執行：
> 
> 請輸入一個整數：
>
> (以鍵盤輸入 10 然後按 Enter)
>
> 你輸入的是 10
>

但是也許使用者不按照規定而是輸入字元比如使用者輸入 A

那麼這張Golang牌會印出什麼呢？

> 請輸入一個整數：
>
> (以鍵盤輸入 A 然後按 Enter)
>
> 你輸入的是 0
>

因為有錯誤發生，所以輸入非數字時，並不會將結果寫進 `num` 之中，而是直接印出預設值。

如果今天要設計成在輸入時加上檢查功能，不允許用戶隨意鍵入那該怎麼做呢？

我們可以利用 `fmt.Scanf()` 的回傳值來達成，該函式預設有兩個回傳值，第一個是「有效的參數個數」(比如要求一個int，用戶確實輸入int那就算有效)，第二個是「錯誤」(錯誤不為nil時代表發生錯誤)

```go=
package main
import "fmt"
func main(){
    var num int
    fmt.Print("請輸入一個整數：")
    n, err := fmt.Scanf("%d", &num)
    if err != nil{    // err != nil 表示有錯誤發生
        fmt.Println("輸入有效的參數數量為", n)
        fmt.Println("錯誤訊息", err)
        fmt.Println("輸入格式有誤")
    }else{
        fmt.Println("輸入有效的參數數量為", n)
        fmt.Println("你輸入的是", num)
    }
}
```

> 執行結果(輸入 10)：
> 請輸入一個整數：10
> 輸入有效的參數數量為 1
> 你輸入的是 10


> 執行結果(輸入 A)：
> 請輸入一個整數：A
> 輸入有效的參數數量為 0
> 錯誤訊息 expected integer
> 輸入格式有誤

利用 `fmt.Scanf()` 所回傳的 `error` 來判斷是否有誤，讓錯誤處理變的非常容易，不用一直 catch you catch me (一般Golang牌傳統上的錯誤處理是用 try-catch-finally 來達成，要拋出自訂的錯誤更是麻煩)

## 利用 errors.New() 自訂錯誤訊息

要自訂 `error` 非常簡單

1. 引入 "erros" 套件
2. 判斷錯誤條件，確認為錯誤情況時回傳 error.New("錯誤訊息") 否則回傳 nil

> 實作一張可以執行整數除法的 Golang 牌。除數不可為 0，為 0 時噴出自訂的錯誤訊息
>

```go=
package main
import (
    "fmt"
    "errors"
)

func div(a, b int) (int, error){
    if b == 0 {
        err := errors.New("除數不可為零")
        // err 的型態：error
        return 0, err
    }
    return (a/b), nil
}

func main(){
    var a, b int
    fmt.Print("輸入兩個整數：")
    _, err := fmt.Scanf("%d %d", &a, &b)

    if err != nil{
        fmt.Println("格式錯誤")
        return  // 直接離開 main()
    }

    ans, err := div(a, b)
    if err != nil{
        fmt.Printf("%s", err)
    }else{
        fmt.Printf("%d / %d = %d",a ,b, ans)
    }
}
```

1. 正確輸入

    > 執行結果：(輸入 30 2)
    > 輸入兩個整數：30 2
    > 30 / 2 = 15

2. 除數為零

    > 執行結果：(輸入 20 0)
    > 輸入兩個整數：20 0
    > 除數不可為零

3. 格式亂打

    > 執行結果：(輸入 A 3)
    > 輸入兩個整數：A 3
    > 格式錯誤

## 利用 fmt.Errorf() 更快地使用自訂錯誤訊息

但是利用 `errors.New()` 還要引用 `errors` 套件，實在太麻煩了因此還可以使用 `fmt.Errorf()` 來實做，用法跟 `fmt.Printf()`, `fmt.Sprintf()` 很像，只是 `Printf()` 是直接印出，`Sprintf()` 是回傳 `string`，而 `Errorf()` 則是回傳 `error`

```go=
package main
import "fmt"

func div(a, b int) (int, error){
    if b == 0 {
        // 使用上比 errors.New() 更好上手
        return 0, fmt.Errorf("%d / %d 除數不可為零", a, b)
    }
    return (a/b), nil
}

func main(){
    var a, b int
    fmt.Print("輸入兩個整數：")
    _, err := fmt.Scanf("%d %d", &a, &b)

    if err != nil{
        fmt.Println("格式錯誤")
        return  // 直接離開 main()
    }

    ans, err := div(a, b)
    if err != nil{
        fmt.Printf("%s", err)
    }else{
        fmt.Printf("%d / %d = %d",a ,b, ans)
    }
}
```


> 執行結果 (輸入 10 0)：
> 輸入兩個整數：10 0
> 10 / 0 除數不可為零


---
## 中場休息：庫洛牌占卜


![](https://i.imgur.com/eX3lugE.jpg)

大家學會了嗎？請幫我顯示出「計算機組織期中期末考」題目與解答謝謝

---

## 自己實作 error 的 interface{}

error 是一個 interface{}，只是官方提供了幾個已經幫我們設計好的函式可以直接使用 `errors.New()` & `fmt.Errorf`

那麼我想自己實作要怎麼實現呢？

> ### type error
> The error built-in interface type is the conventional interface for representing an error condition, with the nil value representing no error.
>
> ```go=
> type error interface {
>     Error() string
> }
> ```
>

> 只要我們自己創立一個新的型態該型態滿足 Error() 方法，那麼也可以視為 error 來使用

```go=
type myError struct{
    msg string
}

func (m *myError) Error() string{
    return m.msg
}
```

如此一來 `*myError` 也符合 `error` 了
我們來實際使用看看！

```go=
package main
import "fmt"

type myError struct{
    msg string
}

func (m *myError) Error() string{
    return m.msg
}

func div(a, b int) (int, error){
    if b == 0 {
        err := &myError{msg: "除數不可為零"}
        return 0, err
    }
    return (a/b), nil
}

func main(){
    var a, b int
    fmt.Print("輸入兩個整數：")
    _, err := fmt.Scanf("%d %d", &a, &b)

    if err != nil{
        fmt.Println("格式錯誤")
        return  // 直接離開 main()
    }

    ans, err := div(a, b)
    if err != nil{
        fmt.Printf("%s", err)
    }else{
        fmt.Printf("%d / %d = %d",a ,b, ans)
    }
}
```

> 執行結果 (輸入 10 0)：
> 輸入兩個整數：10 0
> 除數不可為零


## 參數數量可變的函式

剛剛看到 `fmt.Scanf()` 的介紹

```go=
func Scanf(format string, a ...interface{}) (n int, err error)
```

`a ...interface{}` 到底是什麼意思呢？這個在`fmt.Printf()` 也可以看到。這個用法代表在呼叫該函式時可以給予不同數量的參數

```go=
func Printf(format string, a ...interface{}) (n int, err error)
```

直接舉個例子可以更快懂一些

```go=
package main
import "fmt"

func sum(a ...int) int{
    s := 0
    for _, v := range a {
        s = s + v
    }
    return s
}

func main(){
    fmt.Println(sum(1, 2, 3, 4))
    fmt.Println(sum(2, 4, 6))
    fmt.Println(sum(8))
}
```

對於數量不確定的參數，可以利用 `...type` 來傳遞，比如 `...int`, `...float64`

而在函式中要讀取這些「數量不確定的參數」時可以將他們視為 slice ，直接利用 `for + range` 去存取即可

至於 `fmt.Scanf()`, `fmt.Printf()` 中所使用的 `...interface{}` 是一種很特殊的用法，因為所有的型態都能滿足空的 `interface`，所以所有的型態都能視為 `interface{}` 利用這個方法，可以使參數接受所有的型態

這個空介面的用法前天有補充到： [#13 介面 Interface & 排序實戰 sort.Interface | Golang魔法使](https://ithelp.ithome.com.tw/articles/10236366)

另外要注意的是，參數數量不確定的用法，只能用在最後一個參數，不然 Golang 沒有辦法搞懂到底哪些參數是已經確定數量的哪些是不確定數量的

如果你嘗試不把「可變數量的參數」放在最後一個時：

```go=
func sum(a ...int, b int) int{

}
```

> 編譯錯誤：
> syntax error: cannot use ... with non-final parameter a

## Panic ── 錯誤，而且強制停止

當錯誤發生至需要強制中斷Golang牌時，可以利用 panic() 來實現。使用 panic() 會導致整個程式停止，所以能使用 error 時盡可能不使用 panic()

我們可以稍微修改先前除數為零回傳error的Golang牌，將其改成panic()版本：

```go=
package main
import "fmt"

func div(a, b int) int{
    if b == 0 {
        panic("除數不可為零")        // 立即結束程式！
        fmt.Println("這行不會執行")
    }
    return a / b
}

func main(){
    var a, b int
    fmt.Print("輸入兩個整數：")
    _, err := fmt.Scanf("%d %d", &a, &b)

    if err != nil{
        fmt.Println("格式錯誤")
        return  // 直接離開 main()
    }

    fmt.Printf("%d / %d = %d",a ,b, div(a, b))
}
```

> 執行結果 (輸入 10 0)：
> 輸入兩個整數：10 0
> panic: 除數不可為零
> 
> goroutine 1 [running]:
> main.div(...)
>        C:/Users/liao2/OneDrive/go- tutorial/lesson15.go:6
> main.main()
>        C:/Users/liao2/OneDrive/go-tutorial/lesson15.go:22 +0x330
> exit status 2

---

## 後記

Goalng 中的 interface 真的偏難懂，但超好用，希望大家能在這幾次課程中能理解他的用法

![](https://i.imgur.com/H7QrdcK.jpg)

本文圖片大多來自：
[庫洛魔法使第一季第廿五集](https://ani.gamer.com.tw/animeVideo.php?sn=10793)