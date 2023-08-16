package xutil

import "unsafe"

func ByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
