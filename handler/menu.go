package handler


//GetMenu 获取菜单
func GetMenu() string {
	backStr := "小机器人菜单\n"
	backStr += "[注册]注册账号，开始游戏\n"
	backStr += "[注销]注销账号，清空数据\n"
	backStr += "[查询]显示你的账号信息\n"
	backStr += "[资产排行][身价排行]查询资产、身价排行\n"
	backStr += "[转账+数字+@昵称]进行转账操作\n"
	backStr += "[猜数字]开始猜数字游戏，回复[我猜+数字]答题\n"
	return backStr
}


