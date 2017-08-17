# GoT - Go Templating

GoT is a very simple tool that uses Go's text templating
http://golang.org/pkg/text/template/ to parse a file and output a new file.

It is intended to be used with `go generate` for generating multiple sets of Go
source code files from one template. It isn't the greatest solution to the "No
Generics" problem of Go, but it's pretty simple and straightforward and has
served me well enough for years now.

I do like approaches others have taken to solve this problem -- especially
options like https://github.com/cheekybits/genny -- but I feel like _someday_
Go will have its own solution for the issue and until then, templates are fine
for me.

Since it's just simple text template parsing, you could use GoT for other
things too, though I've really just used it for source code. One example where
I've used it is https://github.com/gholt/holdme and it's just called from
within the package.go file when you use "go generate".

The syntax is pretty simple too:

```
Syntax: got <infile> <outfile> [key[=value]] [key,item1[,item2]...] ...
```

The infile and outfile should be self-explanatory.

The `key[=value]` args simply sets a template key to the value; if you omit the
value, it'll set it as the text `"true"`.

The `key,item1[,item2]...` args will set a template key to a slice containing
the items given. This is useful with the `{{range $v := .key}}` template
syntax. I've used this in places like
https://github.com/getcfs/megacfs/blob/master/oort/package.go#L18 where I
needed to generate a lot of boilerplate code around a set of method names.

I expect this tool to remain for the forseeable future, and to not break
compatibility either. It's just a single 58 line source code file, so shouldn't
be hard to maintain. I doubt the Go authors are going to fundamentally break
how Go templating works either.

> Copyright See AUTHORS. All rights reserved.  
> Use of this source code is governed by a BSD-style  
> license that can be found in the LICENSE file.
