package closure

//闭包与函数字面量
import "fmt"

func init() {
	fmt.Println("init package closure")
}

func Closure(x int) func(int) int {
	return func(i int) int {
		x += i
		return x
	}
}
