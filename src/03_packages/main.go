package main

import (
	"fmt"
	"math"
	"github.com/kpatryk/go_crash_course/src/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
}
