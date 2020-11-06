package logger

import (
	"fmt"
	"reflect"
	"time"

	"github.com/fatih/color"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

func print(ns string, context ...interface{}) {
	ts := green(time.Now().Format("2006-01-02 15:04:05"))
	ns = yellow("[" + ns + "]")
	context = append([]interface{}{ts + " " + ns}, context...)

	fmt.Println(context...)
}

// Bootstrapper prints formatted and colorized messages with "Bootstrapper" namespace
func Bootstrapper(msg string, context ...interface{}) {
	if context == nil {
		context = []interface{}{green(msg)}
	} else {
		context = []interface{}{
			green(msg + ":"),
			cyan(context...),
		}
	}

	print("Bootstrapper", context...)
}

// Config prints formatted and colorized messages with "Config" namespace
func Config(config interface{}) {
	name := reflect.TypeOf(config).Elem().Name()

	print("Config", green(name+" loaded:"), cyan(config))
}

// Router prints formatted and colorized messages with "Router" namespace
func Router(router interface{}, path string) {
	ns := "Router"
	name := reflect.TypeOf(router).Elem().Name()
	prefix := green(name + ns + " initialized {")
	middle := cyan(path)
	suffix := green("}:")

	print(ns, prefix+middle+suffix)
}

// Route prints formatted and colorized messages with "Route" namespace
func Route(method, route string) {
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

	prefix := green("Bound {")
	middle := fmt.Sprintf("%s - %s", cyan(route), method)
	suffix := green("}")

	print("Route", prefix+middle+suffix)
}
