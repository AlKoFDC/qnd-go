package funding

import (
	"fmt"
	"testing"
)

func ExampleSwap() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Output:
	// world hello
}

func TestSwap(t *testing.T) {
	type testIO struct {
		in1     string
		in2     string
		expect1 string
		expect2 string
	}
	for _, value := range []testIO{
		{in1: "world", expect2: "world"},
		{"hello", "", "", "hello"},
		{"sauce", "awesome", "awesome", "sauce"},
	} {
		got1, got2 := swap(value.in1, value.in2)
		if got1 != value.expect1 || got2 != value.expect2 {
			t.Errorf(
				"For \"%s %s\" expected to get \"%s %s\", but got \"%s %s\".",
				value.in1, value.in2, value.expect1, value.expect2, got1, got2,
			)
		}
	}
}

func TestMe(t *testing.T) {
	for _, v := range map[int]int{
		1: 4,
		2: 5,
		3: 6,
	} {
		t.Log(v)
	}
	t.Log("slice")
	for _, v := range []int{10, 9, 8, 7, 6} {
		t.Log(v)
	}
}

func TestSplit(t *testing.T) {
	var a, b int
	const (
		expectA = 7
		expectB = 10
	)
	a, b = split(17)
	t.Log(a, b)
	if a != expectA {
		t.Error("a not as expected")
	}
	if b != expectB {
		t.Fatalf("b not as expected: got %d expect %d", b, expectB)
	}
}

func TestMap(t *testing.T) {
	type Vertex struct {
		Lat, Long float64
	}

	//m := make(map[string]Vertex)

	m := map[string]Vertex{"Bell Labs": {1, 1}, "Google": {37.42202, -122.08408}}
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	t.Log(m["Bell Labs"])
}
