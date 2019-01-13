package userprovisioning

import (
	"os"
	"fmt"
	"encoding/csv"
	"strings"
)

type igsLengedeData struct {
	content [] data
}

// This is the structure of the file coming from Atlantis SIS
type data struct {
	KindOfMemeber string `csv:"gruppen_kz"`
	SchoolID      string `csv:"Schulen-ID"`
	Schoolyear    string `csv:"Schuljahr"`
	UserID        string `csv:"Eindeutige ID"`
	FamilyID      string `csv:"Familien-Nr."`
	Searchstring  string `csv:"Suchbegriff"`
	Birthday      string `csv:"Geb.-Datum"`
	ChildID       string `csv:"Stamm-Kürzel"`
	Sex           string `csv:"Ge."`
	Telefon1      string `csv:"Telefon 1"`
	Telefon2      string `csv:"Telefon 2"`
	FullName      string `csv:"Display-Name"`
	Username      string `csv:"Username"`
	Email         string `csv:"E-Mail"`
	Password      string `csv:"Account-Passwort"`
	Groups        string `csv:"Groups"`
}


func CsvData (path string, readFirstLine bool) (*igsLengedeData) {
	output := new(igsLengedeData)
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer csvFile.Close()
	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		panic(err)
	}
	i := 0
	for _, line := range lines {
		tmp := data{
			KindOfMemeber: line[0],
			SchoolID:      line[1],
			Schoolyear:    line[2],
			UserID:        line[3],
			FamilyID:      line[4],
			Searchstring:  line[5],
			Birthday:      line[6],
			ChildID:       line[7],
			Sex:           line[8],
			Telefon1:      line[9],
			Telefon2:      line[10],
			FullName:      line[11],
			Username:      line[12],
			Email:         line[13],
			Password:      line[14],
			Groups:        line[15],
		}

		if readFirstLine {
			output.content = append(output.content, tmp)
		} else if i > 0 {
			output.content = append(output.content, tmp)
		}
		i++
	}
	return output
}


func (a *igsLengedeData) CreateIgsLengedePerson (institution string) (output []byte) {
	for i := 0; i < len(a.content); i++ {
		var member string
		var kids [] string

		if a.content[i].KindOfMemeber == "E" {
			member = "Parent"
			kids = strings.Split(a.content[i].ChildID, "/")
		}
		if a.content[i].KindOfMemeber == "S" {
			member = "Student"
		}
		if a.content[i].KindOfMemeber == "M" {
			member = "Staff"
		}
		name := strings.Split(a.content[i].FullName, " ")
		fmt.Println(name)
		tmp := Person(institution, a.content[i].UserID, a.content[i].UserID, name[0], name[1], a.content[i].Birthday, a.content[i].Email, "", "", "", member, MakeATelefonSlice([] int{1,2},	[] string {a.content[i].Telefon1, a.content[i].Telefon2}), MakeAChildSlice(institution, kids))

		tmpByte := tmp.ParseToXML()
		for in := 0; i < len(tmpByte); i++ {
			output = append(output,tmpByte[in])
		}
	}
	return
}