package trie

type node struct {
	val      []byte
	children Trie
}

func newNode(val []byte) node {
	return node{
		val:      val,
		children: New(),
	}
}

// Trie implements a Trie using maps
// to store the child nodes at each level of the trie.
type Trie map[byte]node

func New() Trie {
	return make(map[byte]node)
}

// Get returns the value in the trie with the specified key, if it exists.
func (t *Trie) Get(key string) ([]byte, bool) {
	var (
		root    Trie = *t
		elem    node
		present bool
	)

	for i := 0; i < len(key); i++ {
		elem, present = root[key[i]]
		if !present {
			return nil, false
		}

		root = elem.children
	}

	return elem.val, true
}

// Set assigns the specified value to the node in the trie at the provided key's location.
// Calling Set with an empty key is not permitted.
func (t *Trie) Set(key string, val []byte) {
	if len(key) == 0 {
		panic("inserting an empty key is not allowed!")
	}

	var (
		root    Trie = *t
		elem    node
		present bool
	)

	for i := 0; i < len(key); i++ {
		if elem.children != nil {
			root = elem.children
		}

		k := key[i]
		elem, present = root[k]
		if !present {
			elem = newNode(nil)
			root[k] = elem
		}
	}

	elem.val = val
	root[key[len(key)-1]] = elem
}

// Del deletes the value identified by the provided key from the trie,
// and returns it to the caller.
func (t *Trie) Del(key string) []byte {
	var (
		root    Trie = *t
		elem    node
		present bool
	)

	for i := 0; i < len(key); i++ {
		if elem.children != nil {
			root = elem.children
		}

		elem, present = root[key[i]]
		if !present {
			return nil
		}
	}

	val := elem.val
	if len(elem.children) == 0 {
		// If this is a leaf-node, then remove it completely.
		delete(root, key[len(key)-1])
	} else {
		// Not a leaf-node, just delete the value from the trie.
		elem.val = nil
	}

	return val
}
