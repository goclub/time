package xtime

import (
	json "encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestChinaTime(t *testing.T) {
	String := func(v interface{}) string {
		data, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		return string(data)
	}
	{
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.FixedZone("CST", 2*3600))
		assert.Equal(t, err, nil)
		assert.Equal(t, String(NewChinaTime(tValue)), `"2020-07-31 21:29:29"`)
	}
	{
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.FixedZone("CST", 8*3600))
		assert.Equal(t, err, nil)
		assert.Equal(t, String(NewChinaTime(tValue)), `"2020-07-31 15:29:29"`)
	}
	{
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.FixedZone("CST", 0*3600))
		assert.Equal(t, err, nil)
		assert.Equal(t, String(NewChinaTime(tValue)), `"2020-07-31 23:29:29"`)
	}
	{
		type Request struct {
			Time ChinaTime `db:"time"`
		}
		req := Request{}
		err := json.Unmarshal([]byte(`{"time":"2020-07-31 15:37:44"}`), &req)
		assert.NoError(t, err)
		assert.Equal(t, req.Time.In(time.FixedZone("CST", 8*3600)).String(), "2020-07-31 15:37:44 +0800 CST")
	}
	{
		type Reply struct {
			Time ChinaTime `db:"time"`
		}
		reply := Reply{}
		tValue, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-07-31 15:29:29", time.UTC)
		assert.Equal(t, err, nil)
		reply.Time = NewChinaTime(tValue)
		assert.Equal(t, reply.Time.String(), "2020-07-31 23:29:29 +0800 CST")
	}
}
