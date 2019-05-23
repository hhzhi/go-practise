package main

import (
	"fmt"
)

// PersonInfo {}
type PersonInfo struct {
	id   string
	name string
}

func main() {

	// 创建map集合
	var map1 map[string]PersonInfo

	// 初始化集合
	map1 = make(map[string]PersonInfo)

	map1["1"] = PersonInfo{"a", "b"}
	map1["2"] = PersonInfo{"c", "d"}
	map1["3"] = PersonInfo{"e", "f"}

	fmt.Println("map1 length = ", len(map1))

	// 删除集合
	delete(map1, "3")

	fmt.Println("map1 length = ", len(map1))

	// 获取map集合
	person, ok := map1["1"]

	if ok {
		fmt.Println("1 name = ", person.name)
	} else {
		fmt.Println("not find person 1 name")
	}

}
