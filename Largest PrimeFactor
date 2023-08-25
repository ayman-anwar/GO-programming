package main
import(
    "fmt"
    "math"
)

func largestPF(num int) int{
    largest:=1
    for num%2==0{
        largest=2
        num/=2
    }
    
    for i:=3;i<=int(math.Sqrt(float64(num)));i+=2{
        for num%i==0{
            largest=i
            num/=i
        }
        
    }
    if num>2{
        largest=num
        
    }
    
    return largest
}

func main(){
    var num int
    fmt.Println("Enter a number: ")
    fmt.Scanln(&num)
    fmt.Println("The largest prime factor of ",num,"is ",largestPF(num))
}

