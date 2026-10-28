// Package host resolves which GitHub instance gh-image talks to. It follows the
// gh CLI convention: GH_HOST selects the host, defaulting to github.com.
package host

import "os"

const Default = "github.com"

// Name returns the target hostname.
func Name() string {
	if h := os.Getenv("GH_HOST"); h != "" {
		return h
	}
	return Default
}

// BaseURL returns the https origin of the target host.
func BaseURL() string { return "https://" + Name() }

// UploadsURL returns the bearer upload endpoint, or "" where none exists
// (everything except github.com).
func UploadsURL() string {
	if Name() != Default {
		return ""
	}
	return "https://uploads.github.com"
}
