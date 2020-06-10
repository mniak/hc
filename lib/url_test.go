package lib

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeUrl(t *testing.T) {
	list := []struct {
		url         string
		path        string
		expectedUrl string
	}{
		{"localhost", "", "http://localhost"},
		{"10.133.1.2", "", "http://10.133.1.2"},
		{"localhost:1234", "/path", "http://localhost:1234/path"},
		{"10.133.1.2:1234", "/path", "http://10.133.1.2:1234/path"},

		{":1234", "", "http://localhost:1234"},
		{"localhost:1234", "", "http://localhost:1234"},
		{"10.133.1.2:1234", "", "http://10.133.1.2:1234"},
		{":1234", "/path", "http://localhost:1234/path"},
		{"localhost:1234", "/path", "http://localhost:1234/path"},
		{"10.133.1.2:1234", "/path", "http://10.133.1.2:1234/path"},

		{"http://:1234", "", "http://localhost:1234"},
		{"http://localhost:1234", "", "http://localhost:1234"},
		{"http://10.133.1.2:1234", "", "http://10.133.1.2:1234"},
		{"http://:1234", "/path", "http://localhost:1234/path"},
		{"http://localhost:1234", "/path", "http://localhost:1234/path"},
		{"http://10.133.1.2:1234", "/path", "http://10.133.1.2:1234/path"},

		{"https://:1234", "", "https://localhost:1234"},
		{"https://localhost:1234", "", "https://localhost:1234"},
		{"https://10.133.1.2:1234", "", "https://10.133.1.2:1234"},
		{"https://:1234", "/path", "https://localhost:1234/path"},
		{"https://localhost:1234", "/path", "https://localhost:1234/path"},
		{"https://10.133.1.2:1234", "/path", "https://10.133.1.2:1234/path"},

		{"example.com", "", "https://example.com"},
		{"http://example.com", "", "http://example.com"},
		{"https://example.com", "", "https://example.com"},
		{"example.com", "/path", "https://example.com/path"},
		{"http://example.com", "/path", "http://example.com/path"},
		{"https://example.com", "/path", "https://example.com/path"},

		{"example.com:8080", "", "https://example.com:8080"},
		{"http://example.com:8080", "", "http://example.com:8080"},
		{"https://example.com:8080", "", "https://example.com:8080"},
		{"example.com:8080", "/path", "https://example.com:8080/path"},
		{"http://example.com:8080", "/path", "http://example.com:8080/path"},
		{"https://example.com:8080", "/path", "https://example.com:8080/path"},
	}
	for i, tt := range list {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual, err := makeurl(tt.url, tt.path)
			require.NoError(t, err)
			require.EqualValues(t, tt.expectedUrl, actual.String())
		})
	}
}
