// Package jackbox supplies an interface to the jackbox.tv API, allowing
// developers to check game status and room information, and even interact
// with games as a player or audience member.
package jackbox

import (
	"net/url"
	"path"
)

const (
	API_URL_BASE = "https://ecast.jackboxgames.com"

	// Use the Tor Browser user agent for anonymity.
	USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; rv:68.0) Gecko/20100101 Firefox/68.0"
)

func API_URL(components ...string) *url.URL {
	u, _ := url.Parse(API_URL_BASE)
	for _, component := range components {
		u.Path = path.Join(u.Path, component)
	}
	return u
}
