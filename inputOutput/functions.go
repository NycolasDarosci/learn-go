package inputOutput

// fmt = format
// imports are not per package, they are per file
import "fmt"

// PrintData with uppercase, it is exported externally to the package
// public
func PrintData(n int, str string) {
	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println(Url)
	fmt.Println(Number)
	fmt.Println(n)
	fmt.Println(str)
	fmt.Println(Pi)
}

func Math(n1 int, n2 float64) float64 {

	n1AsFloat := float64(n1)

	return n1AsFloat * n2
}
