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

type SubsRequest struct {
	Token    string
	Endpoint string
	Extended int
	UserId   int64
	Count    int
}

func (s SubsRequest) String() string {
	return fmt.Sprintf("%s%d&user_id=%d&access_token=%s&count=%d&v=5.21",
		s.Endpoint,
		s.Extended,
		s.UserId,
		s.Token,
		s.Count)
}
