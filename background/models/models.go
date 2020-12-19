package models

type VkIDModel struct {
	Response struct {
		Count int   `json:"count"`
		Items []int `json:"items"`
	} `json:"response"`
}

type VkUserModel struct {
	Response []struct {
		FirstName       string `json:"first_name"`
		ID              int    `json:"id"`
		LastName        string `json:"last_name"`
		CanAccessClosed bool   `json:"can_access_closed"`
		IsClosed        bool   `json:"is_closed"`
		Sex             int    `json:"sex"`
		Verified        int    `json:"verified"`
		Interests       string `json:"interests"`
		Deactivated     string `json:"deactivated"`
		City            struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"city"`
	} `json:"response"`
}

type VkGroupModel struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			Name string `json:"name"`
		} `json:"items"`
	} `json:"response"`
}
