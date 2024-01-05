package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	fmt.Println("AromaGuard Unpacker")
	fmt.Println(" > This was the dumbest garbage that I have ever seen. Next time, make it worth my time.")
	fmt.Println(" > Developed with love by AdolfCracks <3.")
	fmt.Println()

	programArguments := os.Args

	if len(programArguments) == 2 {
		path := os.Args[1]
		data, err := os.ReadFile(path)

		if err != nil {
			panic(err)
		}
		fmt.Println(" > Executing insane decoding logic...")
		decodedContent, err := base64.StdEncoding.DecodeString(string(data))

		if err != nil {
			panic(err)
		}

		os.WriteFile(path+"-unpacked.jar", decodedContent, 0644)
		fmt.Println(" > Wrote output to " + (path + "-unpacked.jar"))
	} else {
		fmt.Println("You can drag the file you wish to unpack directly into the executable!")
	}
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
