package main
import(
    "fmt"
    "sync"
    "time"
)

func main(){
    var num int
    var wg sync.WaitGroup
    start:=make(chan struct{})
    done:=make(chan struct{})
    fmt.Print("Enter no. of workers:")
    fmt.Scan(&num)
    
    for i:=1;i<=num;i++{
        wg.Add(1)
        go func(id int){
            defer wg.Done()
            
            <-start
            fmt.Printf("Worker %d is starting to work\n",id)
            time.Sleep(time.Second)
            
            fmt.Printf("Worker %d has reached checkpoint\n",id)
            time.Sleep(time.Second)
            
            fmt.Printf("Worker %d has completed the work\n",id)
            time.Sleep(time.Second)
            
            
        }(i)
    }
    fmt.Println("Starting workers.....")
    close(start)
    wg.Wait()
    
    close(done)
    fmt.Println("All workers completed their work")
}
