package dto

import "vizzy/entity"

type Page1Dto struct {
	Title struct {
		Text string `json:"text"`
	} `json:"title"`
	Legend struct {
		Data  []string `json:"data"`
		Right int      `json:"right"`
	} `json:"legend"`
	XAxis struct {
		Type     string   `json:"type"`
		Data     []string `json:"data"`
		Show     bool     `json:"show"`
		AxisLine struct {
			Show bool `json:"show"`
		} `json:"axisLine"`
		SplitLine struct {
			Show bool `json:"show"`
		} `json:"splitLine"`
	} `json:"xAxis"`
	YAxis struct {
		Type     string `json:"type"`
		AxisLine struct {
			Show bool `json:"show"`
		} `json:"axisLine"`
		SplitLine struct {
			Show bool `json:"show"`
		} `json:"splitLine"`
	} `json:"yAxis"`
	Series []entity.Page1EntitySeries `json:"series"`
	Grid   struct {
		ContainLabel bool `json:"containLabel"`
		Left         int  `json:"left"`
		Right        int  `json:"right"`
		Top          int  `json:"top"`
		Bottom       int  `json:"bottom"`
	} `json:"grid"`
}
