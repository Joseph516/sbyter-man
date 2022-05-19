package util

import (
	"net/url"
	"path"
)

func UrlJoin(src string, paths ...string) string {
	u, _ := url.Parse(src) // src: "http://foo"
	tmp := append([]string{u.Path}, paths...)
	u.Path = path.Join(tmp...)
	return u.String()
}
