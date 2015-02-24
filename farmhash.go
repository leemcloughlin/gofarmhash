// farmhash Googles hash algorithm

/*
Farmash is a successor to Cityhash

This code is just a conversion of Googles C++ into go and moved around a bit
to implement the hash interface
*/
package farmhash

const (
	Version = "1.1"
)

type Farmhash32 struct {
}

type Farmhash64 struct {
}

func (f *Farmhash32) Reset() {
}

// These functions based on original C++ namespace util

func Hash32(s []byte, len uint64) uint32 {
	return mkHash32(s, len)
}

func Hash32WithSeed(s []byte, len uint64, seed uint32) uint32 {
	return mkHash32WithSeed(s, len, seed)
}

func Hash64(s []byte, len uint64) uint64 {
	return naHash64(s, len)
}

func Hash64WithSeed(s []byte, len, seed uint64) uint64 {
	return naHash64WithSeed(s, len, seed)
}

func Hash64WithSeeds(s []byte, len, seed0, seed1 uint64) uint64 {
	return naHash64WithSeeds(s, len, seed0, seed1)
}

func Hash128(s []byte, len uint64) uint128 {
	return Fingerprint128(s, len)
}

func Hash128WithSeed(s []byte, len uint64, seed uint128) uint128 {
	return CityHash128WithSeed(s, len, seed)
}

func FingerPrint32(s []byte, len uint64) uint32 {
	return mkHash32(s, len)
}

func FingerPrint64(s []byte, len uint64) uint64 {
	return naHash64(s, len)
}

/*
uint128_t Fingerprint128(const char* s, size_t len) {
  return farmhashcc::Fingerprint128(s, len);
}
*/
