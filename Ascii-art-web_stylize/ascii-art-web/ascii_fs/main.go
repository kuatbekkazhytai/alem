package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Convert(s []string) string {
	str := ""
	for _, w := range s {
		str = str + w
	}
	return str

}

func main() {
	args := os.Args[1:]

	//uslovie 4to est argumenty, esli net to vozvrashaet pustotu
	if len(args) == 0 {
		return
	}

	file := "standard.txt"
	if len(args) > 1 && args[1] == "shadow" {
		file = "shadow.txt"
	} else if len(args) > 1 && args[1] == "thinkertoy" {
		file = "thinkertoy.txt"
	}

	//  cikl vyvoda
	massivarg := strings.Split(args[0], "\n")
	for _, z := range massivarg {
		for i := 1; i < 9; i++ {
			y := 0
			arrayprint := make([]string, len(z))
			for _, w := range z {
				data, err := ioutil.ReadFile(file)
				if err != nil {
					fmt.Print("500 internal server error")
					return
				}
				k := 0
				str := ""
				for _, n := range data {
					if n == 10 {
						k++
						continue
					}
					if k == ((int(w)-32)*9 + i) {
						str = str + string(n)
					} else if k > (int(w)-32)*9+i+1 {
						arrayprint[y] = str
						break
					}
				}
				y++
			}
			fmt.Println(Convert(arrayprint))
		}
	}
}
