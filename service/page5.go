package service

import (
	"vizzy/common"
	"vizzy/dao"
)

type Page5Srv struct {
	Dao dao.Page5Dao
}

func (s *Page5Srv) Query(c *common.RouteContext) {
	record, err := s.Dao.Query()
	if err != nil {
		c.Fail(common.ResJSON{
			Code: 500,
			Data: nil,
			Msg:  "",
		})
		return
	}
	c.Success(record)
}
