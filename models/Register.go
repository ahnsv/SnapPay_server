package models

type Register struct {
	URL         string `json:"url"`
	ProductInfo struct {
		ID        int      `json:"id"`
		Title     string   `json:"title"`
		Subtitle  string   `json:"subtitle"`
		Inventory string   `json:"inventory"`
		Options   []string `json:"options"`
		Price     string   `json:"price"`
		Image     []string `json:"image"`
	} `json:"product_info"`
	RegisterInfo struct {
		Date     string `json:"date"`
		Userinfo struct {
			UserName     string `json:"user_name"`
			UserUsername string `json:"user_username"`
			UserEmail    string `json:"user_email"`
		} `json:"userinfo"`
		RegisterID string `json:"register_id"`
	} `json:"register_info"`
}
