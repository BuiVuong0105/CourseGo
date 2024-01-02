package domain

import (
	"github.com/speps/go-hashids"
)

type HashCode struct {
	hashId *hashids.HashID
}

func NewHashCode() *HashCode {

	hd := hashids.NewData()
	hd.Salt = "v2:35056997918317u85346cccx3i@123"
	hd.MinLength = 3
	hd.Alphabet = "abcdefghijkmnopqrstuvwxyz1234567890"

	hashId, _ := hashids.NewWithData(hd)

	return &HashCode{
		hashId: hashId,
	}
}

func (g *HashCode) GetHashID() *hashids.HashID {
	return g.hashId
}

func (g *HashCode) EndCode(prefixId, number int) [3]string {
	hashID := g.GetHashID()
	qri, _ := hashID.Encode([]int{1, prefixId, number})
	qrm, _ := hashID.Encode([]int{2, prefixId, number})
	code, _ := hashID.Encode([]int{3, prefixId, number})
	return [3]string{qri, qrm, code}
}
