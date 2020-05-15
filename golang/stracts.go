package main
import ("fmt"
        "strconv")

type Person struct{
//   firstName string
//   lastName string
//   city string
//   gender string
//   age int
firstName,lastName,city, gender string
age                             int
}

func (p *Person)hasBirthday(){
    p.age++
}
func (p Person)greet()string{
    return "Hello, My name is "+ p.firstName+" "+p.lastName+" "+strconv.Itoa(p.age)
}
func main(){
person1 := Person{firstName:"sidhu",lastName:"munagala", city:"vijayawada",age:26,gender:"male"}
fmt.Println(person1)
person1.hasBirthday()
person1.hasBirthday()
person1.hasBirthday()
person1.hasBirthday()


fmt.Println(person1.greet())

}