package main
import(
    "fmt"
    "sync"
    "time"
)

type Chopstick struct{
    sync.Mutex
}

func main(){
    var numPhilosophers int
    
    fmt.Print("Enter no. of philosophers: ")
    fmt.Scan(&numPhilosophers)
    
    var wg sync.WaitGroup
    chopsticks:=make([]Chopstick,numPhilosophers)
    
    for i:=0;i<numPhilosophers;i++{
        wg.Add(1)
        go func(id int){
            defer wg.Done()
            
            fmt.Printf("Philosopher %d is thinking\n",id)
            time.Sleep(time.Second)
            left:=id
            right:=(id+1)%numPhilosophers
            
            chopsticks[left].Lock()
            chopsticks[right].Lock()
            
            fmt.Printf("Philosopher %d is eating\n",id)
            time.Sleep(time.Second)
            
            fmt.Printf("Philosopher %d finished eating\n",id)
            chopsticks[left].Unlock()
            chopsticks[right].Unlock()
            
        }(i)
    }
    wg.Wait()
}
