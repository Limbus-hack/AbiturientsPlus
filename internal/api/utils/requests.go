package utils

import "fmt"

type UserRequest struct {
	Token    string
	Endpoint string
	Count    int
	City     int
	Sort     int
	AgeFrom  int
	AgeTo    int
	Fields   string
}

func (u UserRequest) String(q rune) string {
	return fmt.Sprintf("%s%d&count=%d&access_token=%s&city=%d&v=5.21&age_from=%d&age_to=%d&fields=%s&q=%s",
		u.Endpoint,
		u.Sort,
		u.Count,
		u.Token,
		u.City,
		u.AgeFrom,
		u.AgeTo,
		u.Fields,
		string(q))
}
