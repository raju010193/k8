package main
import "fmt"

func main(){
    ids := []int{23,234,34,55,66,44,33,56,77,99}
    for i,ids := range(ids){
    fmt.Printf("%d - ID's %d\n",i,ids)
    }
    sum :=0
    for _,id := range(ids){
    sum+=id
    }
    fmt.Println(sum)
}