package main
import "fmt"
import "os"

func main(){
    mymap:=make(map[string]int)
    var n int
    var s string
    var id int
    fmt.Print("Enter no. of fruits:")
    fmt.Scan(&n)
    for i:=0;i<n;i++{
        fmt.Print("Enter the fruit: ")
        fmt.Scan(&s)
        fmt.Print("Enter the id: ")
        fmt.Scan(&id)
        mymap[s]=id
    }
    var ch int
    fmt.Println("1.Update 2.Delete 3.Print 4.Exit")
    for{
    fmt.Println()
    fmt.Print("Enter your choice:")
    fmt.Scan(&ch)
    if ch==1{
        fmt.Print("Enter the fruit to be updated:")
        fmt.Scan(&s)
        _,check:=mymap[s]
        if check{
            fmt.Print("Enter new id:")
            fmt.Scan(&id)
            mymap[s]=id
        }else{
            fmt.Println("Fruit not found")
        }
    }else if ch ==2{
        fmt.Print("Enter the fruit to be deleted:")
        fmt.Scan(&s)
        _,check:=mymap[s]
        if check{
            delete(mymap,s)
        }else{
            fmt.Println("Fruit not found")
        }
        
    }else if ch==3{
        for fruit,id1:=range(mymap){
            fmt.Println(fruit,":",id1)
        }
    }else{
        os.Exit(0)
    }
    
    
    }
}
