---
tags: Golang魔法使
---
# \#12 方法 Method & 空指標 nil | Golang魔法使
這天是暑假的最後一天，晚上在月峰神社有慶典，千春和唬爛王山崎是在閃屁哦

![](https://i.imgur.com/aVD44Pf.jpg)


接下來的課程對於沒有接觸過物件導向的魔法使們可能會吃力一點點，但也不太會，因為 Golang 並不算是物件導向的語言，所以 Golang 只會用到一些些而已，基本上是很好學的 ~~文組不好說~~

![](https://i.imgur.com/bSrWg11.jpg)

收服這張 Method Golang 牌就可以被 **國民女兒知世** 稱讚了嗎？(我記得前幾天好像說是國民老婆但先不要，年紀太小會被說是蘿莉控)

## 什麼是 Method？

一般來說，我們常用的函式是將參數放在函式的名稱後面比如 `func double(param int){}`，但如果現在我有一套函式，常常都對同一個參數使用時，我們可以透過「Method」讓咒語看起來更簡單

### 不使用 Method

請魔法使們先觀察以下的 Golang 牌

```go=
package main
import "fmt"

type person struct{
    name   string
    groups []string
    crush  *person
}

func getName(p *person) string{
    return p.name
}

func setCrush(p *person, crush *person){
    p.crush = crush
}

func getCrushName(p *person) string{
    return p.crush.name
}

func main(){
    sakura := &person{
        name   : "木之本櫻",
        groups : []string{"啦啦隊"},
    }

    tomoyo := &person{
        name   : "大道寺知世",
        groups : []string{"合唱團"},
    }

    touya := &person{
        name : "木之本桃矢",
    }

    yukito := &person{
        name : "月城雪兔",
    }

    setCrush(sakura, yukito)
    setCrush(tomoyo, sakura)
    setCrush(touya, yukito)
    setCrush(yukito, touya)

    fmt.Println("小櫻的本名是？", getName(sakura))
    fmt.Println("小櫻喜歡的人是？", getCrushName(sakura))

    fmt.Println("知世的本名是？", getName(tomoyo))
    fmt.Println("知世喜歡的人是？", getCrushName(tomoyo))
}
```

> 執行結果：
> 小櫻的本名是？ 木之本櫻
> 小櫻喜歡的人是？ 月城雪兔
> 知世的本名是？ 大道寺知世
> 知世喜歡的人是？ 木之本櫻

相信大家應該都有看得懂，畢竟大家已經當了11天的魔法使的工具人了，可能比較不懂的是「為什麼他們喜歡的人是那樣配」，大家回去補第一季就知道了 (第二季會換)

### 使用 Method 更加好懂

很明顯地可以發現剛剛給的函式 **幾乎都會用到** `p *person` **這一個參數** ，這是個使用 Method 的好時機

```go=
func getName(p *person) string{
    return p.name
}
//可以換成
func (p *person) getName() string{
    return p.name
}
```

```go=
// 原先是這樣呼叫：
getName(sakura)
// 現在是這樣呼叫：
sakura.getName()
```

所以上一課才說比起 `struct{}` 我們更喜歡用 `*struct{}`

**使用Method改寫**

```go=
package main
import "fmt"

type person struct{
    name    string
    groups  []string
    crush   *person
}

func (p *person) getName() string{
    return p.name
}

func (p *person) getGroups() []string{
    return p.groups
}

func (p *person) setCrush(crush *person){
    p.crush = crush
}

func (p *person) getCrushName() string{
    return p.crush.name
}

func main(){
    sakura := &person{
        name   : "木之本櫻",
        groups : []string{"啦啦隊"},
    }

    tomoyo := &person{
        name   : "大道寺知世",
        groups : []string{"合唱團"},
    }

    touya := &person{
        name : "木之本桃矢",
    }

    yukito := &person{
        name : "月城雪兔",
    }

    sakura.setCrush(yukito)
    tomoyo.setCrush(sakura)
    touya.setCrush(yukito)
    yukito.setCrush(touya)

    fmt.Println("小櫻的本名是？", sakura.getName())
    fmt.Println("小櫻喜歡的人是？", sakura.getCrushName())

    fmt.Println("知世的本名是？", tomoyo.getName())
    fmt.Println("知世喜歡的人是？", tomoyo.getCrushName())
}
```

> 執行結果：
> 小櫻的本名是？ 木之本櫻
> 小櫻喜歡的人是？ 月城雪兔
> 知世的本名是？ 大道寺知世
> 知世喜歡的人是？ 木之本櫻
> 

---
## 中場休息

來自Instagram插畫家 [@babyfat_yuan](https://www.instagram.com/babyfat_yuan/) 的作品

![](https://i.imgur.com/jFzPyA7.jpg)

![](https://i.imgur.com/KiGYgyd.jpg)

![](https://i.imgur.com/xBkIJrv.jpg)

![](https://i.imgur.com/0LvoMAo.jpg)


---

## 什麼是 nil ── 知世：小櫻，那個是 nil 啊！

nil 在一般的程式語言其實就是 `null`, `NULL`, 或 `None`

先前提到說，在宣告一個變數時Golang會給一個預設值，比如 int, uint 會給 `0`, string 會給 `""`

那麼如果是一個指標變數，會給什麼呢？Golang會給他一個全部都是 0 的位址，稱為 `nil` 要注意的是，`nil` 常常會導致程式出錯。就舉先前的程式為例，如果有一個角色沒有喜歡的人

```go=
package main
import "fmt"

type person struct{
    name    string
    groups  []string
    crush   *person
}

func (p *person) getName() string{
    return p.name
}

func (p *person) setCrush(crush *person){
    p.crush = crush
}

func (p *person) getCrushName() string{
    return p.crush.name
}

func main(){
    sakura := &person{
        name   : "木之本櫻",
        groups : []string{"啦啦隊"},
    }

    tomoyo := &person{
        name   : "大道寺知世",
        groups : []string{"合唱團"},
    }

    naoko := &person{
        name : "柳澤奈緒子",
        groups : []string{"啦啦隊"},
    }

    touya := &person{
        name : "木之本桃矢",
    }

    yukito := &person{
        name : "月城雪兔",
    }

    // naoko 並沒有設定喜歡的人
    sakura.setCrush(yukito)
    tomoyo.setCrush(sakura)
    touya.setCrush(yukito)
    yukito.setCrush(touya)

    // 對 naoko 存取喜歡的人的名字會...？
    fmt.Println("奈緒子的本名是？", naoko.getName())
    fmt.Println("奈緒子喜歡的人是？", naoko.getCrushName())
}
```


> 執行結果：
> 奈緒子的本名是？ 柳澤奈緒子
> panic: runtime error: invalid memory address or **nil pointer** dereference
> \[signal 0xc0000005 code=0x0 addr=0x8 pc=0x49136d\]
>
> goroutine 1 [running]:
> main.(*person).getLoverName(...)
> 省略...
> 


為什麼會這樣呢？因為在預設情況下，宣告新的 person 時， lover 會預設為 `nil`，`nil` 就是一個沒有指向任何地方的指標，因此如果你想要取得 `nil.name` 當然會噴錯

因為這張Golang牌很小，所以一下就可以找到錯誤了，如果是一張幾百行咒語的Golang牌，會很難找到錯誤的地方

那麼我們要怎麼避免這種錯誤呢？那就是在用 `.` 的時候「小心小心再小心」

![](https://i.imgur.com/LRT8Mea.png)
[庫洛魔法使第一季第六集](https://ani.gamer.com.tw/animeVideo.php?sn=10774)

改寫 `getLoverName()`：
```go=
func (p *person) getCrushName() string{
    if p.lover == nil{
        return "沒有"
    }
    return p.crush.name
}
```

> 執行結果：
> 奈緒子的本名是？ 柳澤奈緒子
> 奈緒子喜歡的人是？ 沒有
> 

## Method 的限制

什麼樣的變數可以擁有 Method？

要注意的是，我們只能對自己定義的型態新增方法，而且只能對同一個 package 中的型態做定義。因為還沒介紹 package 的概念所以現階段都是在同一個稱為 main 的 package，未來才會介紹怎麼創建另一個 package

如果我們對Golang本身內建的型態int新增一個取絕對值的方法：

```go=
package main
import "fmt"

func (n int) absoulte() int{
    if n < 0{
        return -n
    }
    return n
}

func main(){
    num := 3
    fmt.Println(num.absoulte())
}
```

> 執行結果：
> \# command-line-arguments
> .\lesson12b.go:4:6: cannot define new methods on non-local type int
> .\lesson12b.go:13:20: num.absoulte undefined (type int has no field or method absoulte)
> 

如果你還是堅持想要實現對 int 新增方法，該怎麼實現？其實可以透過`型態別名`來新增，因為使用別名會被視為是「我們自定的型態」，而且現在都是在同一個 package (main package)下使用(未來講解)

新增一個 int 的別名為 myInt：

```go=
package main
import "fmt"

type myInt int

func (n myInt) absoulte() myInt{
    if n < 0{
        return -n
    }
    return n
}

func main(){
    var num myInt // 第 14 行
    num = -3
    fmt.Println(num.absoulte())
}
```

> 執行結果：
> 3
> 

另外要注意的是：以上的範例是用 **非指標型態** 定義方法，這樣呼叫 `absoulute()` 時，僅僅是回傳經過更改的值，而原本 **第 14 行** 的 `num` 並沒有被更改。

我們可以試著印出 num

```go=
package main
import "fmt"

type myInt int

func (n myInt) absoulte() myInt{
    if n < 0{
        return -n
    }
    return n
}

func main(){
    var num myInt // 第 14 行
    num = -3
    fmt.Println(num.absoulte())
    // num 實際上是沒有被更改的   
    fmt.Println(num)
}
```

> 執行結果：
> 3
> -3
> 

如果我們希望利用 `.absoulut()` 就能更改 **第 14 行** 的變數，那麼就必需改對「指標變數」來新增方法：

```go=
package main
import "fmt"

type myInt int

func (n *myInt) absoulte() myInt{
    if *n < 0{
        *n = -(*n)
    }
    return *n
}

func main(){
    var num myInt = -3 // 第 14 行
    pointer := &num
    fmt.Println(pointer.absoulte())
    fmt.Println(num)
}
```

> 執行結果：
> 3
> 3

如此一來 **第 14 行** 的 num 就會被更改。

## 後記

各位醉生夢死的大學生們，差不多該醒了哦

![](https://i.imgur.com/Vt71uzE.jpg)

不過我們大學 9/7 就開學了，9/11 上課時教授 14:30 手機鬧鐘響，看來大資工系的生活都相當有規律呢！

剛剛想說，小櫻這樣不就小5了嗎？bear bear 忘記日本是隔年春天才算一年，所以這一季的小櫻還是小4，不過這是20年前的動漫，換算一下現在應該 30 多歲了

祝大家開學愉快

本文多數圖片來自：
[庫洛魔法使第一季第十八集](https://ani.gamer.com.tw/animeVideo.php?sn=10786)