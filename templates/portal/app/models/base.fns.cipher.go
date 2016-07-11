package models

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/satori/go.uuid"
)

func cipher_md5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func cipher_uuid(str string) string {
	return uuid.NewV4().String()
}

func EncodeID(number int64) string {
	if number == 0 {
		return ""
	}
	a := mix(uint64(number))
	return strings.ToUpper(strconv.FormatInt(setVersion(a), 36))
}

func DecodeID(str string) int64 {
	if str == "" || str == "74" {
		return 0
	}
	number, _ := strconv.ParseUint(str, 36, 64)
	return int64(demix(number))
}

func mix(number uint64) []uint64 {
	var ver uint64 = 1
	ret := number
	digit := 0

	for ret > 0 {
		digit++
		ret = ret >> 3
	}

	var i uint64
	md := uint64((digit-1)/5 + 1)
	mix := uint64(number & ((1 << (3 * md)) - 1))

	for digit > 0 {

		md--
		ret += (((number & ((1 << 15) - 1)) + ((mix & (((1 << 3) - 1) << (3 * md))) << (15 - 3*md))) << i)

		number = number >> 15
		digit -= 5
		i += 18
	}
	number = ret
	return []uint64{ver, number}
}

func setVersion(mixed []uint64) int64 {
	return int64(((mixed[1] >> 8) << 12) + (mixed[0] << 8) + (mixed[1] & 255))
}

func getVersion(number uint64) []uint64 {
	return []uint64{(number >> 8) & 15, ((number >> 12) << 8) + (number & 255)}
}

func demix(number uint64) uint64 {
	vs := getVersion(number)
	number = vs[1]
	switch vs[0] {
	case 1:
		var dig uint64
		var ret uint64
		for number > 0 {
			ret += ((number & ((1 << 15) - 1)) << dig)
			number = number >> 18
			dig += 15
		}
		number = ret
	}
	return number
}
