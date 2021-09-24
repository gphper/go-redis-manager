/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-23 21:07:01
 */
package comment

import (
	"strings"
)

// 字典树节点
type TrieNode struct {
	children map[string]*TrieNode
	isEnd    bool
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[string]*TrieNode), isEnd: false}
}

// 字典树
type Trie struct {
	Root *TrieNode
}

// 构造字典树
func NewTrie() *Trie {
	return &Trie{Root: newTrieNode()}
}

// 向字典树中插入一个单词
func (trie *Trie) Insert(word []string) {
	node := trie.Root
	for i := 0; i < len(word); i++ {
		_, ok := node.children[word[i]]
		if !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

type Node struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Children []Node `json:"children"`
}

func GetOne(node *TrieNode) []Node {
	slice := make([]Node, 0)

	for k, _ := range node.children {

		var tmp []Node
		if node.children[k].isEnd {

			if strings.Contains(k, "_") {
				kTypeSlice := strings.Split(k, "_")
				tmp = []Node{Node{Title: kTypeSlice[0], Type: kTypeSlice[1]}}
			}

			tmp = append(tmp, GetRepeat(node.children[k], k)...)
		} else {
			tmp = []Node{Node{
				Title:    k,
				Children: GetOne(node.children[k]),
			}}
		}

		slice = append(slice, tmp...)

	}
	return slice
}

func GetRepeat(node *TrieNode, pre string) []Node {

	slice := make([]Node, 0)
	for k, v := range node.children {

		var tempNode Node
		if strings.Contains(k, "_") {
			kTypeSlice := strings.Split(k, "_")
			pre += ":" + kTypeSlice[0]
			tempNode = Node{Title: pre, Type: kTypeSlice[1]}
		} else {
			pre += ":" + k
			tempNode = Node{Title: pre}
		}

		slice = append(slice, tempNode)

		tmpSlice := GetRepeat(v, pre)

		if len(tmpSlice) > 0 {
			slice = append(slice, tmpSlice...)
		}
	}

	return slice
}
