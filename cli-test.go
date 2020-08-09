package main

import "webhook-runner/interpreter"

func main() {
	interp, err := interpreter.New("./usr")
	if err != nil {
		panic(err)
	}

	if err := interp.LoadFunction("hugo.Bar"); err != nil {
		panic(err)
	}

	if err := interp.RunLoadedHook("hugo.Bar", "Kung"); err != nil {
		panic(err)
	}
}
