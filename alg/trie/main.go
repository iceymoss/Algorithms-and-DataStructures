package main

import "fmt"

type Node struct {
	children map[rune]*Node
	char     string
	isEnding bool
}

func NewNode(char string) *Node {
	return &Node{
		children: make(map[rune]*Node),
		char:     char,
		isEnding: false,
	}
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: NewNode("/"),
	}
}

// Insert 插入单词
func (trie *Trie) Insert(str string) {
	curNode := trie.root
	for _, ch := range str {
		//查看当前是否存在对应字符的k-v对
		value, ok := curNode.children[ch] //读当前节点的子节点
		if !ok {
			value = NewNode(string(ch))
			curNode.children[ch] = value
		}
		//更新当前节点
		curNode = value
	}
	//个单词遍历完所有字符后将结尾字符打上标记
	curNode.isEnding = true
}

func (trie *Trie) Find(str string) bool {
	curNode := trie.root
	for _, ch := range str {
		value, ok := curNode.children[ch]
		if !ok {
			return false
		}
		curNode = value
	}
	if curNode.isEnding == false {
		return false
	}
	return true
}

func (trie *Trie) StartsWith(str string) bool {
	curNode := trie.root
	for _, ch := range str {
		value, ok := curNode.children[ch]
		if !ok {
			return false
		}
		curNode = value
	}
	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("iceymoss")
	trie.Insert("apple")
	fmt.Println(trie.Find("iceymos"))
	fmt.Println(trie.Find("apple"))
	fmt.Println(trie.StartsWith("app"))
}
