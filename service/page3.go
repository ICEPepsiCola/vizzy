package service

import (
	"vizzy/common"
	"vizzy/dao"
)

type Page3Srv struct {
	Dao dao.Page3Dao
}

func (s *Page3Srv) Query(c *common.RouteContext) {
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
