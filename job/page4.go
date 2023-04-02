package job

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"vizzy/dao"
	"vizzy/entity"
)

func Page4Job(dao dao.Page4Dao) (err error) {
	var ss = []byte(`{
	  "xAxis": {
		"type": "value"
	  },
	  "yAxis": {
		"type": "category",
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	  },
	  "series": [
		{
		  "data": [120, 200, 150, 80, 70, 110, 130],
		  "type": "bar"
		}
	  ],
      "grid": {"containLabel":true,"left":10,"right":10,"top":0,"bottom":0}
	}`)
	var origin entity.Page4RelEntity
	err = json.Unmarshal(ss, &origin)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = dao.Create(entity.Page4{
		Entity: origin,
	})
	if err != nil {
		logrus.Error(err)
		return err
	}
	return err
}
