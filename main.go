package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/Soreil/mnemonic"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Provide an output filename")
		os.Exit(1)
	}
	file, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	mnemonic.SetSalt("1test2test3test4test5test6test7test8test")
	for a := byte(0); a < 0xff; a++ {
		for b := byte(0); b < 0xff; b++ {
			for c := byte(0); c < 0xff; c++ {
				for d := byte(0); d < 0xff; d++ {
					ip := net.IPv4(a, b, c, d).String()
					m, err := mnemonic.Mnemonic(ip)
					if err != nil {
						panic(err)
					}
					writer.WriteString(m + "," + ip + "\n")
				}
			}
		}
	}
}
