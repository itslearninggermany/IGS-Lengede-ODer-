package userprovisioning

import (
		"encoding/xml"
		"fmt"
)

type person struct {
	XMLName   xml.Name `xml:"person"`
	Recstatus string   `xml:"recstatus,attr"`
	Sourcedid struct {
		Source string `xml:"source"`
		ID     string `xml:"id"`
	} `xml:"sourcedid"`
	Userid string `xml:"userid"`
	Name   struct {
		Fn   string `xml:"fn"`
		N    struct {
			Family string `xml:"family"`
			Given  string `xml:"given"`
		} `xml:"n"`
	} `xml:"name"`
	Demographics struct {
		Bday string `xml:"bday"`
	} `xml:"demographics"`
	Email string `xml:"email"`
	Tel   [] telefon `xml:"tel"`
	Adr struct {
		Street   string `xml:"street"`
		Locality string `xml:"locality"`
		Pcode    string `xml:"pcode"`
	} `xml:"adr"`
	Institutionrole struct {
		Primaryrole         string `xml:"primaryrole,attr"`
		Institutionroletype string `xml:"institutionroletype,attr"`
	} `xml:"institutionrole"`
	Extension struct {
		Relationship [] child `xml:"relationship"`
	} `xml:"extension"`
}
// Creates a new Person
func Person (institution, personID, userID, givenname, familyname, birthday, email, street, locality, pcode, kindOfPersonStudentOrStaffOrParent string, tel []telefon, relation [] child ) (output *person) {
	a := new (person)
	a.Recstatus = "1"
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = personID
	a.Userid = userID
	a.Name.Fn = givenname + familyname
	a.Name.N.Family = familyname
	a.Name.N.Given = givenname
	a.Email = email
	a.Demographics.Bday = birthday
	a.Tel = tel
	a.Adr.Street = street
	a.Adr.Locality = locality
	a.Adr.Pcode = pcode
	a.Institutionrole.Primaryrole = "Yes"
	if kindOfPersonStudentOrStaffOrParent == "Staff" {
		a.Institutionrole.Institutionroletype = "Staff"
	} else 	if kindOfPersonStudentOrStaffOrParent == "Student" {
		a.Institutionrole.Institutionroletype = "Student"
	} else if kindOfPersonStudentOrStaffOrParent == "Parent" {
		a.Institutionrole.Institutionroletype = "Other"
	} else {
		fmt.Println("The Person is not a Staff, Student or Parent")
	}
	a.Extension.Relationship = relation
	return a
}

func (a *person) ParseToXML () []byte{
	x,_ := xml.Marshal(a)
	return x
}



