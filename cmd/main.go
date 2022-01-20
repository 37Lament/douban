package cmd

import (
	"douban/api"
	"douban/dao"
)

func main() {
	dao.InitDB();
	api.InitEngine()

}