package support

import (
	"time"
	"os"
)

type FileMeta struct {
	Path    string
	Name    string
	ModTime time.Time
}

func NewFileMeta(includeEmpties bool, paths ...string) []FileMeta {
	result := make([]FileMeta, len(paths))

	i := 0
	for _, p := range paths {
		Debug("stat %s\n", p)
		info, err := os.Stat(p)
		if err != nil {
			if os.IsNotExist(err) {
				Info("%s does not exist.\n", p)
				if includeEmpties {
					result[i] = FileMeta{p, "", time.Time{}}
					i += 1
				}
			} else {
				Fail(err)
			}
		} else {
			result[i] = FileMeta{p, info.Name(), info.ModTime()}
			i += 1
		}
	}

	return result[:i]
}

func YoungestFile(files ...FileMeta) FileMeta {
	if len(files) == 0 {
		return FileMeta{}
	}
	youngest := files[0]
	for _, file := range files {
		youngest = youngest.Younger(file)
	}
	return youngest
}

func (file FileMeta) Exists() bool {
	return !file.ModTime.IsZero()
}

func (file FileMeta) Younger(other FileMeta) FileMeta {
	if other.ModTime.After(file.ModTime) {
		return other
	}
	return file
}
