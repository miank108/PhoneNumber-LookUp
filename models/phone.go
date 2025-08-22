package models

type PhoneNumber struct {
	Number  string
	Name    string
	Country string
	Type    string
}

func (p *PhoneNumber) String() string {
	if p == nil {
		return ""
	}
	return "Found: " + p.Name + ", Country: " + p.Country + ", Type: " + p.Type
}
