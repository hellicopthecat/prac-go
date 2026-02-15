package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/hellicopthecat/learngo/rest"
)

func usage() {
	fmt.Printf("Welcome to len less 2\n\n")
	fmt.Printf("Please use the following cmd\n\n")
	fmt.Printf("-port=4000 : 	Set Port of the Server\n")
	fmt.Printf("-mode=rest : 	Choose Between 'html' and 'rest'\n")
	os.Exit(0)
}
func Start() {
	// if len(os.Args) < 2 {
	// 	usage()
	// }

	// rest := flag.NewFlagSet("rest", flag.ExitOnError)

	// portFlag := rest.Int("port", 4000, "Set Port Server")

	// switch os.Args[1] {
	// case "explorer":
	// 	fmt.Println("Start Explorer")
	// case "rest":
	// 	rest.Parse(os.Args[2:])
	// 	fmt.Println("Start REST API")
	// default:
	// 	usage()
	// }
	// if rest.Parsed() {
	// 	fmt.Println(portFlag)
	// 	fmt.Println("start Server")
	// }

	if len(os.Args) < 1 {
		usage()
	}
	port := flag.Int("port", 4000, "Set Port of the Server")
	mode := flag.String("mode", "rest", "Choose Between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
	default:
		usage()
	}
	fmt.Println(*port, *mode)
}
