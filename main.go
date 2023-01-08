package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Printf("init")
}

// printf("%c",a)；输出单个字符。
// printf("%d",a)；输出十进制整数。
// printf("%f",a)；输出十进制浮点数.
// printf("%o",a)；输出八进制数。
// printf("%s",a)；输出字符串。
// printf("%u",a)；输出无符号十进制数。
// printf("%x",a)；输出十六进制数。
func main() {
	log.Printf("main")

	//make
	var b map[string]int    //map(key,value)

	b = make(map[string]int, 10)  {new ()}
	fmt.Printf("打印b2=%v", b)

	b["测试"] = 100
	b["测试1"] = 1001
	fmt.Printf("打印b3=%v", b)

	//指针
	//a := 10
	//b := &a
	//fmt.Printf("a:%d ptr:%p\n", a, &a) //   十六进制表示，前缀 0x
	//fmt.Printf("b:%p type:%T\n", b, b) // type:*int
	//fmt.Println(&b)

	//切片
	////1.声明切片
	//var s1 []int
	//if s1 == nil {
	//	fmt.Println("s1是空")
	//} else {
	//	fmt.Println("s1不是空")
	//}
	//// 2.:=
	//s2 := []int{}
	//// 3.make()
	//var s3 []int = make([]int, 0)
	//fmt.Println(s1, s2, s3)
	//// 4.初始化赋值
	//var s4 []int = make([]int, 0, 0)
	//fmt.Printf("s4=%v", s4)
	//s5 := []int{1, 2, 3}
	//fmt.Printf("s5=%v", s5)
	//// 5.从数组切片
	//arr := [5]int{1, 2, 3, 4, 5}
	//var s6 []int
	//// 前包后不包
	//s6 = arr[1:4]
	//fmt.Printf("取索引2-4的值=%v", s6)

}
