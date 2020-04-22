package models

const (
	ALPHABET_SIZE = 26
)

// type Trie struct {
// 	isWord   bool
// 	children map[int32]*Trie
// }

type trieNode struct {
	isWord   bool
	children [ALPHABET_SIZE]*trieNode
}

type Trie struct {
	root *trieNode
}

func InitTrie() *Trie {
	return &Trie{
		root: &trieNode{},
	}
}

func (t *Trie) Insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWord = true
}

func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	if current.isWord {
		return true
	}
	return false
}
