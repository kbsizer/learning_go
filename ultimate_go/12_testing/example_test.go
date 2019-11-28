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

// TestDownload validates that the http Get function can download content.
func TestDownload_simple(t *testing.T) {
	// setup
	goodURL := "https://www.goinggo.net/post/index.xml"
	badURL := "https://www.goinggo.net/bad_url/index.xml"
	statusCode := 200

	t.Log("GIVEN the need to GET content from an HTTP server.")
	{
		t.Logf("\tTest 0 -- WHEN checking %q for status code %d", goodURL, statusCode)
		{
			resp, err := http.Get(goodURL)
			assert(t, "err", nil, err)

			defer resp.Body.Close()

			assert(t, "status code", statusCode, resp.StatusCode)
		}

		t.Logf("\tTest 1 -- WHEN checking %q for status code %d", badURL, statusCode)
		{
			resp, err := http.Get(badURL)
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

// TestDownload validates the http Get function.  It exercises multiple scenarios using a
// TABLE-DRIVEN test.
func TestDownload_tableDriven(t *testing.T) {
	// define inputs and expectations
	tests := []struct {
		url        string
		statusCode int
	}{
		{"https://www.goinggo.net/post/index.xml", http.StatusOK},
		{"https://rss.cnn.com/rss/cnn_topstories.rss", http.StatusNotFound},
	}

	t.Log("GIVEN user wishes to download XML content from a URL...")
	{
		for i, tt := range tests {
			t.Logf("\tTEST %d: WHEN checking %q for status code %d", i, tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t%s: Get call failed: %v", failed, err)
				}
				defer resp.Body.Close()
				t.Logf("\t%v ", resp)

				assert(t, "status code", tt.statusCode, resp.StatusCode)
			}
		}
	}

}
