package iradix

import (
	"strings"
	"testing"
)

func TestIteratorNext(t *testing.T) {
	r := New()
	r, _, _ = r.Insert([]byte("001"), 1)
	r, _, _ = r.Insert([]byte("002"), 2)
	r, _, _ = r.Insert([]byte("005"), 5)
	r, _, _ = r.Insert([]byte("010"), 10)
	r, _, _ = r.Insert([]byte("100"), 10)

	var got []string
	want := []string{"001", "002", "005", "010", "100"}

	it := r.Root().Iterator()
	for key, _, ok := it.Next(); ok; key, _, ok = it.Next() {
		got = append(got, string(key))
	}

	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("wrong traversal order: got: %v, want: %v", got, want)
	}
}

func TestIteratorPrev(t *testing.T) {
	r := New()
	r, _, _ = r.Insert([]byte("001"), 1)
	r, _, _ = r.Insert([]byte("002"), 2)
	r, _, _ = r.Insert([]byte("005"), 5)
	r, _, _ = r.Insert([]byte("010"), 10)
	r, _, _ = r.Insert([]byte("100"), 10)

	var got []string
	want := []string{"100", "010", "005", "002", "001"}

	it := r.Root().Iterator()
	for key, _, ok := it.Prev(); ok; key, _, ok = it.Prev() {
		got = append(got, string(key))
	}

	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("wrong traversal order: got: %v, want: %v", got, want)
	}
}

func TestIteratorChangeDirection(t *testing.T) {
	r := New()
	r, _, _ = r.Insert([]byte("001"), 1)
	r, _, _ = r.Insert([]byte("002"), 2)
	r, _, _ = r.Insert([]byte("005"), 5)
	r, _, _ = r.Insert([]byte("010"), 10)
	r, _, _ = r.Insert([]byte("100"), 10)

	it := r.Root().Iterator()
	it.Next()

	if _, _, ok := it.Prev(); ok {
		t.Errorf("Iterator should not change direction")
	}

	it = r.Root().Iterator()
	it.Prev()

	if _, _, ok := it.Next(); ok {
		t.Errorf("Iterator should not change direction")
	}
}
