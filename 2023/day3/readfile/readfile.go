package readfile

import (
	"bytes"
	"os"
)

func ReadLines(inputFile string) (chan string, error) {
	out := make(chan string, 1)
	file, err := os.Open(inputFile)
	if err != nil {
		return out, err
	}

	go func() {
		defer file.Close()

		line := ""
		for {
			var data = make([]byte, 1)
			read, err := file.Read(data)
			if err != nil {
				break
			}
			chunk := data[:read]
			split := bytes.SplitN(chunk, []byte("\n"), 2)
			if len(split) == 2 {
				line += string(split[0])
				out <- line
				line = ""
				line += string(split[1])
				continue
			}
			line += string(chunk)
		}

		if line != "" {
			out <- line
		}
		close(out)
	}()

	return out, nil

}
