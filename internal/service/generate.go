package service

import "math/rand"

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_" // 26+26+10+1=63
const letterBits = 6                                                               // 2^6 = 64 > 63
const shortUrlLenght = 10                                                          // 10 * 6 = 60 > 63
const letterMask = (int64(1) << letterBits) - 1                                    // 1111111 - 1 = 0111111

func generate() string {
	var shortUrl = make([]byte, shortUrlLenght)

	var r int64 = rand.Int63() // each 6 bits represent letter from alphabet
	var l int                  // alphabet letter index

	for i := 0; i < len(shortUrl); i++ {
		l = int(r & letterMask)
		if l >= len(alphabet) {
			l = 0 // higher chanse of 'a' occuring
		}
		shortUrl[i] = alphabet[l]
		r = r >> letterBits // cut processed bits
	}

	return string(shortUrl)
}
