package main

import(
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
)

var state [65535]byte
var index = 0
var l = 0

func main() {
	p, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Did you supply a file?")
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(p); i++ {
		switch p[i] {
		default: fmt.Print()
		case '>': index++
		case '<': if index > 0 { index-- }
		case '+': state[index]++
		case '-': state[index]--

		case '.': fmt.Print(string(state[index]))
		case ',': state[index], _ = reader.ReadByte()

		case '[':
			if state[index] == 0 {
				i++
				for l > 0 || p[i] != ']' {
					if p[i] == '[' {
						l++
					} else if p[i] == ']' {
						l--
					}
					i++
				}
			}

		case ']':
			i--
			for l > 0 || p[i] != '[' {
				if p[i] == ']' {
					l++
				} else if p[i] == '[' {
					l--
				}
				i--
			}
			i--
		}

	}
}
