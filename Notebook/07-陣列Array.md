---
tags: Golang魔法使
---
# \#7 陣列 Array | Golang魔法使

![](https://i.imgur.com/7XrFaQN.jpg)
[圖源：庫洛魔法使第一季第八集](https://ani.gamer.com.tw/animeVideo.php?sn=10776)

小櫻的爸爸在友枝國小的運動會上做了果凍給小櫻的朋友吃。一個變數名稱只能代替一個數值，有什麼方法可以只用一個變數就代替六個值呢？

![](https://i.imgur.com/LMMfzs7.jpg)


## 陣列 (Array)
在大多數的魔法語言中，有一個稱為「陣列」的東西，他可以僅用一個變數名稱代替多個變數值，在Golang中「陣列」有幾個限制

1. 同一個陣列中只能存放同一個型態的值，比如整數陣列、浮點數陣列、字串陣列
2. 陣列的長度在宣告後就固定不可以改變了
3. 宣告陣列的長度時只能是使用確定不變的數

### 宣告

#### 方法一
```go=
package main
import "fmt"
func main(){
    var jelly [6]string    // 宣告長度為 6 的 string array
    jelly[0] = "小櫻"
    jelly[1] = "知世"
    jelly[2] = "桃矢"
    jelly[3] = "雪兔"
    jelly[4] = "小櫻他爸"
    jelly[5] = "可能雪兔吃兩個"
    fmt.Println(jelly[0])
}
```

> 執行結果：
> 小櫻
>

> 特別注意的是：在存取陣列時一律由零開始做存取
> 

另一方面，並不是宣告出來的 6 個值都一定要設定：

```go=
package main
import "fmt"
func main(){
    var jelly [6]string    // 宣告長度為 6 的 string array
    jelly[0] = "小櫻"
    jelly[1] = "知世"
    jelly[2] = "桃矢"
    jelly[3] = "雪兔"
    jelly[4] = "小櫻他爸"
    fmt.Println(jelly[5])
}
```

> 執行結果：
> (什麼都沒有)
>

雖然我們宣告了長度為 6 的陣列但其實也不用全部都給值，在沒有魔法使給值的情況下，Golang會自動分配`預設值`，比如 `string` 會預設給`空字串`，`int`, `uint`, `float32`... 數值類的會給 `0`, `bool` 會給 `false`。在有些比較老舊的魔法中如 C, C++, 初始值會是一個不確定的數(記憶體中的殘值)，因為對每個陣列中的值給定預設值其實會花一些時間，一定程度上是會影響效能的

並不是只有陣列有「預設值」的規則，宣告一般變數而且沒有給值的話，都是會給預設值

```go=
package main
import "fmt"
func main(){
    var n int
    fmt.Println(n)
}
```

> 執行結果：
> 0
>

#### 方法二
由於方法一太麻煩了，要打很多字，我們試試看方法二。這一種寫法簡單很多

```go=
package main
import "fmt"
func main(){
    var jelly = [6]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    fmt.Println(jelly[1])
}
```
> 執行結果：
> 知世
> 

**有換行的寫法**
另外要注意，如果有換行的話最後一行仍需以逗號做結尾(跟json不一樣json是最後一個不加，這個是最後一個一定要加)
```go=
package main
import "fmt"
func main(){
    jelly := [6]string{
        "小櫻",
        "知世",
        "桃矢",
        "雪兔",
        "小櫻他爸",    // 一定要加逗號
    }
    fmt.Println(jelly[3])
}
```

> 執行結果：
> 雪兔
> 

**為什麼要這樣設計(我的猜測)**
有學過其他魔法的朋友應該知道有些魔法結尾要加 `;`，其實在 `Golang` 中也是，但是編譯器會自己加上。官方建議不要加分號(可以加，不會報錯)，因為編譯器會自己透過一些方式去判斷，所以如果最後一行不打逗號，編譯器可能會先自己把分號加在那，導致後續會噴錯，(這是我的猜測啦)

為什麼我會這樣猜呢？因為如果你在最後一個值後直接加花括號，**是沒有問題的**，所以並不是說最後一個就一定要打逗號。

```go=
package main
import "fmt"
func main(){
    jelly := [6]string{
        "小櫻",
        "知世",
        "桃矢",
        "雪兔",
        "小櫻他爸"}
    fmt.Println(jelly[3])
}
```
> 執行結果：
> 雪兔
> 


#### 方法三
當然我們也可以使用 `:=` 來實現賦值

```go=
package main
import "fmt"
func main(){
    jelly := [6]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    fmt.Println(jelly[2])
}
```

#### 方法四
如果我懶的自己算長度，也可以請編譯器幫忙計算，只要在 `[]` 放入 `...` 編譯器就會幫你算長度

```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    fmt.Println(jelly[4])
}
```

> 執行結果：
> 小櫻他爸
> 

### 取得相關數值
#### 要求讀寫超過陣列長度的值，魔仗會噴錯

```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    fmt.Println(jelly[5])
}
```

> 執行結果：
> \# command-line-arguments
.\lesson07.go:5:22: invalid array index 5 (out of bounds for 5-element array)

#### 利用 len() 取得陣列長度
len 是 length 的縮寫

```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    fmt.Printf("陣列長度：%d", len(jelly))
}
```
> 執行結果：
> 陣列長度：5

因為Golang不是物件導向語言所以是用len()這種函式的型式來取得陣列長度(和python一樣)，這個取陣列長度在每個語言中都有各自的作法，比如： `C` 完全不能取得長度(除非你自己實作), `java` 是用 `.length()`, `javascript` 和 `dart` 是用 `.length`, `PHP` 是用 `count()`(php的命名邏輯還真獨特呢)，總之各家的寫法都不一樣，還滿有趣的

#### 利用 for 遍歷一個陣列

既然我們都知道陣列的長度了，那我們能否透過 for 迴圈把陣列的內容 print 出來呢？

```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for i:=0; i<len(jelly); i=i+1{
        fmt.Println(jelly[i])
    }
}
```

> 執行結果：
> 小櫻
> 知世
> 桃矢
> 雪兔
> 小櫻他爸
> 

#### 利用 for 和 range 遍歷一個陣列
Golang 提供另一個方法來走訪陣列，這個方法可以用來走訪所有有「迭代器」的資料結構。至於「迭代器」是什麼？我們會在往後的課程中提到！

```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for k, v := range jelly{
        fmt.Printf("jelly[%d] = %s\n", k, v)
    }
}
```

> 執行結果：
> jelly[0] = 小櫻
> jelly[1] = 知世
> jelly[2] = 桃矢
> jelly[3] = 雪兔
> jelly[4] = 小櫻他爸
> 

+ `k` 代表的是 key (鍵)：為 0, 1, 2, 3, 4
+ `v` 代表的是 value (值)：為 "小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"

如果有其中一個值是不需要的可以用 `_` 去忽略，如果你沒用 `_` 又不使用宣告出來的變數，魔仗會噴錯 (上一個單元中多回傳值的部分有提到，可以回去看)

```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for _, v := range jelly{
        fmt.Printf("%s ", v)
    }
}
```
比如說我只想知道 `value` 不想知道 `key` 就把 `key` 的位置以 `_` 代替 

另一個要注意的地方是，使用 `range` 來走訪陣列時，直接更改 `value` 值並不會對陣列內的值造成影響，但是仍可以透過鍵(key)去更改

**直接更改value**
```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for _, v := range jelly{
        v = v + v
    }
    for _, v := range jelly{
        fmt.Println(v)
    }
}
```

> 執行結果：
> 小櫻
> 知世
> 桃矢
> 雪兔
> 小櫻他爸
>

**透過key更改value**
```go=
package main
import "fmt"
func main(){
    jelly := [...]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for k, v := range jelly{
        jelly[k] = v + v
    }
    for _, v := range jelly{
        fmt.Println(v)
    }
}
```

> 執行結果：
> 小櫻小櫻
> 知世知世
> 桃矢桃矢
> 雪兔雪兔
> 小櫻他爸小櫻他爸
>

### 宣告陣列長度時不能使用變數來宣告

陣列長度在編譯完成後就不能再更動，因此在成為Golang牌前，陣列長度就必需是確定的，變數值是由Golang牌發動時才被確定的，所以不能使用變數來設定陣列的長度。Golang中的陣列是「很純的陣列」所以才會這樣規定，有些語言因為有加料所以可以實現「以變數來做為陣列長度」

```go=
package main
import "fmt"
func main(){
    length := 5
    jelly := [length]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for _, v := range jelly{
        fmt.Printf("%s ", v)
    }
}
```

> 執行結果：
> \# command-line-arguments
> .\lesson07.go:5:14: non-constant array bound length

如果你仍希望透過類似的手法來實作你可以選擇宣告一個`常數(constant)`

```go=
package main
import "fmt"
func main(){
    const length int = 5    // 利用 const 宣告一個常數
    jelly := [length]string{"小櫻", "知世", "桃矢", "雪兔", "小櫻他爸"}
    for _, v := range jelly{
        fmt.Printf("%s ", v)
    }
}
```
> 執行結果：
> 小櫻 知世 桃矢 雪兔 小櫻他爸

宣告成常數後就不可以再更動其值，因為該常數能在編譯前確定，所以可以當作陣列的長度

```go=
package main
import "fmt"
func main(){
    const length int = 5
    length = 7
    fmt.Println(length)
}
```

> 執行結果：
> \# command-line-arguments
> .\lesson07.go:5:12: cannot assign to length

## 後記
今天好想沒帶什麼劇情吔

![](https://i.imgur.com/V80C2Pf.jpg)

> 問：
> 小櫻他媽 16 歲時嫁給他爸，小櫻在 3 歲時他媽去世(當年他媽27歲)，小櫻他哥比他大 7 歲。請問小櫻他爸要關幾年？

如果喜歡這個系列希望邦友們能按個讚

本文多數圖片來自：
[圖源：庫洛魔法使第一季第十集](https://ani.gamer.com.tw/animeVideo.php?sn=10778)