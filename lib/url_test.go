package lib

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeUrl(t *testing.T) {
	list := []struct {
		url         string
		path        string
		expectedUrl string
	}{
		// No domain (should assume localhost)
		{":1234", "", "http://localhost:1234"},
		{":1234", "/path", "http://localhost:1234/path"},

		{"http://:1234", "", "http://localhost:1234"},
		{"http://:1234", "/path", "http://localhost:1234/path"},

		{"https://:1234", "", "https://localhost:1234"},
		{"https://:1234", "/path", "https://localhost:1234/path"},

		// localhost
		{"localhost", "", "http://localhost"},
		{"localhost", "/path", "http://localhost/path"},
		{"localhost:1234", "", "http://localhost:1234"},
		{"localhost:1234", "/path", "http://localhost:1234/path"},

		{"http://localhost:1234", "", "http://localhost:1234"},
		{"http://localhost:1234", "/path", "http://localhost:1234/path"},

		{"https://localhost:1234", "", "https://localhost:1234"},
		{"https://localhost:1234", "/path", "https://localhost:1234/path"},

		// IP address
		{"10.133.1.2", "", "http://10.133.1.2"},
		{"10.133.1.2", "/path", "http://10.133.1.2/path"},
		{"10.133.1.2:1234", "/path", "http://10.133.1.2:1234/path"},
		{"10.133.1.2:1234", "", "http://10.133.1.2:1234"},

		{"http://10.133.1.2:1234", "", "http://10.133.1.2:1234"},
		{"http://10.133.1.2:1234", "/path", "http://10.133.1.2:1234/path"},

		{"https://10.133.1.2:1234", "", "https://10.133.1.2:1234"},
		{"https://10.133.1.2:1234", "/path", "https://10.133.1.2:1234/path"},

		// example.com
		{"example.com", "", "https://example.com"},
		{"example.com", "/path", "https://example.com/path"},
		{"example.com:8080", "", "https://example.com:8080"},
		{"example.com:8080", "/path", "https://example.com:8080/path"},

		{"http://example.com", "", "http://example.com"},
		{"http://example.com", "/path", "http://example.com/path"},
		{"http://example.com:8080", "", "http://example.com:8080"},
		{"http://example.com:8080", "/path", "http://example.com:8080/path"},

		{"https://example.com", "", "https://example.com"},
		{"https://example.com", "/path", "https://example.com/path"},
		{"https://example.com:8080", "", "https://example.com:8080"},
		{"https://example.com:8080", "/path", "https://example.com:8080/path"},

		// example.com/path
		{"example.com/path", "", "https://example.com/path"},
		{"example.com/path", "/subpath", "https://example.com/path/subpath"},
		{"example.com:8080/path", "", "https://example.com/path"},
		{"example.com:8080/path", "/subpath", "https://example.com/path/subpath"},

		{"http://example.com/path", "", "http://example.com/path"},
		{"http://example.com/path", "/subpath", "http://example.com/path/subpath"},
		{"http://example.com:8080/path", "", "http://example.com:8080/path"},
		{"http://example.com:8080/path", "/subpath", "http://example.com:8080/path/subpath"},

		{"https://example.com/path", "", "https://example.com/path"},
		{"https://example.com/path", "/subpath", "https://example.com/path/subpath"},
		{"https://example.com:8080/path", "", "https://example.com:8080/path"},
		{"https://example.com:8080/path", "/subpath", "https://example.com:8080/path/subpath"},
	}

	for _, tt := range list {
		t.Run(fmt.Sprintf("url=%s, path=%s", tt.url, tt.path), func(t *testing.T) {
			actual, err := makeurl(tt.url, tt.path)
			require.NoError(t, err)
			require.EqualValues(t, tt.expectedUrl, actual.String())
		})
	}
}
