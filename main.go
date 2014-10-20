package main

import (
	golisp "./golisp"
)

func main() {

	code := `

(define double (lambda (x) (+ x x)))

(print  (map double '(1 2 3 4 56)))

`
	golisp.Excute(code)

}
