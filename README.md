# dig

Package `dig` provides tools for traversing Go maps.

`dig` turns this:

```go
foo = myMap["a"].(map[string]interface{})["b"].(map[string]interface{})["c"].(uint64)
```

into this:

```go
foo = dig.Uint64(myMap, "a", "b", "c")
```

Useful when you depend on external services, such as web-APIs, and you don't
really want to handle all possible error cases or possible panic points caused
by sudden changes on the expected JSON.

## Installation

```
go get -u github.com/xiam/dig
```

## Usage

Start with a `map[string]interface{}` and define which path to follow, take
this `testMap` example:

```go
testMap := map[string]interface{}{
	"first": map[string]interface{}{
		"first.1": map[string]string{
			"col.1": "a",
			"col.2": "b",
			"col.3": "c",
		},
	},
	"second": map[string]map[string]string{
		"second.2": map[string]string{
			"col.4": "d",
			"col.5": "e",
			"col.6": "f",
		},
	},
}
```

suppose you want to get the value of the element with key "col.4". You'll need
a path, the path of a node is defined as the array of map keys that you need to
traverse in order to get to that node, in the above example the path would be:
`second` > `second.2` > `col.4`, in that specific order.

Now suppose you expect "col.4" to be a string, then you'd use `dig.String()` to
get the value.

```go
// s = "d"
s := dig.String(&testMap, "second", "second.2", "col.4")
```

Note that you're passing a pointer to map (`&testMap`) instead of a map.

There are more sugar methods with similar functionality, such as
`dig.Uint64()`, `dig.Float64()`, `dig.Interface()`, etc. all these methods will
attempt to convert the node data type into the expected Go type.

The `dig.String()` function and friends are just wrappers around `dig.Get()`,
and you can also use `dig.Get()` to get a node value, it will also produce an
`error` value if something goes wrong.

```go
var s string
err := dig.Get(&testMap, &s, "second", "second.2", "col.4")
```

Now suppose you'd like to set the value of "col.5", you'd use the `dig.Set()`
function:

```go
dig.Set(&testMap, "modified", "second", "second.2", "col.4")
```

And finally, if you're ever in the need of creating a route on a map, you could
use `dig.Dig()` like this:

```go
var i int

m := map[string]interface{}{}

// Create a path.
dig.Dig(&m, "foo", "bar", "baz")

// Set the value 42.
dig.Set(&m, 42, "foo", "bar", "baz")

// Get the value, expecting 42.
dig.Get(&m, &i, "foo", "bar", "baz")

if i != 42 {
	t.Errorf("Test failed.")
}
```

Please consult the [online documentation][1] to get a list of all available
methods.

## Documentation

See the [online docs][1].

## License

>  Copyright (c) 2013-today José Nieto, https://xiam.dev
>
>  Permission is hereby granted, free of charge, to any person obtaining
>  a copy of this software and associated documentation files (the
>  "Software"), to deal in the Software without restriction, including
>  without limitation the rights to use, copy, modify, merge, publish,
>  distribute, sublicense, and/or sell copies of the Software, and to
>  permit persons to whom the Software is furnished to do so, subject to
>  the following conditions:
>
>  The above copyright notice and this permission notice shall be
>  included in all copies or substantial portions of the Software.
>
>  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
>  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
>  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
>  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
>  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
>  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
>  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

[1]: http://godoc.org/github.com/xiam/dig
