package main

import (
	"fmt"
)

const maxSize = 8 // 元素个数

type elemType int64 // 元素类型

// 错误值
var errCodeMsg = map[errCode]string{
	OK:         "正常",
	InsertErr:  "插入错误",
	AppendErr:  "追加错误",
	GetErr:     "获取元素错误",
	DeleteErr:  "删除错误",
	UnknownErr: "位置错误",
}

type errCode int32 // 错误码

const (
	OK         errCode = 101
	InsertErr  errCode = 102
	AppendErr  errCode = 103
	GetErr     errCode = 104
	DeleteErr  errCode = 105
	UnknownErr errCode = 109
)

type SqList struct {
	data   [maxSize]elemType // 数据元素个数固定
	length int               // 记录顺序表中当前数据元素个数
}

func getElem(list *SqList, pos int) (e elemType, ec errCode) {
	if list.length == 0 || pos < 1 || pos > list.length {
		return e, GetErr
	}

	e = list.data[pos-1]

	return e, OK
}

func getListLen(list *SqList) int {
	return list.length
}

func listInsert(list *SqList, pos int, e elemType) errCode {
	if list.length == maxSize {
		return InsertErr
	}

	if pos < 1 || pos > list.length+1 {
		return InsertErr
	}

	if pos <= list.length {
		for i := list.length - 1; i >= pos-1; i-- {
			list.data[i+1] = list.data[i]
		}
	}

	list.data[pos-1] = e
	list.length++

	return OK
}

func listAppend(list *SqList, e elemType) errCode {
	if list.length == maxSize {
		return AppendErr
	}

	list.data[list.length] = e

	list.length++

	return OK
}

func listDelete(list *SqList, pos int) (e elemType, ec errCode) {
	if list.length == 0 {
		return e, DeleteErr
	}

	if pos < 1 || pos > list.length {
		return e, DeleteErr
	}

	e = list.data[pos-1]

	if pos < list.length {
		for i := pos; i < list.length; i++ {
			list.data[i-1] = list.data[i]
		}
	}

	list.length--

	return e, OK
}

func listPrint(list *SqList) {
	fmt.Println("list length:", getListLen(list))
	fmt.Print("list data: ")
	for i := 0; i < list.length; i++ {
		fmt.Print(list.data[i], " ")
	}
	fmt.Println()
}

func handleError(code errCode) {
	fmt.Println(errCodeMsg[code])
}

func main() {
	sl := new(SqList)

	fmt.Println("=====================[1.print empty list]===================")
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[2.print append list]==================")
	err := listAppend(sl, 5)
	if err != OK {
		handleError(err)
	} else {
		fmt.Println("插入成功！")
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[3.print append list]==================")
	err = listAppend(sl, 15)
	if err != OK {
		handleError(err)
	} else {
		fmt.Println("插入成功！")
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[4.print append list]==================")
	err = listAppend(sl, 25)
	if err != OK {
		handleError(err)
	} else {
		fmt.Println("插入成功！")
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[5.print del list]=====================")
	e, ok := listDelete(sl, 1)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("del:", e)
		listPrint(sl)
	}
	fmt.Println("============================================================")

	fmt.Println("=====================[6.print del list]=====================")
	e, ok = listDelete(sl, 3)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("del:", e)
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[7.print insert list]==================")
	ok = listInsert(sl, 1, 3)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("insert:", 3)
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[8.print insert list]==================")
	ok = listInsert(sl, 1, 9)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("insert:", 9)
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[9.print insert list]==================")
	ok = listInsert(sl, 9, 19)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("insert:", 19)
	}
	listPrint(sl)
	fmt.Println("============================================================")

	fmt.Println("=====================[10.print get list]====================")
	e, ok = getElem(sl, 3)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("get:", e)
	}
	fmt.Println("============================================================")

	fmt.Println("=====================[11.print get list]====================")
	e, ok = getElem(sl, 13)
	if OK != ok {
		handleError(ok)
	} else {
		fmt.Println("get:", e)
	}
	fmt.Println("============================================================")

	fmt.Println("=====================[12.print get list len]================")
	fmt.Println("get list len:", getListLen(sl))
	fmt.Println("============================================================")
}
