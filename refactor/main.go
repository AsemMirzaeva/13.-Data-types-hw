// Sizning hamkasblaringizdan biri noyob main.go faylidan iborat dastur yaratdi
// Kodning barqarorligini yaxshilash uchun sizdan kodni qayta tahrirlashingiz so'raladi (refactor).
// Yangi kod tashkilotini taklif qiling:

// Qaysi paketlarni yaratishingiz kerak?

// Yangi katalog yaratish kerakmi?

// main.go

package main

import (
	"fmt"
	"reservation/email"
	"reservation/invoice"
)

func main() {
	customerName := "Zuhriddin"
	customerEmail := "zuhriddin@gmail.com"
	var nights uint = 12
	emailContents := email.GetContents("Abdug'aniyev", customerName, nights)
	txt := email.Send(emailContents, customerEmail)
	fmt.Println(txt)
	fmt.Println("")
	invoice.CreateAndSave(customerName, nights, 145.32)
}
