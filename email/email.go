package email

import "fmt"

func Send(a string, to string) string {
	txt := ""
	txt = "email:" + to + "\n" + a
	return txt
}

func GetContents(t string, ism string, kun uint) string {
	txt := "Qadrli %s %s,\nSiz xonangizni %d kunga bron qilganingiz tasdiqladik. Kuningiz xayrli o'tsin !"
	return fmt.Sprintf(txt,
		t,
		ism,
		kun)
}
