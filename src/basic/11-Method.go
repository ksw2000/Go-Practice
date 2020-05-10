package main
import (
    "fmt"
    "math"
)

type Vertex struct{
    X,Y float64
}

func (v Vertex) Abs() float64{
    return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v Vertex) Inner(u Vertex) float64{
    return v.X * u.X + v.Y * u.Y
}

//指標接收器
func (v *Vertex) Add(u Vertex) Vertex{
    v.X += u.X
    v.Y += u.Y
    return *v
}

func main(){
    v := Vertex{3, 4}
    fmt.Println(v.Abs());

    w := Vertex{6, 7}
    fmt.Println(w, "dot", "{1, -2} =", w.Inner(Vertex{1, -2}))

    z := &Vertex{1,2}
    fmt.Println(z.Add(Vertex{3, 8}))
}
