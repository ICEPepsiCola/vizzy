package job

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"vizzy/dao"
	"vizzy/entity"
)

func Page5Job(dao dao.Page5Dao) (err error) {
	var ss = []byte(`[
	  { "area": "天津", "count": 28, "percent": "40%", "duration": "1小时" },
	  { "area": "北京", "count": 36, "percent": "50%", "duration": "2小时" },
	  { "area": "香港", "count": 36, "percent": "60%", "duration": "3小时" },
	  { "area": "四川", "count": 36, "percent": "20%", "duration": "4小时" },
	  { "area": "广东", "count": 36, "percent": "30%", "duration": "5小时" },
	  { "area": "广西", "count": 36, "percent": "40%", "duration": "6小时" },
	  { "area": "湖南", "count": 36, "percent": "80%", "duration": "7小时" }
	]`)
	var origin []entity.Page5
	err = json.Unmarshal(ss, &origin)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = dao.Create(origin)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return err
}
