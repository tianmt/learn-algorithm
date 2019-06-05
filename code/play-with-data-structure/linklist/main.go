package main

import (
	"fmt"

	mySingleLinkList "linkList/singleLinkList"
)

func main() {
	list := mySingleLinkList.NewLinkList()
	list.Add("Hello")
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(4)
	list.Append(4)
	fmt.Println(list.FindByIndex(1))
	fmt.Println(list.FindByIndex(10))
	fmt.Println(list.FindByValue("Hello"))
	fmt.Println(list.Length())
	list.Remove(4)
	fmt.Println(list.Length())
	fmt.Println(list.FormatList())
	fmt.Println(list.HasCycle())

	// reverse format
	list.Reverse()
	fmt.Println(list.FormatList())
	fmt.Println(list.Length())
}
