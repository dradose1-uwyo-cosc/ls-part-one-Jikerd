package main
import "os"
import "simplels/functions"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"."}
	}
	useColor:= functions.IsTerminal(os.Stdout)
	functions.SimpleLS(os.Stdout, args, useColor)
}
