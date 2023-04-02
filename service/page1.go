package service

import (
	"vizzy/common"
	"vizzy/dao"
)

type Page1Srv struct {
	Dao dao.Page1Dao
}

func (s *Page1Srv) Query(c *common.RouteContext) {
	record, err := s.Dao.Query()
	if err != nil {
		c.Fail(common.ResJSON{
			Code: 500,
			Data: nil,
			Msg:  "",
		})
		return
	}
	//var data = dto.Page1Dto{
	//	Title:  record.Entity.Title,
	//	Legend: record.Entity.Legend,
	//	XAxis:  record.Entity.XAxis,
	//	YAxis:  record.Entity.YAxis,
	//	Series: record.Entity.Series,
	//	Grid:   record.Entity.Grid,
	//}

	c.Success(record.Entity)
}
