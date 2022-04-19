package main

import (
	"fmt"
)

// 链表的操作方法
type ListMethod interface {
	InsertData(x interface{}) interface{} // 从链表尾部插入值
	SearchData(x interface{}) interface{} // 查找指定元素
	DeleteData(x interface{}) interface{} // 删除指定元素
	PrintList()                           // 打印链表
	ClearList()                           // 清除链表
}

// 链表结构
type List struct {
	val  interface{}
	next *List
}

/*
* @brief: 插入新元素到链表尾部
* @params: 新增元素
* @return:	链表头指针
 */
func (l *List) InsertData(x interface{}) interface{} {
	if l == nil {
		l = new(List)
		l.next = nil
		l.val = x
	} else {
		p := l
		for ; p.next != nil; p = p.next {
		}

		tmp := &List{
			val:  x,
			next: nil,
		}

		p.next = tmp
	}

	return l
}

/*
* @brief: 在链表中查找指定元素
* @params: 查找目标元素
* @return:	找到返回元素指针，未找到返回nil
 */
func (l *List) SearchData(x interface{}) interface{} {
	// if l == nil {
	// 	return nil
	// }

	for p := l; p != nil; p = p.next {
		if p.val == x {
			return p
		}
	}

	return nil
}

/*
* @brief: 从链表中删除目标元素
* @params: 需要被删除的元素
* @return:	返回删除元素后的链表头指针
 */
func (l *List) DeleteData(x interface{}) interface{} {
	if l == nil {
		return nil
	}

	if l.val == x {
		return l.next
	}

	for p := l; p.next != nil; p = p.next {
		if p.next.val == x {
			p.next = p.next.next
			break
		}
	}
	// l.PrintList() // 可以在这里调用链表的其余方法
	return l
}

/*
* @brief: 打印链表
 */
func (l *List) PrintList() {
	if l == nil {
		return
	}

	fmt.Print(l.val)
	l = l.next

	for ; l != nil; l = l.next {
		fmt.Print("->", l.val)
	}
	fmt.Println()
}

func main() {
	var head *List
	fmt.Println(head == nil) //true

	for i, j := 1, 2; i < 12; i, j = i+1, j+1 {
		head = head.InsertData(float32(i) / float32(j)).(*List)
		head = head.InsertData(fmt.Sprintf("%c", i+64)).(*List)
	}

	head.PrintList() // 0.5->A->0.6666667->B->0.75->C->0.8->D->0.8333333->E->0.85714287->F->0.875->G->0.8888889->H->0.9->I->0.90909094->J->0.9166667->K

	//head = head.DeleteData('E').(*List) //这里穿进去，'E'变成了int32类型，因此无法删除
	head = head.DeleteData("E").(*List)

	// tar := head.SearchData(0.9) // 默认float64，和float32值相同，但 == 为false
	tar := head.SearchData(float32(0.9))
	fmt.Println(tar) // <nil>
	head.PrintList() // 0.5->A->0.6666667->B->0.75->C->0.8->D->0.8333333->0.85714287->F->0.875->G->0.8888889->H->0.9->I->0.90909094->J->0.9166667->K

	fmt.Println(head) //&{0.5 0xc000004090}
}
