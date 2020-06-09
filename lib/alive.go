package lib

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// LivenessCheck checks the liveness of an endpoint
func LivenessCheck(base, path string, verbose bool) error {
	url, err := makeurl(base, path)
	if err != nil {
		return fmt.Errorf("invalid URL: %s/%s", base, path)
	}
	resp, err := resty.New().
		SetDebug(verbose).
		NewRequest().
		Get(url.String())
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("the site is not alive. response status %s", resp.Status())
	}
	return nil
}
