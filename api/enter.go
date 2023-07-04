package api

import "shopping-go/api/system"

type RouterGroup struct {
	System system.RouterGroup
}

var ApiRouterGroup = new(RouterGroup)
