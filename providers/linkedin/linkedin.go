package linkedin

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("linkedin", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for Linkedin.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf(
		"https://www.linkedin.com/search/results/all/?keywords=%s",
		url.QueryEscape(q))
}

// Tags returns the tags relevant to this provider.
func (p *Provider) Tags() []string {
	return []string{"social"}
}
