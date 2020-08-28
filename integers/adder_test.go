package integers

import (
	"fmt"
	"testing"
)

// godoc -http=:6060
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	if got != want {
		// %q for strings in quotes and %d for digits
		t.Errorf("Got %d but Want is %d", got, want)
	}
}
