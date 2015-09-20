# data-cache [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/data/cache)
cache library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/data/cache`

# Usage
	cache.Init(cache.Conf{Address: "10.105.99.81:6379"})

	cache.Set("key-a", "value-a")
	v, _ := cache.Get("key-a")
	fmt.Println(v)
