package iradix

import (
	"fmt"
	"strings"
	"testing"
)

func TestWalk(t *testing.T) {
	r := New()
	r, _, _ = r.Insert([]byte("001"), 1)
	r, _, _ = r.Insert([]byte("002"), 2)
	r, _, _ = r.Insert([]byte("005"), 5)
	r, _, _ = r.Insert([]byte("010"), 10)
	r, _, _ = r.Insert([]byte("100"), 10)

	var got []string
	want := []string{"001", "002", "005", "010", "100"}

	r.Root().Walk(func(k []byte, v interface{}) bool {
		got = append(got, string(k))
		return false
	})

	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestWalkBackwards(t *testing.T) {
	r := New()
	r, _, _ = r.Insert([]byte("001"), 1)
	r, _, _ = r.Insert([]byte("002"), 2)
	r, _, _ = r.Insert([]byte("005"), 5)
	r, _, _ = r.Insert([]byte("010"), 10)
	r, _, _ = r.Insert([]byte("100"), 10)

	var got []string
	want := []string{"100", "010", "005", "002", "001"}

	r.Root().WalkBackwards(func(k []byte, v interface{}) bool {
		got = append(got, string(k))
		fmt.Printf("%s\n", k)
		return false
	})

	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("got %v, want %v", got, want)
	}
}
