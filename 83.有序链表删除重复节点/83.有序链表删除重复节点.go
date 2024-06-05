package main

import "fmt"

type NodeList struct {
	Next *NodeList
	Val  int
}

func createNodeList(arr []int) *NodeList {
	if len(arr) == 0 {
		return nil
	}
	head := &NodeList{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &NodeList{Val: val}
		current = current.Next
	}
	return head
}

func printNodeList(n *NodeList) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func main() {

}
