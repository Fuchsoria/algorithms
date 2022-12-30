package hashtable

type Node struct {
	isEnd bool
	child []*Node
}

func newNode() *Node {
	return &Node{
		child: make([]*Node, 128),
	}
}

func (n *Node) Exists(char int) bool {
	return n.child[char] != nil
}

func (n *Node) Next(char rune) *Node {
	ascii := int(char)

	if !n.Exists(ascii) {
		n.child[ascii] = newNode()
	}

	return n.child[ascii]
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: newNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, r := range word {
		node = node.Next(r)
	}

	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.Go(word)
	if node != nil {
		return node.isEnd
	}

	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.Go(prefix) != nil
}

func (t *Trie) Go(word string) *Node {
	node := t.root
	for _, c := range word {
		ascii := int(c)

		if node.Exists(ascii) {
			node = node.Next(c)
		} else {
			return nil
		}
	}

	return node
}
