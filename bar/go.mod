module github.com/enobufs/gomonorepoexp/bar

replace github.com/enobufs/gomonorepoexp/foo => ../foo

go 1.21.5

require github.com/enobufs/gomonorepoexp/foo v0.0.0-00010101000000-000000000000
