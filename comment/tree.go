/*
 * @Description:
 * @Author: gphper
 * @Date: 2021-09-23 21:07:01
 */
package comment

import (
	"sort"
)

// 字典树节点
type TrieNode struct {
	Children map[string]*TrieNode
	Types    string
	IsEnd    bool
}

// 构造字典树节点
func newTrieNode() *TrieNode {
	return &TrieNode{Children: make(map[string]*TrieNode), IsEnd: false, Types: "none"}
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
func (trie *Trie) Insert(word []string, types string) {
	node := trie.Root
	for i := 0; i < len(word); i++ {
		_, ok := node.Children[word[i]]
		if !ok {
			node.Children[word[i]] = newTrieNode()
		}
		node = node.Children[word[i]]
	}

	node.IsEnd = true
	node.Types = types
}

type Node struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	All      string `json:"all"`
	Children []Node `json:"children"`
}

func GetOne(node *TrieNode, allpre string) []Node {
	slice := make([]Node, 0)

	for k, _ := range node.Children {

		var tmp []Node
		if node.Children[k].IsEnd {

			if node.Children[k].Types != "none" {
				var qian string
				if allpre != "" {
					qian = allpre + ":" + k
				} else {
					qian = k
				}
				tmp = []Node{Node{Title: k, Type: node.Children[k].Types, All: qian}}
			}

			tmp = append(tmp, GetRepeat(node.Children[k], k, allpre)...)
		} else {
			var qian string
			if allpre != "" {
				qian = allpre + ":" + k
			} else {
				qian = k
			}
			tmp = []Node{Node{
				Title:    k,
				All:      qian,
				Children: GetOne(node.Children[k], qian),
			}}
		}

		slice = append(slice, tmp...)

	}

	sort.Slice(slice, func(i, j int) bool {

		if slice[i].Title == slice[j].Title {
			return len(slice[i].Children) > len(slice[j].Children)
		} else {
			return slice[i].Title < slice[j].Title
		}

	})

	return slice
}

func GetRepeat(node *TrieNode, pre string, allpre string) []Node {

	slice := make([]Node, 0)
	for k, v := range node.Children {

		var tempNode Node

		if node.Children[k].Types != "none" {
			pre += ":" + k
			tempNode = Node{Title: pre, Type: node.Children[k].Types, All: allpre + ":" + k}
		} else {
			pre += ":" + k
			tempNode = Node{Title: pre}
		}

		slice = append(slice, tempNode)

		tmpSlice := GetRepeat(v, pre, allpre)

		if len(tmpSlice) > 0 {
			slice = append(slice, tmpSlice...)
		}
	}

	return slice
}
