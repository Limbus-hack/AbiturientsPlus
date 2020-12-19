package model

type UserItem struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Sex       int    `json:"sex"`
	Interests string `json:"interests"`
}

type UserResponse struct {
	Count int64      `json:"count"`
	Items []UserItem `json:"items"`
}

type VkUsers struct {
	Response UserResponse `json:"response"`
}

type SubscriptionItem struct {
	Name string `json:"name"`
}

type SubscriptionsResponse struct {
	Count             int64              `json:"count"`
	SubscriptionItems []SubscriptionItem `json:"items"`
}

type VkUserSubscriptions struct {
	SubscriptionsResponse SubscriptionsResponse
}
