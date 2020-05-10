package main
import "fmt"
func main(){
    for i:=2; i<10 ;i++ {
        for j:=2; j<10 ; j++ {
            fmt.Printf("%dx%d=%-2d ",i,j,i*j)
        }
        fmt.Printf("\n")
    }

    for i:=1; i<=19 ;i++ {
        fmt.Println("2x",i,"=",2*i)
    }

    //GOlang use "for" instead of "while"
    /*
    C:
        int i=0;
        while(i<10){
            printf("%d ",i);
            i++;
        }
    */
    i:=0;
    for i<10 {
        fmt.Printf("%d ",i)
        i++
    }
}
