package job

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"vizzy/dao"
	"vizzy/entity"
)

func Page3Job(dao dao.Page3Dao) (err error) {
	var ss = []byte(`{
	  "xAxis": {
		"type": "category",
		"data": ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"]
	  },
	  "yAxis": {
		"type": "value"
	  },
	  "series": [
		{
		  "data": [120, 200, 150, 80, 70, 110, 130],
		  "type": "bar"
		}
	  ],
      "grid": {"containLabel":true,"left":10,"right":10,"top":20,"bottom":20}
	}`)
	var origin entity.Page3RelEntity
	err = json.Unmarshal(ss, &origin)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = dao.Create(entity.Page3{
		Entity: origin,
	})
	if err != nil {
		logrus.Error(err)
		return err
	}
	return err
}
