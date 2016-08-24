package message_parser

import (
	"reflect"
	"unsafe"
)

// Zero-allocation functions for converting []byte to string and visa verse
func StringToByteSlice(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
