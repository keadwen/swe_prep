package sorts

import (
	"reflect"
	"testing"
)

func TestInsertion(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{{
		name: "successful sort",
		data: []int{5, 8, 3, 2, 11},
		want: []int{2, 3, 5, 8, 11},
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Insertion(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Insertion(): got %v, want %v", got, tt.want)
			}
		})
	}
}
