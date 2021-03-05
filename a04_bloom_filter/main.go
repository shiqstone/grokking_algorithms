package main

import "fmt"

func main() {
	filter := NewBloomFilter()
	fmt.Println(filter.Funcs[1].Seed)
	str1 := "hello,bloom filter!"
	filter.Add(str1)
	str2 := "A happy day"
	filter.Add(str2)
	str3 := "Greate wall"
	filter.Add(str3)

	fmt.Println(filter.Set.Count())
	fmt.Println(filter.Contains(str1))
	fmt.Println(filter.Contains(str2))
	fmt.Println(filter.Contains(str3))
	fmt.Println(filter.Contains("freedom!"))
	
	//redis bloom filter
	redisConn := "127.0.0.1:6379"
	redisAuth := "auth"
	con,err := redis.Dial("tcp", redisConn)//redis connect
	defer con.Close()

	bloom := NewRBloomFilter(con) //create bloom_filter
	bloom.Conn.Do("auth", redisAuth)
	bloom.Add(str1)
	bloom.Add(str2)
	bloom.Add(str3)
	fmt.Println(bloom.Exist(str1))
	fmt.Println(bloom.Exist(str2))
	fmt.Println(bloom.Exist(str3))
	fmt.Println(bloom.Exist("freedom!"))
}
