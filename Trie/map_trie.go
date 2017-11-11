package trie

type hashNode struct {
	val      []byte
	children HashTrie
}

func newNode(val []byte) hashNode {
	return hashNode{
		val:      val,
		children: make(map[byte]hashNode),
	}
}

// HashTrie implements a Trie using maps
// to store the child nodes at each level of the trie.
type HashTrie map[byte]hashNode

func NewHash() HashTrie {
	return make(map[byte]hashNode)
}

// Get returns the value in the trie with the specified key, if it exists.
func (t *HashTrie) Get(key string) ([]byte, bool) {
	var (
		root    HashTrie = *t
		node    hashNode
		present bool
	)

	for i := 0; i < len(key); i++ {
		node, present = root[key[i]]
		if !present {
			return nil, false
		}

		root = node.children
	}

	return node.val, true
}

// Set assigns the specified value to the node in the trie at the provided key's location.
// Calling Set with an empty key is not permitted.
func (t *HashTrie) Set(key string, val []byte) {
	if len(key) == 0 {
		panic("inserting an empty key is not allowed!")
	}

	var (
		root    HashTrie = *t
		node    hashNode
		present bool
	)

	for i := 0; i < len(key); i++ {
		if node.children != nil {
			root = node.children
		}

		k := key[i]
		node, present = root[k]
		if !present {
			node = newNode(nil)
			root[k] = node
		}
	}

	node.val = val
	root[key[len(key)-1]] = node
}

// Del deletes the value identified by the provided key from the trie,
// and returns it to the caller.
func (t *HashTrie) Del(key string) []byte {
	var (
		root    HashTrie = *t
		node    hashNode
		present bool
	)

	for i := 0; i < len(key); i++ {
		if node.children != nil {
			root = node.children
		}

		node, present = root[key[i]]
		if !present {
			return nil
		}
	}

	val := node.val
	if len(node.children) == 0 {
		// If this is a leaf-node, then remove it completely.
		delete(root, key[len(key)-1])
	} else {
		// Not a leaf-node, just delete the value from the trie.
		node.val = nil
	}

	return val
}
