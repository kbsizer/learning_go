package hello

// NOTE: Each different major version (v1, v2, and so on) of a Go module uses a different
// module path: starting at v2, the path must end in the major version. In the example, v3
// of rsc.io/quote is no longer rsc.io/quote: instead, it is identified by the module path
// rsc.io/quote/v3. This convention is called semantic import versioning, and it gives
// incompatible packages (those with different major versions) different names.

import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

// Hello demonstrates Go Modules.  See: https://blog.golang.org/using-go-modules
func Hello() string {
	return quote.Hello()
}

// Proverb returns a Go concurrency proverb
func Proverb() string {
	return quoteV3.Concurrency()
}
