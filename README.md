# malgo 🚧

**malgo** is a _work-in-progress_ implementation of **[mal](https://github.com/kanaka/mal)** Lisp interpreter in Go.

I'm using this as a learning tool to:
- better understand the components of Lisp (and programming languages in general)
- practice writing Go 🙂

I've been loosely following the [process guide](https://github.com/kanaka/mal/blob/master/process/guide.md) by the original author. The [mal](https://github.com/kanaka/mal) repo contains over 80 implementations, including one in Go, but I intend on _not_ looking at it until my implementation is more fleshed-out. Looking forward to compare and contrast them at a later point.

### Run
```sh
go get github.com/poly2d/malgo
cd $GOPATH/src/github.com/poly2d/malgo && go run mal.go

```

#### Special Forms
`def! do fn* if let*`

#### Core Functions
- Arithmetic: `+ - * /`
- Comparators: `= < <= > >=`
- List: `list list? empty? count`
- Print: `prn`

### Todos
- Continue guide (steps 5+ and deferrables)
- Handle panics
- Automate testing
  - e.g. Makefile, Github workflow
