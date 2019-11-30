package example2

import (
	"net/http"
	"testing"
)

// Handy constants to make successes and failures stand out
const succeed = "\u2713" // checkmark
const failed = "\u2717"  // "X"

// TestDownload validates the http Get function.  It exercises multiple scenarios using a
// TABLE-DRIVEN test.
func TestDownload(t *testing.T) {
	// define inputs and expectations
	// (common idiom is to use a slice of anonymous struct here)
	tests := []struct {
		url        string // input
		statusCode int    // expected output
	}{
		// TABLE DRIVING TESTS
		{"https://www.goinggo.net/post/index.xml", http.StatusOK},
		{"https://rss.cnn.com/rss/cnn_topstories.rss", http.StatusNotFound},
	}

	t.Log("GIVEN user wishes to download XML content from a URL...")
	{
		// Iterate over each test case in the table
		for i, tt := range tests {
			t.Logf("\tTEST %d: WHEN checking %q for status code %d", i, tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t%s -- HTTP GET call failed: %v", failed, err)
				}
				t.Logf("\t%s -- HTTP GET call", succeed)

				defer resp.Body.Close()

				if resp.StatusCode == tt.statusCode {
					t.Logf("\t%s -- received expected status code: %d", succeed, tt.statusCode)
				} else {
					t.Errorf("\t%s -- Expected status code %d, but got: %d", failed, tt.statusCode, resp.StatusCode)
				}
			}
		}
	}
}
