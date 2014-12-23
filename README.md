# The Rat Programming Language #

**NOTE** This project is not workable at all right now. No clear schedule is made as I don't have any dedicated time for it. I'll commit and push *unstable* code everyday to save my work progress. Please don't use it until I remove this line in README and announce first release version. Thanks for your understanding.

`Rat` is a general purpose language on top of `go`. In short, `rat` is a superset of `go` plus generic and generator.

Full language specification can be viewed in [`language-design` repository](https://github.com/go-rat/language-design) or [gitbook](http://huandu.gitbooks.io/rat/).

## Install `rat` compiler ##

`Rat` is a program written in `go`. [Install `go` command line](http://golang.org/doc/install) if you don't have one.

Install `rat` with following command.

	go install github.com/go-rat/rat

Make sure `rat` is available in your command line. Type `rat help` to see what happens. If `rat` command is not found, you need to add `$(go env GOROOT)/bin` to system `PATH`.

## License ##

This project is licensed under the MIT license that can be found in the LICENSE file.
