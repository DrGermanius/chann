package chann

import (
	"sync/atomic"
	"unsafe"
)

type hchan struct {
	qcount   uint
	dataqsiz uint
	buf      unsafe.Pointer
	elemsize uint16
	closed   uint32
}

const (
	prtlen      = unsafe.Sizeof(uint(0))
	closeoffset = unsafe.Offsetof(hchan{}.closed)
)

//go:noinline
func IsClosed[T any](ch chan T) bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer((*(*uintptr)(unsafe.Pointer(uintptr(unsafe.Pointer(&ch)) + prtlen)))+closeoffset))) != 0
}
