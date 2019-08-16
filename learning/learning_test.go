package learning

import (
	"fmt"
	"testing"
)

func Test_learning(t *testing.T) {
	fmt.Printf("%b\n", (1<<16)-1)
	fmt.Printf("%d\n", (1<<16)-1)
}

func Test_uints(t *testing.T) {
	var unsigned1, unsigned2 int
	unsigned1, unsigned2 = 50, 2

	result := int(unsigned2 - unsigned1)
	fmt.Println("result: ", result)

}
