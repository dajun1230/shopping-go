package api

import "shopping-go/api/system"

type ApiGroup struct {
	System system.ApiGroup
}

var ApiRouterGroup = new(ApiGroup)
