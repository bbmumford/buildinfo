# buildinfo

A running binary's build metadata (version, git commit, build date, app name, app domain) as package variables the build stamps with `-ldflags -X`, plus the process uptime since start. Every service reports the same build facts without each `main` package declaring its own version variables. Standard library only.

## Use

```sh
go get github.com/bbmumford/buildinfo

go build -ldflags "\
  -X github.com/bbmumford/buildinfo.Version=v1.2.3 \
  -X github.com/bbmumford/buildinfo.GitCommit=abc1234 \
  -X github.com/bbmumford/buildinfo.BuildDate=20260318.143052 \
  -X github.com/bbmumford/buildinfo.AppName=api.example.com \
  -X github.com/bbmumford/buildinfo.AppDomain=api.example.com"
```

An unstamped `Version` reads `dev`; the other variables read `unknown`.

```go
log.Printf("starting %s", buildinfo.Short()) // "v1.2.3 (abc1234, 2026-03-18)"
info := buildinfo.Info()                     // the whole BuildInfo struct
up := buildinfo.UptimeString()               // e.g. "2h15m"; Uptime() returns the time.Duration
```

## Consumers

ORBTR library, ai worker, xyz and the api.orbtr.ai endpoint, io constellation, tools and the orbtr.io, mcp, help and app endpoints, and the flyio package.
