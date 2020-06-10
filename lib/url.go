package lib

import (
	"net"
	"net/url"
	"path"
	"strconv"
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

	if host, port, err := net.SplitHostPort(b); err == nil {
		if _, err = strconv.Atoi(port); err == nil && !strings.Contains(host, "//") {
			return makeurl("//"+b, p)
		}
	}

	u, err := url.Parse(b)
	if err != nil {
		return nil, err
	}

	if u.Host == "" {
		u = &url.URL{
			Host: b,
		}
	}

	if u.Host == "" {
		u.Host = "localhost"
	}

	u.Path = path.Join(u.Path, p)
	return normalize(u), nil
}
