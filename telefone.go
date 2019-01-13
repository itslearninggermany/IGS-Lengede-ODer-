package userprovisioning

import ("strconv")

type telefon struct {
	Text    string `xml:",chardata"`
	Teltype string `xml:"teltype,attr"`
}

func Telefon (OneForMobileTwoForOthers int, number string) *telefon {
	a := new (telefon)
	a.Teltype = strconv.Itoa(OneForMobileTwoForOthers)
	a.Text = number
	return a
}

func (a *telefon) GetTelefon () (OneForMobileTwoForOthers int, number string) {
	tmp, _ := strconv.Atoi(a.Teltype)
	return tmp, a.Text
}

func MakeATelefonSlice (oneForMobileTwoForOthers []int, number [] string) (res []telefon) {
	for i :=0; i < len(number); i++ {
		a := Telefon(oneForMobileTwoForOthers[i],number[i])
		res = append(res,*a)
	}
	return
}
