---
tags: Golang魔法使
---
# \#3 字串與Printf | Golang魔法使

![](https://i.imgur.com/NE34mcw.jpg)

## 前言

這次小櫻前往{水族館|aquarium}校外教學，卻不幸遇到可惡的「字串」Golang牌，導致水族館員和企鵝差點被淹死，小櫻的哥哥恰好在這裡打工，縱身一躍跳進水池解救了大家，像這樣

![](https://i.imgur.com/gCenn8T.jpg)

要如何把這個囂張的 Golang 牌繩之以法呢？「知己知彼百戰百勝」，螢幕前小櫻的工具人，一起來了解字串 Golang 牌的運作吧！

## 了解字串前要先了解什麼是字元

在了解字串前我們必需先了解什麼是字元，字元其實就是指「單一一個字」，這個字也是由數字所組成，在 Go 語言中是以 UTF-8 來做為字元的編碼方式，舉個例子，小櫻的「小」在是以 23567 的二進位來儲存「櫻」則是以 27387 的二進位來儲存

要宣告一個字元我們會用 `rune` (中文：符文)來表示，`rune` 跟 `int32` 是一樣的東西

```go=
package main
import "fmt"
func main(){
    var a rune = '小' // 以單引號包住字元
    var b rune = '櫻'
    fmt.Println(a)
    fmt.Println(b)
}

```

> 執行結果：
> 23567
> 27387
> 

**為什麼會這樣？**

因為 `rune` 就是 `int32`，通過 `fmt.Println` 當然會把他當作數字來做處理，如果我們要以「文字」印出，則必需使用 `fmt.Printf` ， `printf` 指的就是 `print` + `format` ，`format`(格式)決定要以什麼形式印出，使用方法如下：

```go=
package main
import "fmt"
func main(){
    var a rune = '小'
    var b rune = '櫻'
    fmt.Printf("%c", a)
    fmt.Printf("%c", b)
}
```

其中的 `%c` 就是告訴魔仗我要印字元出來，至於要印哪個字元呢？則是看 `,` 後面第一個變數

> 執行結果：
> 小櫻
> 

當然也可以用 `%c %c` 他會依序對應 `,` 後面的變數

```go=
package main
import "fmt"
func main(){
    var a rune = '小'
    var b rune = '櫻'
    fmt.Printf("%c%c", a, b)
}
```

> 執行結果：
> 小櫻
> 


> 另外常用的還有 `%d`, `%f`, `%s`
>
> `%c` 的 c 是指 character 以字元印出
>
> `%d` 的 d 是指 decimal 以十進位印出
>
> `%e` 的 e 是指 exponent 以科學計號方式印出
>
> `%f` 的 f 是指 float 以浮點數印出
> 
> `%s` 的 s 是指 string 以字串印出
> 
> `%%` 是用來印出 %


## 單引號其實是將「文字」轉為「數字」

當我們使用 `'小'` 這個用法時，電腦會直接把他視為 `23567` 你甚至可以試試把 `'小'` 減去 `23566` 讓他變成 `1`

```go=
package main
import "fmt"
func main(){
    var a rune = '小' - 23566
    fmt.Printf("%d", a)
}

```

> 執行結果：
> 1
>

`'小' - 23566` 對於Go語言來說會直接視為 `23567-23566` 比如：`fmt.Printf("%d", a)` 則視為 `fmt.Printf("%d", 1)`

## 字串是什麼？

相信大家都知道字元是什麼了，「字元不過就是個數字」，那「字串」呢？字串其實就是一串「字元」。(詳細原理會再後續課程說明)

字串怎麼使用呢？舉個例子大家應該就懂了！

```go=
package main
import "fmt"
func main(){
    var release string = "封印解除"
    release2 := "レリーズ"
    fmt.Println(release)
    fmt.Println(release2)
}
```

> 執行結果：
> 封印解除
> レリーズ
> 


如同整數一樣，字串也元運算符號，也就是「加法」，在資訊科學(~~魔法科學~~)中稱為 `concatenate`

```go=
package main
import "fmt"
func main(){
    str1 := "隱藏著黑暗力量的鑰匙啊"
    str2 := "在我面前顯示你真正的力量"
    fmt.Println(str1 + str2)
}
```

> 執行結果：
> 隱藏著黑暗力量的鑰匙啊在我面前顯示你真正的力量
> 

## 如何將數字型態 int 轉成字串型態 string

如果我們只是要印出數字可以用 Printf 來印出

**法一：Printf**
```go=
package main
import "fmt"
func main(){
    sakura := "小櫻今年"
    grade := "年級"
    fmt.Printf("%s%d%s", sakura, 4, grade)
}
```
> 執行結果：
> 小櫻今年4年級
> 

**法二：Print**
```go=
package main
import "fmt"
func main(){
    sakura := "小櫻今年"
    grade := "年級"
    fmt.Print(sakura, 4, grade)
}
```
> 執行結果：
> 小櫻今年 4 年級
> 

`fmt.Print(sakura, 4, grade)` 可以視為

```go=
fmt.Print(sakura)
fmt.Print(" ")
fmt.Print(4)
fmt.Print(" ")
fmt.Print(grade)
```


以上提到的方法都是直接印在螢幕上(標準輸出 standard output)。那有沒有方法直接將結果印在變數裡？

**法三：使用 Sprintf**
```go=
package main
import "fmt"
func main(){
    sakura := "小櫻今年"
    grade := "年級"
    result := fmt.Sprintf("%s%d%s", sakura, 4, grade)
    fmt.Printf(result)
}
```

這個用法會將原本應該印出於螢幕上的字串印在變數 result 當中

## 字串轉成數字？

如果今天我是想把 `"4"` (字串) 轉成 `4` (整數)那要怎麼轉呢？
其實這個可以把字串拆成字元再用字元去做運算，但是剛剛就說了，依現在各位魔法使的實力，還不夠，所以我們就先跳過吧！

## print 使用方法整理

Go 語言常用的印出方法可以簡單分為三種 `fmt.Print`, `fmt.Println`, `fmt.Printf`

其中，`fmt.Println` 與 `fmt.Print` 差不多，只是 `fmt.Println` 會在每次印出時都會在行尾做換行。`fmt.Printf` 可以指定以哪 format 印出，而其他兩者都是以預設 format 印出

<!--
## 與其他程式語言的比較

初學者可以跳過

### C 語言

剛剛說了 Golang 在字串加數字的一些邏輯，我們來看一下其他語言

在 C 語言中我們可以用一樣`sprintf`的方法來做，要注意的事第一個參數放的事有足夠大小的字串指標

> `sprintf()` 如果成功，則返回寫入的字符的總數，不包括`\0`，否則在發生故障的情況下，返回一個負數。

```c
#include<stdio.h>
int main(){
    char sakura[64] = "小櫻今年";
    char* grade = "年級";
    sprintf(sakura, "%s%d%s", sakura, 4, grade);
    printf(sakura);
    return 0;
}
```

### javascript
在 `javascript` ES5 之前我們可以把字串直接與數字相加，在進行字串與數字相加時 javascirpt 會直接把數字做為字串處理

```javascript
var sakura = "小櫻今年";
var grade = "年級";
var result = sakura + 4 + grade;
console.log(result);
```

而 ES6 則可以直接將變數嵌入到字串中

```javascript
var sakura = "小櫻今年";
var grade = "年級";
var result = `${sakura} 4 ${grade}`;
console.log(result);
```

### php
在 `php` 中字串相加用 `.` 數字用 `+` ，缺點是物件導向時只能用 `->` 不能用 `.` 但 php 一開始設計時就不是要走物件導向所以一開始問題不大，現在要用物件才發現有夠麻煩 `.` 要打 `->`

另一種方法是把變數鑲進雙引號裡面，這感覺跟 `dart` 和新版的 `javascript` 有點類似

```php=
<?php
    $sakura = "小櫻今年";
    $grade = "年級";
    $result = $sakura.(4).$grade;
    echo $result;
    
    $result = "{$sakura}4{$grade}";
    echo $result;
?>
```

### java, python
在 `java` 中就好玩了，雖然大家都覺得`java` 是強型態，`python` 是弱型態，但是 `java` 卻能支援字串加數字，因為 java 在做相加時會自動將 4 轉成字串

```java=
public class Strcat{
    public static void main(String[] args){
        String sakura = "小櫻今年";
        String grade = "年級";
        String result = sakura + 4 + grade;
        System.out.println(result);
    }
}
```

```python=
sakura = "小櫻今年"
grade = "年級"
# print(sakura + 4 + grade) 報錯
print(sakura + str(4) + grade)
```

### dart

dart 在數字加字串時習慣會直接像用 php 的那種方法將變數嵌入

```dart=
void main(){
    var sakura = "小櫻今年";
    var grade = "年級";
    var number = 4;
    print(sakura + number.toString() + grade);
    print("$sakura$number$grade");
}
```

### rust
在 rust 中則是類似用 sprintf() 的做法來實現
```rust=
fn main(){
    let sakura = "小櫻今年";
    let grade = "年級";
    let result = format!("{}{}{}", sakura, 4, grade);
    print!("{}", result);
}
```

rust 中的 string 挺複雜的

## kotlin

!-->

---

## 後記

這部分2021年被我刪減掉一些，感覺之前教太難:smile:

圖片大多來自：[庫洛魔法使第一季第三集](https://ani.gamer.com.tw/animeVideo.php?sn=10771)以及[庫洛魔法使第一季第四集](https://ani.gamer.com.tw/animeVideo.php?sn=10771)