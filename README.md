```go
package main

import (
	"github.com/z-song/golisp"
)

func main() {

	code := `

;(print (type (lambda (x) (+ x x))))

;(print (int (double "123.112312")))

;(print (type (string 123)))

;(print ((lambda (x) (+ x x)) 43))

(define gt (lambda (x) (<= x 5)))

;(print (car (list 12 "hi")))

;(print (cdr (list 1 2 3 4 5 6)))

;(print (filter gt [1 5 6 8 3 10]))

;(print (fill [] 100 "hello"))

(print (append ["hello"] 123 "world" [1 2 3 4 5]))

;(print (!= "hello" (- 100 99)))

;(print (map double [1 2 4 5]))

;(define double (lambda (x) (+ x x)))

;(print (map double '(1 2 3 4)))

;(print (append "hello" " world"))

;(print (split "hello world" "o"))


	   `

	golisp.EvalString(code)

}

```