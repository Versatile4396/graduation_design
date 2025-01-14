package Log

import "fmt"

// \033 这是标记变换颜色的起始标记，这之后的[1;31;40m则是定义颜色，1表示代码的意义或者说是显示方式，31表示前景颜色，40则是背景颜色。
// 在这定义之后终端就会显示你设定的样式，如果只是要改变一行的样式则在结尾加入\033[0m表示恢复终端默认样式。

func Info(msg string) {
	fmt.Printf("\033[1;32;42m%s\033[0m\n\n", msg)
}
func Warning(msg string) {
	fmt.Printf("\033[1;33;40m%s\033[0m\n\n", msg)
}

func Error(msg string) {
	fmt.Printf("\033[1;31;41m%s\033[0m\n\n", msg)
}

func Fmt(info interface{}) {
	fmt.Printf("\033[1;36;40m%+v\033[0m\n\n", info)
}
