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
	namespace = yellow("[" + namespace + "]")
	context = append([]interface{}{timestamp + " " + namespace}, context...)

	fmt.Println(context...)
}

func Boot(message string, context ...interface{}) {
	if context == nil {
		context = []interface{}{green(message)}
	} else {
		context = []interface{}{
			green(message + ":"),
			cyan(context...),
		}
	}

	print("Bootstrapper", context...)
}

func InitRouter(name, path string) {
	prefix := green(name + "Router initialized {")
	middle := cyan(path)
	suffix := green("}:")

	print("Router", prefix+middle+suffix)
}

func BindRoute(method, route string) {
	switch method {
	case "GET":
		method = color.HiCyanString(method)
	case "POST":
		method = color.HiGreenString(method)
	case "PUT":
		method = color.HiYellowString(method)
	case "PATCH":
		method = color.HiBlueString(method)
	case "DELETE":
		method = color.HiRedString(method)
	default:
		method = color.HiMagentaString(method)
	}

	route = fmt.Sprintf("%s - %s", cyan(route), method)

	print("Route", green("Bound {")+route+green("}"))
}
