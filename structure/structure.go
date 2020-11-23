package structure

//结构体和结构体方法
//interface
import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender byte
}

type Methods interface {
	FmtPersion()
}

func (person *Person) FmtPersion() {
	fmt.Println("name=", person.Name)
	fmt.Println("age=", person.Age)
	fmt.Println("gender=", person.Gender)
}
