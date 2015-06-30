package hashme

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

//Md5
func ComputeMd5Checksum(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

func ValidateMd5Checksum(data []byte, sum []byte) bool {
	ourSum := ComputeMd5Checksum(data)
	return bytes.Equal(ourSum, sum)
}

func StringToMd5(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateMd5WithString(data string, sum string) bool {
	ourSum := StringToMd5(data)
	return ourSum == sum
}

//Sha1
func ComputeSha1Checksum(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func StringToSha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateSha1Checksum(data []byte, sum []byte) bool {
	ourSum := ComputeSha1Checksum(data)
	return bytes.Equal(ourSum, sum)
}

func ValidateSha1WithString(data string, sum string) bool {
	ourSum := StringToSha1(data)
	return ourSum == sum
}

//Sha256
func ComputeSha256Checksum(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func StringToSha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateSha256Checksum(data []byte, sum []byte) bool {
	ourSum := ComputeSha256Checksum(data)
	return bytes.Equal(ourSum, sum)
}

func ValidateSha256WithString(data string, sum string) bool {
	ourSum := StringToSha256(data)
	return ourSum == sum
}

//Sha512
func ComputeSha512Checksum(data []byte) []byte {
	h := sha512.New()
	h.Write(data)
	return h.Sum(nil)
}

func StringToSha512(data string) string {
	h := sha512.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateSha512Checksum(data []byte, sum []byte) bool {
	ourSum := ComputeSha512Checksum(data)
	return bytes.Equal(ourSum, sum)
}

func ValidateSha512WithString(data string, sum string) bool {
	ourSum := StringToSha512(data)
	return ourSum == sum
}
