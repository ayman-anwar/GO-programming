package main
import "fmt"
func main(){
    var num int
    a:=0
    b:=1
    fmt.Printf("Enter a number:")
    fmt.Scan(&num)
    copy:=0
    for i:=0;i<num;i++{
        if (a%2==0){
            fmt.Println(a)
        }
        copy=a+b
        a=b
        b=copy
    }
}
