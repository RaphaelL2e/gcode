package millcode

import "fmt"

func PrintN(n int) {
	for i := 1; i <= n; i++ {
		fmt.Print(n)
	}

}

func PrintN2(n int) {
	if n > 0 {
		PrintN2(n - 1)
		fmt.Print(n)
	}
}
