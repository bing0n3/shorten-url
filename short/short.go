package short

import (
	"crypto/md5"
	"encoding/hex"
)

type shorter struct {
	url string
}

func short() {

}

func (shorter *shorter) IdToShortUrl(id int64) string {
	strMap := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	runeMap := []rune(strMap)
	shortUrl := []rune(nil)
	for (id > 0) {
		shortUrl = append(shortUrl, runeMap[id % 62])
		id /= 62
	}
	shortUrlStr := Reverse(shortUrl)

	return shortUrlStr
}

func (shorter *shorter) ShortUrlToId(shortUrl string) int64{
	var id int64
	id = 0
	runes := []rune(shortUrl)
	for i := 0; i < len(runes); i++ {
		switch {
		case  'a' <= runes[i] && runes[i] <= 'z':
			id = id * 62 + int64(runes[i] - 'a')
			break
		case  'A' <= runes[i] && runes[i] <= 'Z':
			id = id * 62 + int64(runes[i] - 'A') + 26
			break
		case  '0' <= runes[i] && runes[i] <= '9':
			id = id * 62 + int64(runes[i] - '0') + 52
			break
		}
	}
	return id
}

func GetMD5Hash(url string) string {
	hasher := md5.New()
	hasher.Write([]byte(url))
	return hex.EncodeToString(hasher.Sum(nil))
}


func Reverse(r []rune) string {
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}