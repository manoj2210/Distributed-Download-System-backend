package models

var SchedulerArray=make(map[string]*Scheduler)

type Scheduler struct {
	lock int
	TotalChunks int
	Record map[int]string
}

func NewScheduler(grpID string,t int) {
	SchedulerArray[grpID]=&Scheduler{
		TotalChunks: t,
		lock: -1,
	}
}

func (s *Scheduler)Allocate(uID string) int{
	if s.lock<s.TotalChunks {
		s.lock += 1
		s.Record[s.lock] = uID
		return s.lock
	}
	return -1
}

func DisplayArray(grpID string)*Scheduler{
	return SchedulerArray[grpID]
}
