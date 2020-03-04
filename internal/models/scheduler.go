package models

//var SchedulerArray=make(map[string]*Scheduler)

type Scheduler struct {
	Url string `json:"url" bson:"url"`
	GroupID string `json:"groupID" bson:"groupID"`
	TotalChunks int64 `json:"totalFiles" bson:"totalFiles"`
	Data []Record `json:"data" bson:"data"`
	Ptr int64 `json:"ptr" bson:"ptr"`
}

type Record struct{
	UserID string `json:"userID"" bson:"userID"`
	FileNo int64 `json:"fileNo" bson:"fileNo"`
	Acknowledged bool `json:"ack" bson:"ack"'`
}

func NewScheduler(url,grpID string)*Scheduler{
	return &Scheduler{
		Url:		 url,
		GroupID:     grpID,
		TotalChunks: 0,
		Data:        nil,
		Ptr:0,
	}
}

func NewRecord(uID string,f int)*Record{
	return &Record{
		UserID:       uID,
		FileNo:       int64(f),
		Acknowledged: false,
	}
}

//func (s *Scheduler)Allocate(uID string) int{
//	if s.lock<int(s.TotalChunks) {
//		s.lock += 1
//		s.Record[s.lock] = uID
//		return s.lock
//	}
//	return -1
//}
//
//func NewScheduler(grpID string,t int) {
//	SchedulerArray[grpID]=&Scheduler{
//		TotalChunks: int64(t),
//		lock: -1,
//		Record: make(map[int]string),
//	}
//}
//
//func DisplayArray(grpID string)(*Scheduler,error){
//	k,ok:=SchedulerArray[grpID]
//	if ok{
//		return k,nil
//	}
//	return nil,errors.New("No such Data Available")
//
//}
