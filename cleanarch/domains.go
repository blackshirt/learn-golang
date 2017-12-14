////////////////////
//// domains.go ////
////////////////////

/*
Its domains of the business in the clean architecture pattern and the core parts
of the applications.Its contains business models, rules and no external dependencies.
Sometimes its called with entities or models part, represent your businnes objects
of your application.
*/

package cleanarch

type Train struct {
	Id   int    `json:"id"`
	Nama string `json:"nama_diklat`
}

type TrainRepo interface {
	GetById(int) (*Train, error)
}
