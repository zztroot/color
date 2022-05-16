# color
Output colored command line. 输出彩色命令行

# 安装
```cmd
go get github.com/zztroot/color
```

# example
```go
s := Coat(fmt.Sprintf("%s%d", "生命诚可贵", 1995), ReverseVideo)
fmt.Println(s)
s = CoatMany("大家好", Red, Underline)
fmt.Println(s)
s = CoatFormat("<c?red?>来得及发垃圾了12345的<c?yellow?>555 dd <c?green?> <c?Underline?> 这 是绿色的")
fmt.Println(s)
```
# output
