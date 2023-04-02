package job

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"vizzy/dao"
	"vizzy/entity"
)

func Page2Job(dao dao.Page2Dao) (err error) {
	var ss = []byte(`[{
      "name": "Access From",
      "type": "pie",
      "radius": "90%",
      "data": [
        { "value": 1048, "name": "Search Engine" },
        { "value": 735, "name": "Direct" },
        { "value": 580, "name": "Email" },
        { "value": 484, "name": "Union Ads" },
        { "value": 300, "name": "Video Ads" }
      ],
      "emphasis": {
        "itemStyle": {
          "shadowBlur": 10,
          "shadowOffsetX": 0,
          "shadowColor": "rgba(0, 0, 0, 0.5)"
        }
      }
    }]`)
	var origin []entity.Page2EntitySeries
	err = json.Unmarshal(ss, &origin)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = dao.Create(entity.Page2{
		Entity: entity.Page2RelEntity{
			Title: struct{}{},
			//Legend: struct {
			//	Orient string `json:"orient"`
			//	Left   string `json:"left"`
			//}{Orient: "vertical", Left: "left"},
			Series: origin,
			Grid: struct {
				ContainLabel bool `json:"containLabel"`
				Left         int  `json:"left"`
				Right        int  `json:"right"`
				Top          int  `json:"top"`
				Bottom       int  `json:"bottom"`
			}{
				ContainLabel: true,
				Left:         10,
				Right:        10,
				Top:          20,
				Bottom:       20,
			},
		},
	})
	if err != nil {
		logrus.Error(err)
		return err
	}
	return err
}
