package model

type Item struct {
	FirstName string `json:"first_name"`
	Id        int64  `json:"id"`
	LastName  string `json:"last_name"`
	TrackCode string `json:"track_code"`
	Sex       int    `json:"sex"`
	Interests string `json:"interests"`
}

type Response struct {
	Count int64  `json:"count"`
	Item  []Item `json:"items"`
}

type VkUsers struct {
	Response Response `json:"response"`
}
