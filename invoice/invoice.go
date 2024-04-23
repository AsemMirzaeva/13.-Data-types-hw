
package invoice

import (
	"fmt"
	"strconv"
)

func CreateAndSave(ism string, kun uint, narx float32) {
	invoice := "Invoice\n"
	invoice += "Ismi: " + ism + "\n"
	invoice += "Necha kun: " + strconv.Itoa(int(kun)) + "\n"
	invoice += "Narxi (kun): $" + strconv.FormatFloat(float64(narx), 'f', 2, 32) + "\n"
	invoice += "Umumiy: $" + strconv.FormatFloat(float64(kun)*float64(narx), 'f', 2, 32) + "\n"

	fmt.Println(invoice)
}
