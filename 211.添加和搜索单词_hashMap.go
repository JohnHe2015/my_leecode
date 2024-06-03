/*
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。

实现词典类 WordDictionary ：

WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。

示例：

输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // 返回 False
wordDictionary.search("bad"); // 返回 True
wordDictionary.search(".ad"); // 返回 True
wordDictionary.search("b.."); // 返回 True

提示：

1 <= word.length <= 25
addWord 中的 word 由小写英文字母组成
search 中的 word 由 '.' 或小写英文字母组成
最多调用 104 次 addWord 和 search
*/

// hashMap
package main

import "fmt"

type WordDictionary struct {
	//define key是长度、value是字符串数组
	words map[int][]string
}

func Constructor() WordDictionary {
	return WordDictionary{words: make(map[int][]string)}
}

func (dict *WordDictionary) addWord(word string) {
	wordLen := len(word)
	dict.words[wordLen] = append(dict.words[wordLen], word)
}

func (dict *WordDictionary) searchWord(word string) bool {
	wordLen := len(word)
	if words, ok := dict.words[wordLen]; ok {
		for _, w := range words {
			if match(w, word) {
				return true
			}
		}
	}
	return false
}

func match(word, pattern string) bool {
	if len(word) != len(pattern) {
		return false
	}
	for i := 0; i < len(word); i++ {
		if pattern[i] != '.' && pattern[i] != word[i] {
			return false
		}
	}
	return true
}

func main() {
	wordDict := Constructor()
	wordDict.addWord("hello")
	wordDict.addWord("nice")
	fmt.Println(wordDict.searchWord("hello"))
	fmt.Println(wordDict.searchWord("nice."))
}
