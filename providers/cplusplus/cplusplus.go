package cplusplus

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("cplusplus", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for cplusplus.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://www.cplusplus.com/search.do?q=%s", url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"cpp"}
}
