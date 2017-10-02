package listy_test

import (
	"fmt"

	"github.com/xdg/listy"
)

func Example() {

	strings := []string{"foo", "bar", "bar", "baz"}
	ints := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}

	// Wrap slices in listy types to "box" them

	xs := listy.S(strings)
	ys := listy.I(ints)

	// Call methods on boxed types. Chain them as needed.

	if xs.Contains("baz") {
		fmt.Println("Has 'baz'")
	}

	if ys.Contains(9) {
		fmt.Println("Has '9'")
	}

	uxs := xs.Uniq()
	uys := ys.Uniq()

	mxs := uxs.Reverse().Map(func(s string) string { return fmt.Sprintf("'%s'", s) })
	mys := uys.Reverse().Map(func(i int) int { return -1 * i })

	// Unbox them to get built-in slice types back

	fmt.Println("Unique xs:", uxs.Unbox())
	fmt.Println("Unique ys:", uys.Unbox())

	fmt.Println("Reverse mapped unique xs:", mxs.Unbox())
	fmt.Println("Reverse mapped unique ys:", mys.Unbox())

	// Output:
	// Has 'baz'
	// Has '9'
	// Unique xs: [foo bar baz]
	// Unique ys: [3 1 4 5 9 2 6]
	// Reverse mapped unique xs: ['baz' 'bar' 'foo']
	// Reverse mapped unique ys: [-6 -2 -9 -5 -4 -1 -3]
}
