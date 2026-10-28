package host

import "testing"

func TestName(t *testing.T) {
	t.Setenv("GH_HOST", "")
	if Name() != "github.com" || BaseURL() != "https://github.com" || UploadsURL() != "https://uploads.github.com" {
		t.Fatalf("bad defaults: %s %s %s", Name(), BaseURL(), UploadsURL())
	}
	t.Setenv("GH_HOST", "ghe.example.com")
	if Name() != "ghe.example.com" || BaseURL() != "https://ghe.example.com" || UploadsURL() != "" {
		t.Fatalf("bad override: %s %s %s", Name(), BaseURL(), UploadsURL())
	}
}
