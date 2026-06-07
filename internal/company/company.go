package company

type SearchResult struct {
	Name           string `json:"company_name"`
	Number         string `json:"company_number"`
	Status         string `json:"company_status"`
	DateOfCreation string `json:"date_of_creation"`
}

type SearchResponse struct {
	Items []SearchResult `json:"items"`
}

type Company struct {
	Name             string `json:"company_name"`
	Number           string `json:"company_number"`
	Status           string `json:"company_status"`
	DateOfCreation   string `json:"date_of_creation"`
	AccountsOverdue  bool   `json:"accounts_overdue"`
	LastAccountsDate string `json:"last_accounts_date"`
}
