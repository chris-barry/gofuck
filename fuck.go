package main

import(
	"os"
	"fmt"
	"bufio"
	"io/ioutil"
)

var state [65535]byte
var index = 0

var whileIndex = 0

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

		// TODO: I don't think this works.
		case '[':
			whileIndex, whileState = index, i
			fmt.Println("HEHE")

		case ']':
			fmt.Println(state[whileIndex])
			if state[whileIndex] != 0 {
				index = whileState
			}
		}
	}
}
