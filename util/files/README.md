# util-files [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/util/files)
file library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/util/files`

# Usage
*IsExist*

	fmt.Println(files.IsExist("./example.txt"))
	fmt.Println(files.IsExist("./README"))
	// Output: 
	// true
	// false

*Read*

	s, _ := files.Read("./example.txt")
	fmt.Println(s)
	// Output: Hello world!

*Write*

	filename := "./example.txt"
	data := "Hello goyy!"
	if err := files.Write(filename, data, 0644); err != nil {
		log.Fatalf("Write %s: %v", filename, err)
	}
	s, _ := files.Read(filename)
	fmt.Println(s)
	files.Write(filename, "Hello world!", 0644) // recover
	// Output: Hello goyy!
