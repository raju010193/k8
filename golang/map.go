package main
import "fmt"
func main(){
students := make(map[int]string)
students[1] = "sidhu"
students[2] = "swamy"
students[3] = "raju"
fmt.Println(students)
fmt.Println(len(students))
delete(students,2)
fmt.Println(students)
for k,v := range students{
    fmt.Printf("%d - %s\n",k,v)
}
}