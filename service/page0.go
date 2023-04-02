package service

import (
	"math/rand"
	"vizzy/common"
)

type Page0Srv struct {
}

func (s *Page0Srv) Query(c *common.RouteContext) {
	c.Success(map[string]int{
		"num1": rand.Intn(1000),
		"num2": rand.Intn(100),
		"num3": rand.Intn(100),
		"num4": rand.Intn(100),
	})
}
