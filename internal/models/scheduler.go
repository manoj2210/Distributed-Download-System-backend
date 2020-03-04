package models

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
	var R []Record
	R=append(R,Record{"none",0,false})
	return &Scheduler{
		Url:		 url,
		GroupID:     grpID,
		TotalChunks: 0,
		Data:        R,
		Ptr:0,
	}
}

func NewRecord(uID string,f int64)*Record{
	return &Record{
		UserID:       uID,
		FileNo:       f,
		Acknowledged: false,
	}
}
