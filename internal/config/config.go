package config

import (
	"os"
	"strings"
)

var Version = "3.0.1"

func APIURL() string {
	if v := os.Getenv("API_URL"); v != "" {
		return strings.TrimRight(v, "/")
	}
	return "https://cloud.untag.ru"
}

func PostHogAPIKey() string {
	if v := os.Getenv("POSTHOG_API_KEY"); v != "" {
		return v
	}
	return ""
}

func PostHogHost() string {
	if v := os.Getenv("POSTHOG_HOST"); v != "" {
		return v
	}
	return "https://eu.i.posthog.com"
}
