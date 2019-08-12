// Demonstrates command-line parsing and basic file I/O.
// Based on https://www.freecodecamp.org/news/writing-command-line-applications-in-go-2bc8c0ace79d/
package ken

import (
	"fmt"
	"log"
	"os"

	flag "github.com/ogier/pflag" // https://github.com/spf13/pflag
)

// init() is a built-in function.  If present, it is executed before main()
func init() {
	log.Println("init() called...")
}

// main() is a built-in function. If present, the file can be a standalone executable.
func main() {
	log.Println("main() called...")
	var srcPath, destPath string

	// using "os" package to print all command-line arguments
	for ndx, value := range os.Args {
		log.Printf("os.Args[%d] = %s\n", ndx, value)
	}

	// using "flags" package to define flags/options and read their values
	// For more powerful CLI support, see: https://github.com/spf13/cobra
	// Params
	// 1:target variable, 2: long (POSIX) flag name, 3: short flag name, 4: default value, 5: "usage" text
	flag.StringVarP(&srcPath, "source", "s", "", "Source Path")
	flag.StringVarP(&destPath, "dest", "d", "", "Destination Path")
	log.Printf("before calling Parse(), Parsed() = %t", flag.Parsed())
	flag.Parse()
	log.Printf("after calling Parse(), Parsed() = %t", flag.Parsed())

	// use flags package to print remaining (non-flag) command-line arguments
	for ndx, value := range flag.Args() {
		log.Printf("flag.Args[%d] = %s\n", ndx, value)
	}
	fmt.Printf("Copying from %s to %s...\n", srcPath, destPath)
}
