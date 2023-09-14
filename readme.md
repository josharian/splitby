I find [bufio.SplitFunc](https://pkg.go.dev/bufio#SplitFunc)s difficult to write.

This package provides a common class of bufio.SplitFuncs, those that split on a separator or using a regular expression.

(To avoid pulling in heavy dependencies, it does not actually import package regexp. It depends only on bufio, bytes, and errors.)
