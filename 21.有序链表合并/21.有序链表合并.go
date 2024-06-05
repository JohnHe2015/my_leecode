/*
	真题描述：将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有结点组成的。

	示例： 输入：1->2->4, 1->3->4 输出：1->1->2->3->4->4
*/

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

func combineNodeList(n1 *NodeList, n2 *NodeList) *NodeList {
	dummy := &NodeList{}
	tail := dummy
	for n1 != nil && n2 != nil {
		if n1.Val < n2.Val {
			tail.Next = n1
			n1 = n1.Next
		} else {
			tail.Next = n2
			n2 = n2.Next
		}
		tail = tail.Next
	}

	if n1 != nil {
		tail.Next = n1
	} else {
		tail.Next = n2
	}
	return dummy.Next
}

func main() {
	node1 := createNodeList([]int{1, 2, 3})
	node2 := createNodeList([]int{4, 5, 6, 8, 10})
	//printNodeList(node1)
	//printNodeList(node2)
	node3 := combineNodeList(node1, node2)
	printNodeList(node3)
}
