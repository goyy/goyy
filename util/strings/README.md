# util-strings [![GoDoc](http://godoc.org/gopkg.in/goyy/goyy.v0?status.png)](http://godoc.org/gopkg.in/goyy/goyy.v0/util/strings)
string library for Go

# Installation
`go get gopkg.in/goyy/goyy.v0/util/strings`

# Usage
*IsBlank*

	fmt.Println(strings.IsBlank(""))
	fmt.Println(strings.IsBlank("  "))
	fmt.Println(strings.IsBlank(" \t\r\n "))
	fmt.Println(strings.IsBlank("a"))
	fmt.Println(strings.IsBlank(" a  "))
	fmt.Println(strings.IsBlank(" \t\r\n a \t\r\n "))
	// Output:
	// true
	// true
	// true
	// false
	// false
	// false

*IsNotBlank*

	fmt.Println(strings.IsNotBlank(""))
	fmt.Println(strings.IsNotBlank("  "))
	fmt.Println(strings.IsNotBlank(" \t\r\n "))
	fmt.Println(strings.IsNotBlank("a"))
	fmt.Println(strings.IsNotBlank(" a  "))
	fmt.Println(strings.IsNotBlank(" \t\r\n a \t\r\n "))
	// Output:
	// false
	// false
	// false
	// true
	// true
	// true

*Left*

	fmt.Println(strings.Left("abc", 2))
	fmt.Println(strings.Left("abc", 4))
	// Output:
	// ab
	// abc

*Right*

	fmt.Println(strings.Right("abc", 2))
	fmt.Println(strings.Right("abc", 4))
	// Output:
	// bc
	// abc

*Mid*

	fmt.Println(strings.Mid("abc", 0, 2))
	fmt.Println(strings.Mid("abc", 0, 4))
	fmt.Println(strings.Mid("abc", -2, 2))
	// Output:
	// ab
	// abc
	// bc

*Before*

	fmt.Println(strings.Before("abc", "c"))
	fmt.Println(strings.Before("abcba", "b"))
	// Output:
	// ab
	// a

*After*

	fmt.Println(strings.After("abc", "a"))
	fmt.Println(strings.After("abcba", "b"))
	// Output:
	// bc
	// cba

*BeforeLast*

	fmt.Println(strings.BeforeLast("abc", "c"))
	fmt.Println(strings.BeforeLast("abcba", "b"))
	// Output:
	// ab
	// abc

*AfterLast*

	fmt.Println(strings.AfterLast("abc", "a"))
	fmt.Println(strings.AfterLast("abcba", "b"))
	// Output:
	// bc
	// a

*Between*

	fmt.Println(strings.Between("yabcz", "y", "z"))
	fmt.Println(strings.Between("yabczydefz", "y", "z"))
	// Output:
	// abc
	// abc

*BetweenSame*

	fmt.Println(strings.BetweenSame("tagabctag", "tag"))
	// Output: abc

*Repeat*

	fmt.Println(strings.Repeat("-", 10))
	// Output: ----------

*PadStart*

	fmt.Printf("%#q\n", strings.PadStart("bat", 5))
	// Output: `  bat`

*PadEnd*

	fmt.Printf("%#q\n", strings.PadEnd("bat", 5))
	// Output: `bat  `

*PadLeft*

	fmt.Println(strings.PadLeft("bat", 5, "*"))
	// Output: **bat

*PadRight*

	fmt.Println(strings.PadRight("bat", 5, "*"))
	// Output: bat**

*Pad*

	fmt.Printf("%#q\n", strings.Pad("bat", 5))
	// Output: ` bat `
