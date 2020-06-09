package lib

import (
	"net"
	"net/url"
	"path"
	"strings"
)

func normalize(u *url.URL) *url.URL {
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
		b = "localhost" + b
	}
	u, err := url.Parse(b)
	if err != nil {
		return nil, err
	}
	_, _, err = net.SplitHostPort(b)
	if err == nil {
		u = &url.URL{
			Host: b,
		}
	}
	if u.Host == "" {
		u = &url.URL{
			Host: b,
		}
	}
	u.Path = path.Join(u.Path, p)
	return normalize(u), nil
}
