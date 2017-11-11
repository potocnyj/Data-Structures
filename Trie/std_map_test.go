package trie

// Add benchmarks for Set/Get using a builtin map[string][]byte in place of a Trie.

import (
	"fmt"
	"testing"
)

func BenchmarkStdMapSet(b *testing.B) {
	for _, size := range cases {
		b.Run(fmt.Sprintf("keySize%d", size), func(b *testing.B) {
			data := getTestData(b, 50000, size)
			var tree map[string][]byte

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if i%len(data) == 0 {
					tree = make(map[string][]byte)
				}

				tree[data[i%len(data)]] = []byte{byte(i)}
			}
		})
	}
}

func BenchmarkStdMapGet(b *testing.B) {
	for _, size := range cases {
		b.Run(fmt.Sprintf("keySize%d", size), func(b *testing.B) {
			data := getTestData(b, 50000, size)
			tree := make(map[string][]byte)
			for j, key := range data {
				tree[key] = []byte{byte(j)}
			}

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, ok := tree[data[i%len(data)]]
				if !ok {
					b.Fail()
				}
			}
		})
	}
}
