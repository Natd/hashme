package hashme

import "testing"

func TestStringToMd5(t *testing.T) {
	data := StringToMd5("theNatd")

	if data != "d790e581b76c0a9fe3653f4e8c70c4ef" {
		t.Error("Md5 not match")
	}
}

func TestStringToSha1(t *testing.T) {
	data := StringToSha1("theNatd")

	if data != "65902a5061df4e745ced312db55c3673d2b77179" {
		t.Error("Sha1 not match")
	}
}

func TestStringToSha256(t *testing.T) {
	data := StringToSha256("theNatd")

	if data != "b4d05937c4c66eda8b47e6cdf2cf34e92ad8e3e731ba211ee4513436ee462265" {
		t.Error("Sha256 not match")
	}

}

func TestStringToSha512(t *testing.T) {
	data := StringToSha512("theNatd")

	if data != "1871c6500959a8f5ec4ec1260971d9846991961fcfa79644a3693b715baaa1ea223cd40956c6bbe1ca4b9bc99c864d2f654b39151fa7dbfbc10aaee18dbb3084" {
		t.Error("Sha256 not match")
	}
}
