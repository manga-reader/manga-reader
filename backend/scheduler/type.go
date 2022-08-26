package scheduler

import "github.com/manga-reader/manga-reader/backend/usecases"

type Scheduler struct {
	usecase *usecases.Usecase
}

func NewScheduler(usecase *usecases.Usecase) *Scheduler {
	return &Scheduler{
		usecase: usecase,
	}
}
