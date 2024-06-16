package main

//TODO
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
	cur := n
	for n != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return n
}

func main() {
	node1 := createNodeList([]int{1, 2, 2, 3, 4, 4, 5, 6, 7, 7})
	result := deleteDuplicateElement(node1)
	printNodeList(result)
}
