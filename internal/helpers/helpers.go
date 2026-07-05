package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func CreateJobID() string {
	buf := make([]byte, 16)

	_, _ = rand.Read(buf)
	buf[6] = (buf[6] & 0x0F) | 0x40
	buf[8] = (buf[8] & 0x3F) | 0x80

	uuid := make([]byte, 36)
	uuid[8] = '-'
	uuid[13] = '-'
	uuid[18] = '-'
	uuid[23] = '-'
	hex.Encode(uuid[0:8], buf[0:4])
	hex.Encode(uuid[9:13], buf[4:6])
	hex.Encode(uuid[14:18], buf[6:8])
	hex.Encode(uuid[19:23], buf[8:10])
	hex.Encode(uuid[24:36], buf[10:16])

	return string(uuid)
}
