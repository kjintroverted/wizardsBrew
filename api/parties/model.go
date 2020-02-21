package parties

// Party describes a campaign party
type Party struct {
	ID       string   `json:"id" db:"id"`
	Name     string   `json:"name" db:"name"`
	PhotoURL string   `json:"photoURL" db:"photo_url"`
	Admin    string   `json:"admin" db:"admin"`
	Members  []string `json:"members" db:"members"`
}
