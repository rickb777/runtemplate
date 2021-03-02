package support

import (
	"os"
	"time"
)

type FileMeta struct {
	Path     string
	Name     string
	ModTime  time.Time
	Embedded string
}

func EmbeddedFileMeta(path, name string, content string) FileMeta {
	return FileMeta{path, name, time.Time{}, content}
}

func SingleFileMeta(path, name string) FileMeta {
	Debug("stat %q\n", path)
	if path == "" {
		return FileMeta{path, "", time.Time{}, ""}
	}

	info, err := fs.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			Debug("%q does not exist.\n", path)
			return FileMeta{path, "", time.Time{}, ""}
		} else {
			Fail(err)
		}
	}

	if name == "" || name[0] == '/' {
		name = info.Name()
	}
	return FileMeta{path, name, info.ModTime(), ""}
}

func NewFileMeta(includeEmpties bool, paths ...string) []FileMeta {
	result := make([]FileMeta, len(paths))

	i := 0
	for _, p := range paths {
		fm := SingleFileMeta(p, "")
		if fm.Exists() {
			result[i] = fm
			i += 1
		} else {
			if includeEmpties {
				result[i] = fm
				i += 1
			}
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
