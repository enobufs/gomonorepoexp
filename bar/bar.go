package bar

import "github.com/enobufs/gomonorepoexp/foo"

func Name() string {
	return "Bar2 (powered by " + foo.Name() + ")"
}
