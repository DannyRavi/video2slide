package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

var map_hash = make(map[string]bool)
var map_keys = make(map[uint32]string)
var append_hash = []string{}

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"

func insertHashSimple(inputHash string) bool {

	if map_hash[inputHash] {
		return false // Already in the map
	}

	if theHolder.itr >= theHolder.itr_ref {
		append_hash = append(append_hash, inputHash)
		map_hash[inputHash] = true
		theHolder.itr = 1
		// logrus.Info(inputHash)
		return true
	}
	theHolder.itr++
	return false

}

type HashInfo struct {
	Md5    string `json:"md5"`
	Sha1   string `json:"sha1"`
	Sha256 string `json:"sha256"`
	Sha512 string `json:"sha512"`
}

func CalculateBasicHashes(rd io.Reader) HashInfo {

	md5 := md5.New()
	sha1 := sha1.New()
	sha256 := sha256.New()
	sha512 := sha512.New()

	// For optimum speed, Getpagesize returns the underlying system's memory page size.
	pagesize := os.Getpagesize()

	// wraps the Reader object into a new buffered reader to read the files in chunks
	// and buffering them for performance.
	readers := bufio.NewReaderSize(rd, pagesize)

	// creates a multiplexer Writer object that will duplicate all write
	// operations when copying data from source into all different hashing algorithms
	// at the same time
	multiWriter := io.MultiWriter(md5, sha1, sha256, sha512)

	// Using a buffered reader, this will write to the writer multiplexer
	// so we only traverse through the file once, and can calculate all hashes
	// in a single byte buffered scan pass.
	//
	_, err := io.Copy(multiWriter, readers)
	if err != nil {
		panic(err.Error())
	}

	var info HashInfo

	info.Md5 = hex.EncodeToString(md5.Sum(nil))
	info.Sha1 = hex.EncodeToString(sha1.Sum(nil))
	info.Sha256 = hex.EncodeToString(sha256.Sum(nil))
	info.Sha512 = hex.EncodeToString(sha512.Sum(nil))

	return info
}

// HasherReader calculates the hash of a byte stream
// As an underlying io.Reader is read from, the hash is updated
type HasherReader struct {
	hash   hash.Hash
	reader io.Reader
}

// NewHasherReader creates a new HasherReader from a provided io.Raeder
func NewHasherReader(r io.Reader) HasherReader {
	hash := sha1.New()
	reader := io.TeeReader(r, hash)
	return HasherReader{hash, reader}
}

// Hash returns the hash value
// Ensure all contents of the underlying io.Reader have been read
func (h HasherReader) Hash() string {
	return hex.EncodeToString(h.hash.Sum(nil))
}

// Read allows HasherReader to conform to io.Reader interface
func (h HasherReader) Read(p []byte) (n int, err error) {
	return h.reader.Read(p)
}
