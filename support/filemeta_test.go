package support

import (
	"testing"
	"os"
	"time"
)

type fileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
	isDir   bool
}

func (fi fileInfo) Name() string {
	return fi.name
}

func (fi fileInfo) Size() int64 {
	return fi.size
}

func (fi fileInfo) Mode() os.FileMode {
	return fi.mode
}

func (fi fileInfo) ModTime() time.Time {
	return fi.modTime
}

func (fi fileInfo) IsDir() bool {
	return fi.isDir
}

func (fi fileInfo) Sys() interface{} {
	return nil
}

var _ os.FileInfo = fileInfo{}

//-------------------------------------------------------------------------------------------------

type osStub struct {
	fi os.FileInfo
	e  error
}

func (o osStub) Stat(name string) (os.FileInfo, error) {
	return o.fi, o.e
}

var _ OS = osStub{}

//-------------------------------------------------------------------------------------------------

func TestBlank(t *testing.T) {
	m := SingleFileMeta("", "")
	if m.Exists()  {
		t.Error("Expected not to exist")
	}
}

func TestHappy(t *testing.T) {
	now := time.Now()
	fi := fileInfo{"foo", 123, 0, now, false}
	fs = osStub{fi, nil} // global
	m := NewFileMeta(true, "/a/b/c/foo")
	if len(m) != 1 {
		t.Errorf("Expected 1 but got %d", len(m))
	}
	if m[0].Path != "/a/b/c/foo"  {
		t.Errorf("Expected '/a/b/c/foo' but got '%s'", m[0].Path)
	}
	if m[0].Name != "foo"  {
		t.Errorf("Expected 'foo' but got '%s'", m[0].Name)
	}
	if m[0].ModTime != now  {
		t.Errorf("Expected %v but got %v", now, m[0].ModTime)
	}
	if !m[0].Exists()  {
		t.Error("Expected exists")
	}
}

func TestYoungest(t *testing.T) {
	now := time.Now()
	a := FileMeta{"", "a", now.Add(-2*time.Minute)}
	b := FileMeta{"", "b", now.Add(-1*time.Minute)}
	c := FileMeta{"", "c", now}

	y1 := YoungestFile(c, a, b)
	if y1 != c {
		t.Errorf("Expected %v but got %v", c, y1)
	}

	y2 := YoungestFile(a, b, c)
	if y2 != c {
		t.Errorf("Expected %v but got %v", c, y2)
	}
}
