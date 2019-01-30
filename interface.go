package main

import "fmt"

type notifer interface {
	notify()
}

type user struct {
	name string
	email string
}

type admin struct {
	name string
	age  int
}

/*
	使用user指针接收者实现了notofy接口
 */
func (u *user) notify()  {
	fmt.Println("name:", u.name)
	u.name = u.name + " user"
}

/*
	使用admin指针接收者实现了notofy接口
 */
func (u *admin) notify(){
	fmt.Println("name:", u.name)
	u.name = u.name + " admin"
}

/*
	接收一个实现了notifer接口值

	多态是指可以根据类型的具体实现采取不同行为的能力。
	如果一个类型实现了某个接口，所有使用这个接口的地方，都可以支持这种类型
 */
func sendNotify(n notifer){
	n.notify()

}

func main()  {

	weihua := user{"张三", "qq@qq.com"}
	sendNotify(&weihua)

	huang := admin{"李四", 25}
	sendNotify(&huang)

	a := admin{"www",1}

	var i notifer
	i = &a
	i.notify()

}

