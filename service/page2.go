package service

import (
	"vizzy/common"
	"vizzy/dao"
)

type Page2Srv struct {
	Dao dao.Page2Dao
}

func (s *Page2Srv) Query(c *common.RouteContext) {
	record, err := s.Dao.Query()
	if err != nil {
		c.Fail(common.ResJSON{
			Code: 500,
			Data: nil,
			Msg:  "",
		})
		return
	}
	c.Success(record.Entity)
}
