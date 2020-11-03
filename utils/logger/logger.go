package logger

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen)
	cyan  = color.New(color.FgCyan)
)

func print(context ...interface{}) {
	fmt.Println(
		append(
			[]interface{}{green.Sprint(time.Now().Format("2006-01-02 15:04:05"))},
			context...,
		)...,
	)
}

func printNs(namespace string, context ...interface{}) {
	print(
		append(
			[]interface{}{color.YellowString("[%s]", namespace)},
			context...,
		)...,
	)
}

// Boot will return colorized message with Bootstrapper namespace
func Boot(message string, context ...interface{}) {
	namespace := "Bootstrapper"
	if context == nil {
		printNs(namespace, green.Sprint(message))
	} else {
		printNs(namespace, green.Sprintf("%s:", message), cyan.Sprint(context...))
	}
}

// Log ...
func Log(context ...interface{}) {
	print(context...)
}

// LogNs ...
func LogNs(namespace string, context ...interface{}) {
	printNs(namespace, context...)
}
