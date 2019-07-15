package util

import (
	"fmt"
	"net/http"
	"unsafe"
)

//use return value readonly
func ConvertBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

//use return value readonly
func SuccessJSON() []byte {
	json := fmt.Sprintf("{\"code\":\"%s\",\"message\":\"%s\"}\n", http.StatusText(http.StatusOK), "done")
	return ConvertBytes(json)
}

//use return value readonly
func InternalErrorJSON(err error) []byte {
	json := fmt.Sprintf("{\"code\":\"%s\",\"message\":\"%s\"}\n", http.StatusText(http.StatusInternalServerError), err)
	return ConvertBytes(json)
}

//use return value readonly
func AuthErrorJSON(err error) []byte {
	json := fmt.Sprintf("{\"code\":\"%s\",\"message\":\"%s\"}\n", http.StatusText(http.StatusUnauthorized), err)
	return ConvertBytes(json)
}
