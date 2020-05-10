package main
import "fmt"
func split(sum int) (x, y int) {    //多值返回
	x=sum % 10
	y=sum - x
	return
}

func gcd(x,y int) int{
    if y>x {
        return gcd(y,x)
    }else if x%y!=0 {
        return gcd(y,x%y)
    }
    return y
}

func lcm(x,y int) int{
    return x*y/gcd(x,y)
}

func gcd2(a int,b int) int{
    if a%b==0{
        return b
    }
    return gcd2(b,a%b)
}

func main() {
    a,b := split(15)
	fmt.Println("split(15):",a,b)
    fmt.Println("gcd(12,16):",gcd(12,16))
    fmt.Println("gcd2(12,16):",gcd(12,16))
    fmt.Println("lcm(12,16):",lcm(12,16))
}
