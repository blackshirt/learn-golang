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

type Category struct {
	Id           int    `json:"id"`
	Code         string `json:"code"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
}

type Penyelenggara struct {
	Id        int    `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
	Alamat    string `json:"alamat"`
	Kota      string `json:"kota"`
}

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

type TrainRekap struct {
	Diklat        string `json:"diklat"`
	Penyelenggara string `json:"penyelenggara"`
	Peserta       int    `json:"peserta"`
	Tahun         string `json:"year"`
}

type TrainRepo interface {
	GetById(int) (*Train, error)
	Fetch(int, int) ([]*Train, error)
	Posts(*Train) error
	RekapTrain() ([]*TrainRekap, error)
}
