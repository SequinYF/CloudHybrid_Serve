package model


func GetList(name string) (info UserInfo) {
	return
}

func ExistByName(name string) bool {
	//	var info UserInfo
	if CHdb == nil {
		println("nil")
	}
	v := CHdb.HasTable(&UserInfo{})
	if !v {
		panic("创建表失败")
	}
	/*
		if v.Value != nil {
			return true
		}
	*/
	return false
}

func AddUser(name string, passwd string, dir string) bool {
	CHdb.Create(&UserInfo{
		Name:   name,
		Passwd: passwd,
		Dir:    dir,
	})
	return true
}


