// Source: http://bit.ly/learning-go-00001

package main

func main(){
	func main() {
		// use an os.File as source for alphaReader
		file, err := os.Open("./alpha_reader3.go")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
		
		reader := newAlphaReader(file)
		p := make([]byte, 4)
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			fmt.Print(string(p[:n]))
		}
		fmt.Println()
	}
}