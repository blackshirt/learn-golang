///////////////////////
///// usecase.go //////
///////////////////////

/*
This layer contains application specific business rules. For sometimes its called
engines. The application level rules and use-cases orchestrate the domain model
and add richer rules and logic including persistence.
*/

package cleanarch

type trainUCase struct {
	trainRepo TrainRepo
}

type TrainUCase interface {
	GetById(int) (*Train, error)
	Fetch(int, int) ([]*Train, error)
}

// Get a single result
func (tuc *trainUCase) GetById(id int) (*Train, error) {
	res, err := tuc.trainRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// For multi results set
func (tuc *trainUCase) Fetch(start, limit int) ([]*Train, error) {
    res, err := tuc.trainRepo(start, limit)
    if err != nil {
        return nil, err
    }
    return res, nil
}

// Constructor
func NewTrainUCase(trep TrainRepo) *trainUCase {
	return &trainUCase{trainRepo: trep}
}
