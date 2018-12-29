package support

import "os"

// OS is a seam for testing
type OS interface {
	Stat(name string) (os.FileInfo, error)
}

type osFacade struct{}

func (o osFacade) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

var fs OS = osFacade{}
