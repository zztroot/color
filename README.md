# color
Output colored command line. 输出彩色命令行

# install
```cmd
go get github.com/zztroot/color
```
或者使用go mod tidy

# example
```go
package main

import (
	"fmt"
	"github.com/zztroot/color"
)

func main() {
	fmt.Println(color.Coat("我是绿色", color.Green))
	
	fmt.Println(color.CoatMany("我是红色加下划线", color.Red, color.Underline))
	
	// 在需要输入的字符串前面加上标识符，如：<c?red?> 红色
	fmt.Println(color.CoatFormat("<c?red?>我是红色<c?yellow?>我是黄色 <c?green?> <c?Underline?> 我 是绿色加下划线"))
	
	name := "zhongtian"
	fmt.Println(color.Coat(fmt.Sprintf("%s：我是蓝色的", name), color.Blue))
}
```
# output
![截图_选择区域_20220516104131](https://user-images.githubusercontent.com/46484337/168510917-8b79bf7b-b6e9-452c-bf8e-180e992b16aa.png)
