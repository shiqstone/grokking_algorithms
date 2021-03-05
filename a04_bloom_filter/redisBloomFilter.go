package main

import "github.com/garyburd/redigo/redis"

type HashFunc func(string) uint64

func NewHashFunc() []HashFunc {
	hs := make([]HashFunc, 0)
	var hfunc HashFunc
	hfunc = BKDRHash
	hs = append(hs, hfunc)
	hfunc = SDBMHash
	hs = append(hs, hfunc)
	hfunc = DJBHash
	hs = append(hs, hfunc)
	return hs
}

func BKDRHash(str string) uint64 {
    seed := uint64(131) // 31 131 1313 13131 131313 etc..
    hash := uint64(0)
    for i := 0; i < len(str); i++ {
        hash = (hash * seed) + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}
func SDBMHash(str string) uint64{
    hash := uint64(0)
    for i := 0; i < len(str); i++ {
        hash = uint64(str[i]) + (hash << 6) + (hash << 16) - hash
    }
    return hash & 0x7FFFFFFF
}
func DJBHash(str string) uint64{
    hash := uint64(0)
    for i := 0; i < len(str); i++ {
        hash = ((hash << 5) + hash) + uint64(str[i])
    }
    return hash & 0x7FFFFFFF
}

type RBloomFilter struct {
	Conn redis.Conn
	Key string
	Funcs []HashFunc
}

func NewRBloomFilter(conn redis.Conn) *RBloomFilter {
	return &RBloomFilter{Conn: conn, Key: "bloom_filter", Funcs: NewHashFunc()}
}

func (b *RBloomFilter)Add(str string) error{
	var err error
	for _,f := range b.Funcs {
		offset := f(str)
	   _,err := b.Conn.Do("setbit", b.Key, offset,1)
	   if err != nil {
		   return err
	   }
   }
	return err
}
func (b *RBloomFilter)Exist(str string) bool{
	var a int64=1
	for _,f := range b.Funcs {
		offset := f(str)
	   bitValue,_ := b.Conn.Do("getbit", b.Key, offset)
	   if bitValue != a {
		   return false
	   }
   }
   return true
}
