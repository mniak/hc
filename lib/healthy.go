package lib

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// HealthCheck checks the health of an endpoint
func HealthCheck(base, path string, verbose bool) error {
	url, err := makeurl(base, path)
	if err != nil {
		return fmt.Errorf("not a valid URL: %s", url)
	}
	resp, err := resty.New().
		SetDebug(verbose).
		NewRequest().
		SetResult(&healthCheckResponse{}).
		SetError(&healthCheckResponse{}).
		Get(url.String())
	if err != nil {
		return err
	}

	if resp.IsError() {
		result, ok := resp.Error().(*healthCheckResponse)
		if !ok {
			return fmt.Errorf("the site is not healthy. response status %s. the body could not be parsed", resp.Status())
		}
		if result.IsHealthy {
			return fmt.Errorf("the status code indicates failure, but IsHealthy=true. response status %s", resp.Status())
		}
		return formatHealthCheckErrors("the site is not healthy. IsHealthy=false.", *result)
	}
	result, ok := resp.Result().(*healthCheckResponse)
	if !ok {
		return fmt.Errorf("the response status code indicates success (%s) but the body could not be parsed", resp.Status())
	}
	if !result.IsHealthy {
		return formatHealthCheckErrors("the response status code indicates success (%s) but IsHealthy=false", *result)
	}
	return nil
}

func formatHealthCheckErrors(msg string, hcr healthCheckResponse) error {
	var sb strings.Builder
	sb.WriteString(msg)
	sb.WriteRune('\n')
	for _, r := range hcr.Results {
		if !r.Check.IsHealthy {
			sb.WriteString(fmt.Sprintf("  [unhealthy] %s: %s (%s)", r.Name, r.Check.Message, r.Check.Duration))
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
