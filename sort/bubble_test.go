package sort

import(
	"testing"
	"reflect"
)

func TestBubble(t *testing.T){
	got := Bubble([]int{3, 2, 1, 0})

	want := []int{0, 1, 2, 3}

	if ! reflect.DeepEqual(got, want) {
		t.Fatalf("Array did not sorted propertly")
	}
}
