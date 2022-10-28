package main

import (
	"fmt"
	"time"
)

func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

func main() {

	foo := map[string]interface{}{
		"Nama":           "Ramdoni",
		"Umur":           28,
		"Tempat Tinggal": "Jauh",
	}
	timeMap(foo)
	fmt.Println(foo)

}
