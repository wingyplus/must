# must

A convenient library to do a `must` pattern

## Problems

Before Go 1.18, if you want to panic when the regular expression cannot compile, you need to
do it like this:

```go
package main

import "regexp"

var re = must(regexp.Compile("<invalid_regexp>"))

func must(re *regexp.Regexp, err error) {
	// handle an error and if panic when error is not nil.
}
```

Ok. The Go standard library provides a convenient way to do this style by provides `regexp.MustCompile`
for us. They do the same thing for package like `template` also. But why do we needs to write
this kind of pattern for each type?

## Go 1.18 save our life

Now Go 1.18 lands with generic feature. Now we can change it to this!:

```go
package main

import (
	"regexp"

	"github.com/wingyplus/must"
)

var re = must.Must(regexp.Compile("<regex>")) // The `Must` is generic!
var tmpl = must.Must(template.New("name").Parse("text")) // Yeah, use the same logic.
```

## This maybe violate the Go Proverb

If you thing this violate the "[A little copying is better than a little dependency.](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=9m28s)" of the Go Proverbs rule. The just copy the code, I don't mind if you do it. But I hope this will reduce the LoC of your code base, less or more
