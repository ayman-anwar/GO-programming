package main
import(
    "fmt"
    "math"
)

type shape interface{
    Area() float64
}

type circle struct{
    radius float64
}

func (c circle) Area() float64{
    return math.Pi*c.radius*c.radius
}

type rectangle struct{
    width float64
    height float64
}

func (r rectangle) Area() float64{
    return r.height*r.width
}

func printArea(s shape){
    fmt.Printf("Area:%0.2f\n",s.Area())
}

func main(){
    var r float64
    var h float64
    var w float64
    fmt.Print("Enter the radius: ")
    fmt.Scan(&r)
    circle:=circle{radius:r}
    printArea(circle)
    
    fmt.Print("Enter height:")
    fmt.Scan(&h)
    fmt.Print("Enter width:")
    fmt.Scan(&w)
    rectangle:=rectangle{height:h,width:w}
    printArea(rectangle)
}
