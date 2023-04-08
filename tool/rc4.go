package tool

import "crypto/rc4"

func Rc4(input, key []byte) (output []byte) {
	output = make([]byte, len(input))
	c, _ := rc4.NewCipher(key)
	c.XORKeyStream(output, input)
	return output
}
