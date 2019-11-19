// Sample test to show how to write a basic unit test in Golang.
package example1

// PACKAGE NAMES IN TESTS: Use the same package as the production code for whitebox testing
// (accessing unexported functions, methods and definitions); use production package plus
// "_test" to perform backbox testing (only exported definitions are visible to the test)
//

import (
	"net/http"
	"testing"
)

// Handy constants to make successes and failures stand out
const succeed = "\u2713" // checkmark
const failed = "\u2717"  // "X"

// USING GO TEST:
//       $ go test
// This will suppress informational messages when tests pass, but will automatically
// switch to verbose mode (see below) if there is a failure.
//       $ go test -v
// Force go test into verbose mode (writing out messages) whether tests are successful or
// not.

// HANDY TESTING METHODS:
//   t.Log() / t.Logf()     -- add information; does not affect result of test
//   t.Error() / t.Errorf() -- log message and set status of test to FAILED
//   t.Fatal() / t.Fatalf() -- log message and end test *immediately* with status of FAILED

// EXTRA BRACKETS TO ADD READIBILITY AND DEMARCATE SCOPE
// Using the "given-when-then" framework (see: https://martinfowler.com/bliki/GivenWhenThen.html)
// (Bill Kennedy calls it "given-when-should")



// TestDownload validates that the http Get function can download content.
func TestDownload(t *testing.T) {
	// setup
	goodUrl := "https://www.goinggo.net/post/index.xml"
	badUrl := "https://www.goinggo.net/bad_url/index.xml"
	statusCode := 200

	t.Log("GIVEN the need to GET content from an HTTP server.")
	{
		t.Logf("\tTest 0 -- WHEN checking %q for status code %d", goodUrl, statusCode)
		{
			resp, err := http.Get(goodUrl)
			assert(t, "err", nil, err)

			defer resp.Body.Close()

			assert(t, "status code", statusCode, resp.StatusCode)
		}

		t.Logf("\tTest 1 -- WHEN checking %q for status code %d", badUrl, statusCode)
		{
			resp, err := http.Get(badUrl)
			assert(t, "err", nil, err)

			defer resp.Body.Close()

			assert(t, "status code", statusCode, resp.StatusCode)
		}
	}
}

func assert(t *testing.T, msg string, expected interface{}, actual interface{}) {
	if expected == actual {
		t.Logf("\t%s -- %s: %v", succeed, msg, actual)
	} else {
		t.Errorf("\t%s -- %s: Expected %v, got %v", failed, msg, expected, actual)
	}
}
