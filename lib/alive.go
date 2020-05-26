package lib

import (
	"fmt"
	"net/url"
	"path"

	"github.com/go-resty/resty/v2"
)

// LivenessCheck checks the liveness of an endpoint
func LivenessCheck(baseURL, urlpath string, verbose bool) error {
	fullURL, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("not a valid URL: %s", fullURL)
	}
	fullURL.Path = path.Join(fullURL.Path, urlpath)
	if fullURL.Scheme == "" {
		fullURL.Scheme = "https"
	}
	resp, err := resty.New().
		SetDebug(verbose).
		NewRequest().
		Get(fullURL.String())
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("the site is not alive. response status %s", resp.Status())
	}
	return nil
}
