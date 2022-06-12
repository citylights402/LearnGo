package learn

import (
	// "demo1/packa"
	"fmt"
	"go_code/src/demo1/packa"
)

func main() {
	i := 100
	ptr := &i

	fmt.Println(i)           //100
	fmt.Println(ptr)         //0xc0000aa058
	fmt.Println(*ptr)        //100
	fmt.Printf("%p\n", &ptr) //0xc000006028

	*ptr = 200
	fmt.Println(i)

	fmt.Println(packa.Variable2)
	// packa.variable1

}
