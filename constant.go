package main

import "fmt"

const Pi = 3.14

/**
 * deklarasi global constant
 * yang dipanggil di dalam fungsi
 */

const (
	StatusOk                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNonAutoritativeInfo = 203
	StatusNoContent           = 205
	StatusPartialContent      = 206
)

func main() {

	action := func() {
		fmt.Println("Status Ok : ", StatusOk)
		fmt.Println("Status Created : ", StatusCreated)
		fmt.Println("Status Accepted : ", StatusAccepted)
		fmt.Println("Status Non Autoritative Info : ", StatusNonAutoritativeInfo)
		fmt.Println("Status Non Content : ", StatusNoContent)
		fmt.Println("Status Partial Content : ", StatusPartialContent)
	}

	action()
}
