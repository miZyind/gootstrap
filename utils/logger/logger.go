package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func print(namespace string, context ...interface{}) {
	timestamp := green(time.Now().Format("2006-01-02 15:04:05"))
	namespace = yellow(namespace)
	prefix := fmt.Sprintf("%s [%s]", timestamp, namespace)
	context = append([]interface{}{prefix}, context...)

	fmt.Println(context...)
}

func Boot(message string, context ...interface{}) {
	if context == nil {
		context = []interface{}{green(message)}
	} else {
		context = []interface{}{
			green(fmt.Sprintf("%s:", message)),
			cyan(context...),
		}
	}

	print("Boot", context...)
}

func InitRouter(name, path string) {
	prefix := green(fmt.Sprintf("%s bound {", name))
	middle := cyan(path)
	suffix := green("}:")

	print("Router", prefix+middle+suffix)
}
