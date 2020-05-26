package lib

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/go-resty/resty/v2"
)

func HealthCheck(baseUrl, urlpath string, verbose bool) error {
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
		SetResult(&healthCheckResponse{}).
		SetError(&healthCheckResponse{}).
		Get(fullUrl.String())
	if err != nil {
		return err
	}

	if resp.IsError() {
		result, ok := resp.Error().(*healthCheckResponse)
		if !ok {
			return fmt.Errorf("The site is not healthy. Response status %s. The body could not be parsed.", resp.Status())
		}
		if result.IsHealthy {
			return fmt.Errorf("The status code indicates failure, but IsHealthy=true. Response status %s.", resp.Status())
		}
		return formatHealthCheckErrors("The site is not healthy. IsHealthy=false.", *result)
	} else {
		result, ok := resp.Error().(*healthCheckResponse)
		if !ok {
			return fmt.Errorf("The response status code indicates success (%s) but the body could not be parsed.", resp.Status())
		}
		if !result.IsHealthy {
			return formatHealthCheckErrors("The response status code indicates success (%s) but IsHealthy=false.", *result)
		}
	}
	return nil
}

func formatHealthCheckErrors(msg string, hcr healthCheckResponse) error {
	var sb strings.Builder
	sb.WriteString(msg)
	sb.WriteRune('\n')
	for _, r := range hcr.Results {
		if !r.Check.IsHealthy {
			sb.WriteString(fmt.Sprintf("  [fail] %s: %s (%s)", r.Name, r.Check.Message, r.Check.Duration))
			sb.WriteRune('\n')
		}
	}
	return errors.New(sb.String())
}

type healthCheckResponse struct {
	IsHealthy           bool                `json:"IsHealthy"`
	HasRegisteredChecks bool                `json:"HasRegisteredChecks"`
	Results             []healthCheckResult `json:"Results"`
	TotalDuration       string              `json:"TotalDuration"`
}

type healthCheckResult struct {
	Name  string           `json:"Name"`
	Check healthCheckCheck `json:"Check"`
}

type healthCheckCheck struct {
	IsHealthy bool   `json:"IsHealthy"`
	Message   string `json:"Message"`
	Duration  string `json:"Duration"`
}
