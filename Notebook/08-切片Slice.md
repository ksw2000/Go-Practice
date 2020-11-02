---
tags: Golang魔法使
---
# \#8 切片 Slice | Golang 魔法使

![](https://i.imgur.com/kUcHx04.jpg)


昨天介紹的陣列有一個致命缺點，就是他的長度是固定的不能改變

而今天要來介紹一個比陣列更有彈性的組合型別──切片。

但也因為比較彈性，所以要學的部份就比較多


# 切片 Slice

## 切片原理

切片是由三個部分組成而成

1. 陣列
2. 容量 capacity
3. 長度 length

直接看例子：

```go=
package main
import "fmt"
func main(){
    role  := [...]string{"小櫻", "知世", "小可", "小狼", "苺鈴"}
    slice := role[0:3]          // [0] [1] [2]
    fmt.Println(slice)
    fmt.Println(role)           // 陣列
    fmt.Println(len(slice))     // 長度
    fmt.Println(cap(slice))     // 容量
}
```

> 執行結果：
> \[小櫻 知世 小可\]
> \[小櫻 知世 小可 小狼 苺鈴\]
> 3
> 5

其中重點就在 `role[0:3]` 這到底是什麼意思呢？這個就是只我在 0 的左邊切一刀 3 的左邊切一刀

> `切` \[0 "小櫻"\]\[1 "知世"\]\[2 "小可"\]`切`\[3 "小狼"\]\[4 "苺鈴"\]

這樣 `slice` 就會從 `role` 裡面挑出 0, 1, 2 形成切片

而所謂的容量就是指原本 `role` 的長度：`5`個角色，而長度則是指切出來的長度`3`

利用指向不同的陣列，可以實現容量、長度都可以變動的切片

## 宣告一個有預設長度、容量的切片

除了直接從陣列切來，也可以用 `make` 來宣告一個新的切片

其中第一個參數是切片的型態比如 `[]int` `[]string`

第二個參數填入長度 len ，第三個參數填入容量 cap

其中**長度不可超過容量**，容量是底層陣列的長度，一般情況長度是可以輕易改變的，若要改變容量則必需重新更換底層陣列

```go=
package main
import "fmt"
func main(){
    slice := make([]string, 3, 5)
    slice[0] = "小櫻"
    slice[1] = "知世"
    slice[2] = "小可"
    fmt.Println(slice)
    fmt.Println(len(slice))     // 長度
    fmt.Println(cap(slice))     // 容量
}
```

> 執行結果：
> 
> [小櫻 知世 小可]
> 3
> 5

## 宣告一個長度與容量相同的切片

在宣告時可以省略第三個參數，如果僅指定第二個參數，那麼就代表宣告一個長度與容量相同的切片

```go=
package main
import "fmt"
func main(){
    slice := make([]string, 5)
    slice[0] = "小櫻"
    slice[1] = "知世"
    slice[2] = "小可"
    fmt.Println(slice)
    fmt.Println(len(slice))     // 長度
    fmt.Println(cap(slice))     // 容量
}
```

> 執行結果：
> 
> [小櫻 知世 小可 ]
> 5
> 5

## 快速宣告一個切片

除了用 `make` 宣告還可以用向陣列一樣宣告的方法

```go=
package main
import "fmt"
func main(){
    slice := []string{"小櫻", "知世", "小可"}
    fmt.Println(slice)
    fmt.Println(len(slice))     // 長度
    fmt.Println(cap(slice))     // 容量
}
```

> 執行結果：
> 
> [小櫻 知世 小可]
> 3
> 3

要注意的是，宣告陣列時是用 `[...]string{}` 而使用切片時是用 `[]string{}`

## append()

宣告切片時，如果不確定要用多少容量，可以先宣告一個空切片，並且使用append() 來插入新的值。在使用 append() 時不需要考慮切片的容量，golang會自己解決，這就是切片好用的地方

```go=
package main
import "fmt"
func main(){
    slice := []string{}         // 宣告空切片
    slice = append(slice, "小櫻")
    slice = append(slice, "知世")
    slice = append(slice, "小可")

    fmt.Println(slice)
    fmt.Println(len(slice))     // 長度
    fmt.Println(cap(slice))     // 容量(append 時容量不夠會自動增加)
}
```

## 走訪

跟陣列一樣 slice 也可以利用 `for` 進行走訪
```go=
package main
import "fmt"
func main(){
    slice := []string{"小櫻", "知世", "小可"}
    for i := 0; i<len(slice); i = i+1{
        fmt.Printf("%d : %s\n", i, slice[i])
    }
}
```

跟陣列一樣 slice 也可以利用 `for + range` 進行走訪

```go=
package main
import "fmt"
func main(){
    slice := []string{"小櫻", "知世", "小可"}
    for k, v := range slice{
        fmt.Printf("%d : %s\n", k, v)
    }
}
```

> 執行結果：
> 0 : 小櫻
> 1 : 知世
> 2 : 小可
>

如果不使用 `value` 值則規則與走訪陣列一致，將 `value` 設成 `_` 即可

## 後記

小櫻真嗨

![](https://i.imgur.com/C2BokxD.jpg)

本文多數圖片來自：
[圖源：庫洛魔法使第一季第十二集](https://ani.gamer.com.tw/animeVideo.php?sn=10780)
