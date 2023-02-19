package trie

type TrieNode struct {
	isEndOfWord bool
	children    [26]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for i := range word {
		char := word[i] - 'a'
		if cur.children[char] == nil {
			cur.children[char] = &TrieNode{}
		}
		cur = cur.children[char]
	}
	cur.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for i := range word {
		char := word[i] - 'a'
		if cur.children[char] == nil {
			return false
		}
		cur = cur.children[char]
	}

	return cur.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for i := range prefix {
		char := prefix[i] - 'a'
		if cur.children[char] == nil {
			return false
		}
		cur = cur.children[char]
	}

	return true
}
