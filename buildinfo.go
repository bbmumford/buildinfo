/*
 * Copyright (c) 2026 HSTLES / ORBTR Pty Ltd. All Rights Reserved.
 * Queries: licensing@hstles.com
 */
package buildinfo

import (
	"fmt"
	"runtime"
	"time"
)

// Variables set at build time via -ldflags.
// Example:
//
//	go build -ldflags "-X github.com/hstles/library/buildinfo.Version=v1.2.3 \
//	  -X github.com/hstles/library/buildinfo.GitCommit=abc1234 \
//	  -X github.com/hstles/library/buildinfo.BuildDate=20260318.143052 \
//	  -X github.com/hstles/library/buildinfo.AppName=bootstrap.hstles.com \
//	  -X github.com/hstles/library/buildinfo.AppDomain=bootstrap.hstles.com"
var (
	// Version is the semantic version tag (e.g., "v1.2.3").
	// Falls back to "dev" if not set.
	Version = "dev"

	// GitCommit is the short git commit hash (e.g., "abc1234").
	// Falls back to "unknown" if not set.
	GitCommit = "unknown"

	// BuildDate is the UTC build timestamp in compact format (e.g., "20260318.143052").
	// Falls back to "unknown" if not set.
	BuildDate = "unknown"

	// AppName is the endpoint name (e.g., "bootstrap.hstles.com", "devices.orbtr.io").
	// Set per-endpoint at build time.
	AppName = "unknown"

	// AppDomain is the public domain for the endpoint (e.g., "bootstrap.hstles.com").
	// Usually the same as AppName for HSTLES/ORBTR endpoints.
	AppDomain = "unknown"
)

// BuildInfo holds all build metadata in a structured form.
type BuildInfo struct {
	Version   string `json:"version"`
	GitCommit string `json:"commit"`
	BuildDate string `json:"buildDate"`
	BuildID   string `json:"buildId"`
	AppName   string `json:"appName"`
	AppDomain string `json:"appDomain"`
	GoVersion string `json:"goVersion"`
}

// Info returns the full build info struct.
func Info() BuildInfo {
	return BuildInfo{
		Version:   Version,
		GitCommit: GitCommit,
		BuildDate: BuildDate,
		BuildID:   buildID(),
		AppName:   AppName,
		AppDomain: AppDomain,
		GoVersion: runtime.Version(),
	}
}

// Short returns a human-readable one-liner: "v1.2.3 (abc1234, 2026-03-18)"
func Short() string {
	date := BuildDate
	if len(date) >= 8 {
		// Convert "20260318.143052" to "2026-03-18"
		date = date[:4] + "-" + date[4:6] + "-" + date[6:8]
	}
	return fmt.Sprintf("%s (%s, %s)", Version, GitCommit, date)
}

// buildID returns a unique build identifier combining timestamp and commit:
// "20260318.143052-abc1234"
func buildID() string {
	if BuildDate == "unknown" || GitCommit == "unknown" {
		return "dev"
	}
	return fmt.Sprintf("%s-%s", BuildDate, GitCommit)
}

// startTime records when the process started (set on package init).
var startTime = time.Now()

// Uptime returns the process uptime as a duration.
func Uptime() time.Duration {
	return time.Since(startTime)
}

// UptimeString returns uptime in human-readable format (e.g., "2h15m").
func UptimeString() string {
	d := Uptime()
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm%ds", int(d.Minutes()), int(d.Seconds())%60)
	}
	return fmt.Sprintf("%dh%dm", int(d.Hours()), int(d.Minutes())%60)
}
