package main
import(
	"fmt"
)
type Node struct{
	data int
	next *Node
}
type List struct{
	//head and tail point to the memory address of the data in the list, hence defining them as pointers.
	head *Node
	len int
}
func (l List)checklength(){
	if l.len<0{
		fmt.Println("Non-existent list")
	}
}

func (l *List)appendList(value int){
	newNode:=&Node{data: value}
	if l.len==0{//inserting at the beginning
		l.head=newNode
		l.head.next=nil
		l.len++
	}else {
		//inserting after the first
		current:=l.head
		current.next=newNode
		l.len++
	}
}

func (l List) printList(){
	for l.len!=0{
		fmt.Print(l.head.data)
		l.head=l.head.next
		l.len--
		fmt.Print("-> ")
	}
	fmt.Print(nil)
	fmt.Println("")
}
func main(){
	fmt.Println("Linked Lists...")
	newList:= List{}
	newList.appendList(10)
	newList.printList()
	newList.appendList(20)
	newList.printList()
}



/* type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
	tail *Node
}

func newList() *List {
	return &List{}
}

func (l *List) append(data int) {
	newNode := &Node{data: data}
	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
	if l.head == nil {
		l.head = newNode
	}
}

func (l *List) display() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	fmt.Println("Linked Lists...")

	list := newList()
	list.append(10)
	list.append(20)
	list.append(30)

	list.display()
} */