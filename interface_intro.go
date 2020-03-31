package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

// 定义了String方法，实现了fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

// 面向对象编程,需要将变量和方法绑定在一起.在go里面通过struct组织变量,再通过interface表示接口,将方法与变量绑定在一起.
// 扣税计算
type PayTaxForUser struct {
	// 定义用户税收相关结构
	taxFree int64
	taxAble int64
	total   int64
	userId  int64
}

// 用于初始化对象, 一般是赋予必要变量的初始值,类似于构造函数
// 如果有db连接等,还可以增加一个init函数用于初始化
func NewPayTaxService(userId int64, taxFree int64) *PayTaxForUser {
	if userId == 0 {
		return nil
	}

	return &PayTaxForUser{
		userId:  userId,
		taxFree: taxFree,
	}
}

func (self *PayTaxForUser) CalcTax() int64 {
	var tax int64 = 0
	if self.taxAble <= 800*100 {
		tax = 0

	} else if self.taxAble <= 4000*100 {
		tax = (self.taxAble - 800*100) * 20 / 100

	} else {
		tax = self.taxAble*80*40/10000 - 7000*100
	}

	return tax
}

func (self *PayTaxForUser) minus(anotherPay PayTaxForUser) PayTaxForUser {

	resPay := PayTaxForUser{
		taxAble: self.taxAble - anotherPay.taxAble,
		taxFree: self.taxFree - anotherPay.taxFree,
		total:   self.total - anotherPay.total,
	}

	return resPay
}

func (self *PayTaxForUser) add(anotherPay PayTaxForUser) PayTaxForUser {

	resPay := PayTaxForUser{
		taxAble: self.taxAble + anotherPay.taxAble,
		taxFree: self.taxFree + anotherPay.taxFree,
		total:   self.total + anotherPay.total,
	}

	return resPay
}

func main() {
	// 没做初始化则默认都是0值

	// 使用方式1 不初始化使用默认值,一般不这么用,没有意义
	a := PayTaxForUser{}
	fmt.Println(a)
	fmt.Println(a.CalcTax())

	a.taxFree = 100
	a.taxAble = 100000
	fmt.Println(a.CalcTax())

	// 使用方式2, 通过专门的new接口初始化
	b := NewPayTaxService(123, 100)
	fmt.Println(b.CalcTax())
	b.taxAble = 200000
	fmt.Println(b.CalcTax())

	// 使用方式3, 直接初始化
	c := PayTaxForUser{
		userId:  234,
		taxFree: 100,
		taxAble: 3000000,
	}
	fmt.Println(c.CalcTax())
}
