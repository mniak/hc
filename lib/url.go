package lib

import (
	"fmt"
	"net"
	"net/url"
	"path"
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
	u, err := url.Parse(b)
	if err != nil {
		return nil, fmt.Errorf("not a valid URL: %s", b)
	}
	u.Path = path.Join(u.Path, p)
	return normalize(u), nil
}
