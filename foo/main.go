// Package foo is nothing interesting
package foo

import "fmt"

// Foo uses the range over int feature.
func Foo() {
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("go1.22 has lift-off!")
}
