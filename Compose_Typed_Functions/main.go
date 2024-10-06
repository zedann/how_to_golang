package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type TransferFunc func(string) string

type Server struct {
	transferFunc TransferFunc
}

func (s *Server) handleRequest(filename string) error {

	newFilename := s.transferFunc(filename)
	fmt.Println("new file name : ", newFilename)

	return nil
}

func hashFilename(filename string) string {
	hash := sha256.Sum256([]byte(filename))
	newFilename := hex.EncodeToString(hash[:])
	return newFilename
}

func prefixFilename(prefix string) TransferFunc {
	return func(filename string) string {
		return prefix + filename
	}
}

func prefixHashFilename(prefix string) TransferFunc {
	return func(filename string) string {
		return prefix + hashFilename(filename)
	}
}

func main() {
	server := &Server{
		transferFunc: prefixFilename("pre_"),
	}

	server.handleRequest("profile_pic.jpg")
}
