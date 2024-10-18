package sort

import(
	"testing"
	"reflect"
)

func TestInsertion(t *testing.T) {
	got := Insertion([]int{5, 4, 3, 2, 1, 0})
	want := []int{0, 1, 2, 3, 4, 5}
	
	if ! reflect.DeepEqual(got, want) {
		t.Log("Got: ", got)
		t.Log("Want: ", want)
		t.Fatalf("Array did not sorted propertly")
	}
}