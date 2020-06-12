/*
__author__ = 'lawtech'
__date__ = '2020/06/12 1:30 下午'
*/

package _535

import (
	"strconv"
	"strings"
)

type Codec struct {
	hash  map[int]string
	count int
}

func Constructor() Codec {
	return Codec{
		hash:  make(map[int]string),
		count: 0,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.hash[this.count] = longUrl
	tmp := this.count
	this.count++
	return "http://tinyurl.com/" + strconv.Itoa(tmp)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	num, _ := strconv.Atoi(strings.Replace(shortUrl, "http://tinyurl.com/", "", 1))
	return this.hash[num]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
