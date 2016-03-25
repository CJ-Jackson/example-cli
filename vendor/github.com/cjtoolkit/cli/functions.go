package cli

import (
	"fmt"
	"io"
)

func splitKeyValue(keyValue string) (key, value string) {
	value = "1"
	pos := -1
	for key, char := range keyValue {
		if char == '=' {
			pos = key
			break
		}
	}
	if pos == -1 {
		key = keyValue
		return
	}
	key = keyValue[:pos]
	value = keyValue[pos+1:]
	return
}

func execTrueFalse(assert bool, trueFn, falseFn func()) {
	if assert {
		trueFn()
	} else {
		falseFn()
	}
}

func handleErrorAndPanicAgain(prefix string) {
	switch r := recover().(type) {
	case string:
		panic(prefix + r)
	case error:
		panic(prefix + r.Error())
	}
}

func handleErrorAndExit(prefix string, exitFn func(), w io.Writer) {
	switch r := recover().(type) {
	case string:
		w.Write([]byte(prefix))
		w.Write([]byte(r))
		exitFn()
	case error:
		w.Write([]byte(prefix))
		w.Write([]byte(r.Error()))
		exitFn()
	}
}

func execFunctionIfNil(value interface{}, fn func()) {
	if "<nil>" == fmt.Sprint(value) {
		fn()
	}
}

func execFunctionIfNotNil(value interface{}, fn func()) {
	if "<nil>" != fmt.Sprint(value) {
		fn()
	}
}

func Run() {
	main_cli.run()
}

func RegisterGlobal(global GlobalInterface) {
	main_cli.registerGlobal(global)
}

func RegisterCommand(command CommandInterface) {
	main_cli.registerCommand(command)
}
