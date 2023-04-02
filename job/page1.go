package job

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"vizzy/dao"
	"vizzy/entity"
)

type page1ServerData struct {
	Type, Name string
	X          string
	Y          int
}

func Page1Job(dao dao.Page1Dao) (err error) {
	var ss = []byte(`[
		{"type": "line", "name": "实线", "x": "1", "y": 120},
		{"type": "line", "name": "虚线", "x": "1", "y": 220},
		{"type": "line", "name": "实线", "x": "2", "y": 132},
		{"type": "line", "name": "虚线", "x": "2", "y": 182},
		{"type": "line", "name": "实线", "x": "3", "y": 101},
		{"type": "line", "name": "虚线", "x": "3", "y": 131},
		{"type": "line", "name": "实线", "x": "4", "y": 120},
		{"type": "line", "name": "虚线", "x": "4", "y": 150},
		{"type": "line", "name": "实线", "x": "5", "y": 120},
		{"type": "line", "name": "虚线", "x": "5", "y": 140},
		{"type": "line", "name": "实线", "x": "6", "y": 130},
		{"type": "line", "name": "虚线", "x": "6", "y": 120}
	]`)
	var origin []page1ServerData
	err = json.Unmarshal(ss, &origin)
	if err != nil {
		logrus.Error(err)
		return err
	}
	names := make(map[string]bool)
	xMap := make(map[string]bool)
	yMap := make(map[string][]int)
	for _, d := range origin {
		names[d.Name] = true
		yMap[d.Name] = append(yMap[d.Name], d.Y)
		xMap[d.X] = true
	}
	nameArray := make([]string, 0, len(names))
	for name := range names {
		nameArray = append(nameArray, name)
	}
	xArray := make([]string, 0, len(xMap))
	for x := range xMap {
		xArray = append(xArray, x)
	}
	var series []entity.Page1EntitySeries
	for name, yValues := range yMap {
		series = append(series, entity.Page1EntitySeries{
			Name: name,
			Type: "line",
			Data: yValues,
			//LineStyle: struct {
			//	Type string `json:"type"`
			//}{},
		})
	}

	err = dao.Create(entity.Page1{
		Entity: entity.Page1RelEntity{
			Title: struct {
				Text string `json:"text"`
			}{Text: ""},
			Legend: struct {
				Data  []string `json:"data"`
				Right int      `json:"right"`
			}{
				Data:  nameArray,
				Right: 10,
			},
			XAxis: struct {
				Type     string   `json:"type"`
				Data     []string `json:"data"`
				Show     bool     `json:"show"`
				AxisLine struct {
					Show bool `json:"show"`
				} `json:"axisLine"`
				SplitLine struct {
					Show bool `json:"show"`
				} `json:"splitLine"`
			}{
				Type: "category",
				Data: xArray,
				Show: false,
				AxisLine: struct {
					Show bool `json:"show"`
				}{
					Show: false,
				},
				SplitLine: struct {
					Show bool `json:"show"`
				}{
					Show: false,
				},
			},
			YAxis: struct {
				Type     string `json:"type"`
				AxisLine struct {
					Show bool `json:"show"`
				} `json:"axisLine"`
				SplitLine struct {
					Show bool `json:"show"`
				} `json:"splitLine"`
			}{
				Type: "value",
				AxisLine: struct {
					Show bool `json:"show"`
				}{
					Show: false,
				},
				SplitLine: struct {
					Show bool `json:"show"`
				}{
					Show: false,
				},
			},
			Series: series,
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
