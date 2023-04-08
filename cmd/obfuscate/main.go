package main

import (
	"LazyConvert/tool"
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadFile(os.Args[1])

	key := []byte("ZnVja3lvdQo=")
	out := tool.Rc4(data, key)

	err := ioutil.WriteFile("hello.rc4", out, 0644)
	if err != nil {
		panic(err)
	}
}
