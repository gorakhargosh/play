package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func sha1_hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func md5_hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func main() {
	s := "sha1 this string"
	fmt.Println(s)
	fmt.Printf("%s\n", sha1_hash(s))
	fmt.Printf("%s\n", md5_hash(s))
}
