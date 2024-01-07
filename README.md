# gomonorepoexp

## Goal

## Repository structure

```
.
├── bar
│   ├── bar.go
│   ├── go.mod
└── foo
    ├── foo.go
    └── go.mod
```

* `foo` package has no dependency
* `bar` package depends on `foo`


### Rules

* Both foo and bar has their own versions (e.g. `foo/v0.0.2`)
* 

## Using go-workspace
I tried to use go-workspace at the root with a hope that I could run unit tests of all the sub-modules in the repo.

### Prepration

Create go.work file with foo bar modules.

```sh
go work init
go work use foo
go work use bar
```

Or in one line..

```sh
go work init foo bar
```

### Run tests
```sh
go test -v ./foo
go test -v ./bar
```

Unfortunately, `./...` does not work...

```sh
go test -v ./...
pattern ./...: directory prefix . does not contain modules listed in go.work or their selected dependencies
```

## Issues I came across

### Submodule in Monorepo
* Atomic changes involing multiple submodules is not easy...
   - You will have to add `replace` directive manually
   - You will need to commit dependency's changes first
   - Then tag new version with the commit (e.g. `tag foo/v0.0.2`)
   - Then change the code that imports the updated dependency
   - Commit the change, then tag a new version (e.g. `tag bar/v0.0.2`)
* `replace` is meant to be fore development purpose. (Parent module ignores the dependency's `replace` directive)
   - See https://go.dev/ref/mod#go-mod-file-replace

### Workspace in Monorepo
* Even using go-workspace, `go test ./...` does not work
* `go test -v ./bar` works, but it ignores the version specified by `require` directive in ./bar/go.mod (this feels like a bug in go-workspaces)



## References

* [A Story of Monorepos in Go](https://betterprogramming.pub/a-story-of-monorepos-in-go-9b8c718d0f75)
* [Building a Monorepo in Golang](https://earthly.dev/blog/golang-monorepo)
* [Golang Workspaces](https://earthly.dev/blog/go-workspaces/)
* [Go Workspaces: Simplifying Multi-Modular Projects](https://bysabbir.medium.com/go-workspaces-simplifying-multi-modular-projects-dc1a489302a)