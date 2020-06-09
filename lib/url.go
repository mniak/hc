package lib

import (
	"net"
	"net/url"
	"path"
	"strings"
)

func normalize(u *url.URL) *url.URL {

	if strings.HasPrefix(u.Host, ":") {
		u.Host = "localhost" + u.Host
	}

	h := u.Hostname()
	ip := net.ParseIP(h)
	if h == "localhost" || ip != nil {
		if u.Scheme == "" {
			u.Scheme = "http"
		}
	} else {
		if u.Scheme == "" {
			u.Scheme = "https"
		}
	}
	return u
}

func makeurl(b, p string) (*url.URL, error) {
	if strings.HasPrefix(b, ":") {
		return makeurl("localhost"+b, p)
	}

	_, _, err := net.SplitHostPort(b)
	if err == nil {
		return makeurl("http://"+b, p)
	}

	u, err := url.Parse(b)
	if err != nil {
		return nil, err
	}

	if u.Host == "" {
		u = &url.URL{
			Scheme: "http",
			Host:   b,
		}
	}

	if u.Host == "" {
		u.Host = "localhost"
	}

	u.Path = path.Join(u.Path, p)
	return normalize(u), nil
}
