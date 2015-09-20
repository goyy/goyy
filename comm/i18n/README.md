# comm-i18n [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/comm/i18n)
i18n library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/comm/i18n`

# Usage
	en_US := map[string]string{
		"say":  "Hello, world!",
		"sayf": "Hello, %s!",
	}
	locales := map[string]map[string]string{
		i18n.Locale_en_US: en_US,
	}
	i18N := i18n.New(locales, i18n.Locale_en_US)
	fmt.Println(i18N.Message("say"))           // Output: Hello, world!
	fmt.Println(i18N.Messagef("sayf", "goyy")) // Output: Hello, goyy!
