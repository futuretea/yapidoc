package router

import (
	"github.com/futuretea/go-yapi"
	"github.com/futuretea/yapidoc/logs"
	"github.com/gin-gonic/gin"
)

type Service struct {
	Name string
	APIS []yapi.InterfaceData
}

type Result struct {
	Name     string
	Services []Service
}

func containThis(str string, arr []string) bool {
	for _, v := range arr {
		if str == v {
			return true
		}
	}
	return false
}

func abortWith403(c *gin.Context, err error, msg string) {
	if err != nil {
		logs.Fatal(msg, err)
		c.AbortWithStatus(403)
	}
}

func render(c *gin.Context, apiURL string, apiToken string, tag string) *Result {
	yapiClient, err := yapi.NewClient(nil, apiURL, apiToken)
	if err != nil {
		logs.Fatal("new yapi client err", err)
	}
	project, _, err := yapiClient.Project.Get()
	abortWith403(c, err, "get project err")
	result := Result{
		Name:     project.Data.Name,
		Services: []Service{},
	}
	catMenus, _, err := yapiClient.CatMenu.Get(project.Data.ID)
	abortWith403(c, err, "get catmenus err")
	for _, catmenu := range catMenus.Data {
		service := Service{
			Name: catmenu.Name,
			APIS: []yapi.InterfaceData{},
		}
		interfaceListParam := yapi.InterfaceListParam{
			CatID: catmenu.ID,
			Page:  1,
			Limit: 1000000,
		}
		interfaces, _, err := yapiClient.Interface.GetList(&interfaceListParam)
		abortWith403(c, err, "get interfaces err")
		for _, i := range interfaces.Data.List {
			r, _, err := yapiClient.Interface.Get(i.ID)
			abortWith403(c, err, "get interface err")
			if containThis(tag, r.Data.Tag) || tag == "" {
				service.APIS = append(service.APIS, r.Data)
			}
		}
		result.Services = append(result.Services, service)
	}
	return &result
}
