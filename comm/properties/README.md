# comm-properties [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/comm/properties)
properties library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/comm/properties`

# Usage
	// example.properties:
	// say=Hello, world!
	// sayf=Hello, %s!
	p, _ := properties.New("./example.properties")
	fmt.Println(p.Property("say"))           // Output: Hello, world!
	fmt.Println(p.Propertyf("sayf", "goyy")) // Output: Hello, goyy!
