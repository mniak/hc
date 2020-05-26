package lib

import (
	"fmt"
	"net/url"
	"path"

	"github.com/go-resty/resty/v2"
)

func CheckLiveness(baseUrl, urlpath string, verbose bool) error {
	fullUrl, err := url.Parse(baseUrl)
	if err != nil {
		return fmt.Errorf("Not a valid URL: %s\n", fullUrl)
	}
	fullUrl.Path = path.Join(fullUrl.Path, urlpath)
	if fullUrl.Scheme == "" {
		fullUrl.Scheme = "https"
	}
	resp, err := resty.New().
		SetDebug(verbose).
		NewRequest().
		Get(fullUrl.String())
	if err != nil {
		return fmt.Errorf("Error checking liveness: %s\n", err.Error())
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("The site is not alive. Response status %s\n", resp.Status())
	}
	return nil
}
