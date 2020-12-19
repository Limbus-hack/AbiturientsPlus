package utils

import "fmt"

type CityRequest struct {
	Token     string
	Count     int
	Endpoint  string
	Sort      int
	CountryId int
	NeedAll   int
}

type UserRequest struct {
	Token    string
	Endpoint string
	Count    int
	City     int
	Sort     int
	AgeFrom  int
	AgeTo    int
}

func (c CityRequest) String() string {
	return fmt.Sprintf("%s%d&count=%d&access_token=%s&country_id=%d&v=5.21&need_all=%d",
		c.Endpoint,
		c.Sort,
		c.Count,
		c.Token,
		c.CountryId,
		c.NeedAll,
	)
}

func (u UserRequest) String() string {
	return fmt.Sprintf("%s%d&count=%d&access_token=%s&city=%d&v=5.21&age_from=%d&age_to=%d",
		u.Endpoint,
		u.Sort,
		u.Count,
		u.Token,
		u.City,
		u.AgeFrom,
		u.AgeTo)
}
