package singleLinkList

import "fmt"

type ElemType interface{}

type Node struct {
	Data ElemType
	Next *Node
}

type List struct {
	headNode *Node
	length   int
}

// 创建结点
func NewListNode(data ElemType) *Node {
	return &Node{data, nil}
}

// 创建链表
func NewLinkList() *List {
	return &List{nil, 0}
}

// 获取下一个结点
func (this *Node) NextNode() *Node {
	return this.Next
}

// 获取结点值
func (this *Node) GetValue() ElemType {
	return this.Data
}

// 长度
func (this *List) Length() int {
	return this.length
}

// 判空
func (this *List) IsEmpty() bool {
	if this.headNode == nil {
		return true
	}

	return false
}

// 从头部添加元素
func (this *List) Add(data ElemType) {
	node := NewListNode(data)

	node.Next = this.headNode
	this.headNode = node

	this.length++
}

// 从尾部添加元素
func (this *List) Append(data ElemType) {
	node := NewListNode(data)

	if this.IsEmpty() {
		this.headNode = node
	} else {
		tmp := this.headNode
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = node
	}

	this.length++
}

// 插入任意位置
func (this *List) Insert(index int, data ElemType) {
	if index <= 0 {
		this.Add(data)
	} else if index > this.Length() {
		this.Append(data)
	} else {
		pre := this.headNode
		for cnt := 0; cnt < index-1; cnt++ {
			pre = pre.Next
		}

		node := NewListNode(data)
		node.Next = pre.Next
		pre.Next = node

		this.length++
	}
}

// 删除指定元素
func (this *List) Remove(data ElemType) {
	pre := this.headNode
	if pre.Data == data {
		this.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
				this.length--
			} else {
				pre = pre.Next
			}
		}
	}
}

// 删除指定位置元素
func (this *List) RemoveIndex(index int) {
	pre := this.headNode

	if index <= 0 {
		this.headNode = pre.Next
	} else if index >= this.Length() {
		// 无操作
		return
	} else {
		for cnt := 0; cnt < index-1 && pre.Next != nil; cnt++ {
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}

	this.length--
}

// 判断包含元素
func (this *List) FindByValue(data ElemType) (node *Node, ok bool) {
	tmp := this.headNode
	for tmp != nil {
		if tmp.Data == data {
			return tmp, true
		}
		tmp = tmp.Next
	}

	return nil, false
}

// 查找指定位置元素
func (this *List) FindByIndex(index int) (node *Node, ok bool) {
	if index <= 0 || index > this.Length() {
		return nil, false
	} else {
		tmp := this.headNode
		for i := 0; i < index; i++ {
			tmp = tmp.Next
		}
		return tmp, true
	}
}

// 格式化链表元素为字符串
func (this *List) FormatList() string {
	lstring := ""

	for tmp := this.headNode; tmp != nil; tmp = tmp.Next {
		lstring += fmt.Sprintf("[%+v]", tmp.GetValue())
		if tmp.Next != nil {
			lstring += "->"
		}
	}

	return lstring
}

// 判断链表有环
func (this *List) HasCycle() bool {
	if this.headNode != nil {
		fast := this.headNode
		slow := this.headNode
		for fast != nil && slow != nil {
			slow = slow.Next
			fast = fast.Next.Next
			if fast == slow {
				return true
			}
		}
	}

	return false
}

// 链表反转
func (this *List) Reverse() {
	if this.headNode == nil || this.headNode.Next == nil {
		return
	}

	oldHeadNode := this.headNode
	cur := oldHeadNode.Next

	for cur != nil {
		this.Add(cur.Data)
		oldHeadNode.Next = cur.Next
		this.length--
		cur = oldHeadNode.Next
	}

	return
}
