package main
import "fmt"

func isPalindrome(num int) bool{
    rev:=0
    temp:=num
    for temp!=0{
        rev=rev*10+temp%10
        temp/=10
    }
    return rev==num
}

func largestPalindrome()(int,int,int){
    maxP:=0
    n1:=0
    n2:=0
    product:=0
    for i:=999;i>=100;i--{
        for j:=999;j>=100;j--{
            product=i*j
            if product>=maxP && isPalindrome(product){
                maxP=product
                n1=i
                n2=j
            }
        }
    }
    return maxP,n1,n2
}

func main(){
    large,num1,num2:=largestPalindrome()
    fmt.Println(large)
    fmt.Println(num1,num2)
}
