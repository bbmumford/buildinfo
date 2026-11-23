# buildinfo

buildinfo holds a running binary's build metadata (version, git commit, build date, app name, app domain) as package variables the build stamps with `-ldflags -X`, plus the process uptime since start. Every service reports the same build facts without each `main` package declaring its own version variables. Standard library only.

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

## Dependencies

None: buildinfo uses only the standard library.

## Used by

Importers in the ORBTR estate:

- `github.com/ORBTR/Library/endpoint`: every endpoint reports its build.
- `github.com/ORBTR/api.orbtr.ai`, `github.com/ORBTR/worker/cmd/orbtr-worker`
  and `github.com/ORBTR/xyz/runtime/cmd/orbtrd`.
- `github.com/ORBTR/app.orbtr.io`, `github.com/ORBTR/help.orbtr.io`,
  `github.com/ORBTR/mcp.orbtr.io`, `github.com/ORBTR/orbtr.io` and
  `github.com/ORBTR/constellation/domain/foundation/notification`.
- `github.com/bbmumford/flyio`.

## Platforms and build

Pure Go, no cgo and no build tags. The module's `go` directive is 1.25.0.

## Licence

MIT. Copyright (c) 2026 Brandon Miller-Mumford. The full text is in
[LICENSE](LICENSE).

## Contributing

[AGENTS.md](AGENTS.md) holds the repository's two rules, for people and coding
agents alike: no temporary documents in the repository, and conventional commit
messages that stand on their own. [CLAUDE.md](CLAUDE.md) imports it for Claude
Code.
