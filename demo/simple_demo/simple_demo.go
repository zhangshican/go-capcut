package main

import (
	"fmt"
	"github.com/zhangshican/go-capcut/internal/types"
)

func main() {
	fmt.Println("测试时间工具")

	// 简单测试
	result, err := types.Tim("5s")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("5s = %d 微秒\n", result)
	}

	// 测试时间范围
	tr := types.NewTimerange(1000000, 2000000)
	fmt.Printf("时间范围: %s\n", tr)
}
