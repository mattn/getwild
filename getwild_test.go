package getwild

import (
	"reflect"
	"testing"
)

func TestSimple(t *testing.T) {
	tests := []struct {
		input []string
		want  []string
	}{
		{
			input: []string{"cmd", "-z", "*.go", "-y"},
			want:  []string{"cmd", "-z", "getwild.go", "getwild_test.go", "-y"},
		},
		{
			input: []string{"cmd", "-z", "*test*", "-y"},
			want:  []string{"cmd", "-z", "getwild_test.go", "-y"},
		},
		{
			input: []string{"cmd", "-z", "*getwild-and-chance*", "-y"},
			want:  []string{"cmd", "-z", "*getwild-and-chance*", "-y"},
		},
	}

	for _, test := range tests {
		got := getwild(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want %v, but %v:", test.want, got)
		}
	}
}
