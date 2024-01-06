package bar

import "github.com/enobufs/gomonorepoexp/foo"

func Name() string {
	return "Bar. (powered by " + foo.Name() + ")"
}
