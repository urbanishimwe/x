package ds

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Any means any tipe can be added
type Any interface{}

// ListNode hold all nodes in a linkedlist
type ListNode struct {
	data Any
	next *ListNode
}

// List every linkedlist have head and tail node
type List struct {
	head *ListNode
	tail *ListNode
}

func (List *List) insert(nodeData Any) {
	node := &ListNode{
		next: nil,
		data: nodeData,
	}

	if List.head == nil {
		List.head = node
	} else {
		List.tail.next = node
	}

	List.tail = node
}

// LinkedList return new LinkedList
func LinkedList() List {
	return List{}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist := List{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	printLinkedList(llist.head)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
