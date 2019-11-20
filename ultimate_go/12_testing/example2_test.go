package example2

import "testing"

import "net/http"

// TestDownload validates the http Get function.  It exercises multiple scenarios using a
// TABLE-DRIVEN test.
func TestDownload(t *testing.T) {
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
			t.Log("\tTEST %d: WHEN checking %q for status code %d", i, tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t%s: Get call failed: %v", failed, err)
				}
				t.Logf("\t%s ")
			}
		}
	}

}
