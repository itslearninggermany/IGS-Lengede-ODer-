package userprovisioning

type child struct {
	Sourcedid struct {
	Source string `xml:"source"`
	ID     string `xml:"id"`
		} `xml:"sourcedid"`
	Label string `xml:"label"`
	}


func Child (institution, childID string) child {
	a := new (child)
	a.Sourcedid.Source = institution
	a.Sourcedid.ID = childID
	a.Label = "child"
	return *a
}


func (a *child) GetChild () (institution, childID string) {
	return a.Sourcedid.Source, a.Sourcedid.ID
}

func MakeAChildSlice (institution string, childIDs [] string) (res []child) {
		for i :=0; i < len(childIDs); i++ {
			a := Child(institution,childIDs[i])
			res = append(res, a)
		}
		return
}
