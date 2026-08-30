package main

import "testing"

func TestNormalizeAssetURL(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"/assets/images/foo.png", "/assets/images/foo.png"},
		{"http://localhost:8080/assets/images/foo.png", "/assets/images/foo.png"},
		{"http://localhost/assets/icons/bar.svg", "/assets/icons/bar.svg"},
		{"https://example.com/photo.jpg", "https://example.com/photo.jpg"},
	}

	for _, tt := range tests {
		if got := normalizeAssetURL(tt.in); got != tt.want {
			t.Errorf("normalizeAssetURL(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
