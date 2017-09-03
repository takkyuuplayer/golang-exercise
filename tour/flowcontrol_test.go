package tour

import "testing"
import (
	"fmt"
	"math"
	"runtime"
)

func TestForLoop(t *testing.T) {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	if fmt.Sprintf("%v", sum) != "45" {
		t.Fatal("sum up should be 45: " + fmt.Sprintf("%v", sum))
	}
}

func TestWhileLoop(t *testing.T) {
	sum := 1
	for sum < 1000 {
		sum += sum
	}

	if fmt.Sprintf("%v", sum) != "1024" {
		t.Fatal("while loop defined by for loop: " + fmt.Sprintf("%v", sum))
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func TestIfStatement(t *testing.T) {
	if fmt.Sprintf("%v", pow(2, 10, 60)) != "60" {
		t.Fatal("variable can be defined in if block: " + fmt.Sprintf("%v", pow(2, 10, 60)))
	}
}

func TestSwitch(t *testing.T) {

	switch os := runtime.GOOS; os {
	case "darwin":
		if fmt.Sprintf("%v", os) != "darwin" {
			t.Fatal("switch statement went wrong: " + fmt.Sprintf("%v", os))
		}
	case "linux":
		if fmt.Sprintf("%v", os) != "linux" {
			t.Fatal("switch statement went wrong: " + fmt.Sprintf("%v", os))
		}
	default:
		fmt.Printf("%s. ", os)
	}
}

func deferStatement() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func deferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

}
