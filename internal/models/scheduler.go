package models

import(
	"errors"
)

var SchedulerArray=make(map[string]*Scheduler)

type Scheduler struct {
	lock int
	TotalChunks int64
	Record map[int]string
}

func NewScheduler(grpID string,t int) {
	SchedulerArray[grpID]=&Scheduler{
		TotalChunks: int64(t),
		lock: -1,
		Record: make(map[int]string),
	}
}

func (s *Scheduler)Allocate(uID string) int{
	if s.lock<int(s.TotalChunks) {
		s.lock += 1
		s.Record[s.lock] = uID
		return s.lock
	}
	return -1
}

func DisplayArray(grpID string)(*Scheduler,error){
	k,ok:=SchedulerArray[grpID]
	if ok{
		return k,nil
	}
	return nil,errors.New("No such Data Available")

}
