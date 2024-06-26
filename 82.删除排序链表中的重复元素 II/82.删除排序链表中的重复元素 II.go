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

func deleteDuplicateElement(n *NodeList) *NodeList {
	dummy := &NodeList{}
	dummy.Next = n
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			value := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == value {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func main() {
	node1 := createNodeList([]int{1, 2, 2, 3, 3, 4, 5, 5, 7})
	result := deleteDuplicateElement(node1)
	printNodeList(result)
}
