package tinyurl

import (
	"fmt"
	"strconv"
)

type Codec struct {
	urlList []string
	m       map[string]int
}

func Constructor() Codec {
	return Codec{
		urlList: []string{},
		m:       map[string]int{},
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	count := 0
	i := 0
	for ; i < len(longUrl); i++ {
		if longUrl[i] == '/' {
			count++
			if count == 3 {
				i++
				break
			}
		}
	}
	if i < len(longUrl) {
		prefix := longUrl[:i]
		if c, ok := this.m[longUrl]; ok {
			return prefix + fmt.Sprint(c)
		} else {
			this.m[longUrl] = len(this.urlList)
			this.urlList = append(this.urlList, longUrl)
			return prefix + fmt.Sprint(len(this.urlList)-1)
		}
	}
	return longUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	count := 0
	i := 0
	for ; i < len(shortUrl); i++ {
		if shortUrl[i] == '/' {
			count++
			if count == 3 {
				i++
				break
			}
		}
	}
	if i >= len(shortUrl) {
		return shortUrl
	}
	index, _ := strconv.ParseInt(shortUrl[i:], 10, 32)
	return this.urlList[int(index)]
}
