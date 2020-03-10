package main

import (
"fmt"
"io"
"os"
"bufio"
)

func main(){
	file, err := os.Open("./planets.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	
	defer file.Close()
	reader := bufio.NewReader(file)

	for{
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Println(line)
	}
}

//lol. What the actual fuck?!