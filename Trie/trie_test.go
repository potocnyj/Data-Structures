package trie

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := New()

	in := []byte("bar")
	key := "foo"
	tree.Set(key, in)

	out, ok := tree.Get(key)
	if !ok || !bytes.Equal(in, out) {
		t.Errorf("Did not find val %s at key %s in tree, got %s instead (present: %t)", in, key, out, ok)
	}

	key2 := "baz"
	in = []byte{1}
	tree.Set(key2, in)

	tree.Del(key)

	out, ok = tree.Get(key)
	if ok {
		t.Errorf("found deleted key %s in tree: %s", key, out)
	}

	out, ok = tree.Get(key2)
	if !ok || !bytes.Equal(in, out) {
		t.Errorf("Did not find val %s at key %s in tree, got %s instead (present: %t)", in, key2, out, ok)
	}
}

func getTestData(t testing.TB, elements, size int) []string {
	res := make([]string, 0, elements)

	buf := make([]byte, size)
	for i := 0; i < elements; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		res = append(res, string(buf))
	}

	return res
}

var cases = []int{
	4,
	8,
	16,
	32,
	64,
	256,
	512,
}

func BenchmarkHashSet(b *testing.B) {
	for _, size := range cases {
		b.Run(fmt.Sprintf("keySize%d", size), func(b *testing.B) {
			data := getTestData(b, 5000, size)
			var tree Trie

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if i%len(data) == 0 {
					tree = New()
				}

				tree.Set(data[i%len(data)], []byte{byte(i)})
			}
		})
	}
}

func BenchmarkHashGet(b *testing.B) {
	for _, size := range cases {
		b.Run(fmt.Sprintf("keySize%d", size), func(b *testing.B) {
			data := getTestData(b, 5000, size)
			tree := New()
			for j, key := range data {
				tree.Set(key, []byte{byte(j)})
			}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, ok := tree.Get(data[i%len(data)])
				if !ok {
					b.Fail()
				}
			}
		})
	}
}
