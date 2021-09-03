// package diffie_hellman_key_exchange
// implements Diffie Hellman Key Exchange Algorithm
// ref: https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
// watch: https://www.youtube.com/watch?v=NmM9HA2MQGI

package diffie_hellman_key_exchange

const (
	generator       = 3
	primeNum  int64 = 6700417 // prime discovered by Euler
)

// GenerateSharedKey : generates a key using private client key, generator & primeNum
// can be made public
// shareKey = (g^key) % primeNum
func GenerateSharedKey(privKey int64) int64 {
	return modularExponentiation(generator, privKey, primeNum)
}

// GenerateMutualKey : generates a mutual key that can only be used by "alice" & "bob"
// mutualKey = (shareKey^privKey) % primeNum
func GenerateMutualKey(privKey, shareKey int64) int64 {
	return modularExponentiation(shareKey, privKey, primeNum)

}

// r = (b^e) % mod
func modularExponentiation(b, e, mod int64) int64 {
	// runs in O(log(n)) where n = e
	// uses expontentiation by squaring to speed up things
	if mod == 1 {
		return 0
	}
	var r int64 = 1
	b = b % mod
	for e > 0 {
		if e&1 == 1 {
			r = (r * b) % mod
		}
		e = e >> 1
		b = (b * b) % mod
	}
	return r
}
