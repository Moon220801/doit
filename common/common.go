package common

import (
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type JsonTime time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

//實現json反序列化，從傳遞的字串中解析成時間物件
func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = JsonTime(now)
	return
}

//實現json序列化，將時間轉換成字串byte陣列
func (t JsonTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

//mongodb是儲存bson格式，因此需要實現序列化bsonvalue(這裡不能實現MarshalBSON，MarshalBSON是處理Document的)，將時間轉換成mongodb能識別的primitive.DateTime
func (t *JsonTime) MarshalBSONValue() (bsontype.Type, []byte, error) {
	targetTime := primitive.NewDateTimeFromTime(time.Time(*t))
	return bson.MarshalValue(targetTime)
}

//實現bson反序列化，從mongodb中讀取資料轉換成time.Time格式，這裡用到了bsoncore中的方法讀取資料轉換成datetime然後再轉換成time.Time
func (t *JsonTime) UnmarshalBSONValue(t2 bsontype.Type, data []byte) error {
	v, _, valid := bsoncore.ReadValue(data, t2)
	if valid == false {
		return errors.New(fmt.Sprintf("%s, %s, %s", "讀取資料失敗:", t2, data))
	}
	*t = JsonTime(v.Time())
	return nil
}
