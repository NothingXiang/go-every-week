package gee_cache

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(cloneBytes(v.b))
}

// ByteSlice return the copy of slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// Len return the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
