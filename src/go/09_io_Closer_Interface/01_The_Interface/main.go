
import "io"


// look into the io Package:
rc := &io.ReadCloser{}
wc := &io.WriteCloser{}
rwc := &io.ReadWriteCloser{}

// looks like this:
type ReadCloser interface {
	Reader
	Closer
}
type WriteCloser interface {
	Writer
	Closer
}
type ReadWriteCloser interface {
	Reader
	Writer
	Closer
}
