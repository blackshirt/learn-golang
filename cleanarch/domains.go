/////////////////////
//// domains.go /////
/////////////////////

/*
Its domains of the business in the clean architecture pattern and the core parts
of the applications.Its contains business models, rules and no external dependencies.
Sometimes its called with entities or models part, represent your businnes objects
of your application.
*/

package cleanarch

type Train struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Begin         string `json:"begin"`
	End           string `json:"end"`
	Location      string `json:"location"`
	Address       string `json:"address"`
	City          string `json:"city"`
	Jp            int    `json:"jp,string"`
	Category      int    `json:"category"`
	Penyelenggara int    `json:"penyelenggara"`
}

type TrainRepo interface {
	GetById(int) (*Train, error)
	Fetch(int, int) ([]*Train, error)
	Posts(*Train) error
}
