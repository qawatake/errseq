package erriter_test

import (
	"erriter"
	"errors"
	"fmt"
)

func ExampleNew() {
	s := erriter.New(func(yield func(int) bool) error {
		for i := range 5 {
			if i == 3 {
				return errors.New("3!!!")
			}
			if !yield(i) {
				return nil
			}
		}
		return nil
	})

	var err error
	for v := range s(&err) {
		fmt.Println(v)
	}

	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("No error")
	}

	// Output:
	// 0
	// 1
	// 2
	// Error: 3!!!
}
