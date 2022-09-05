package eqmatch

import "errors"

type Persons struct {
	List   []Person            // 人员列表
	Quick  map[string]Person   // 根据ID快速查询用户信息
	Total  int                 // 人员总数
	Num    int                 // 每个人(被)面试的人数
	Result map[string][]string // 返回的结果
}

type Person struct {
	Id   string //人员id
	Type string // 人员类型, 如: 前端, 后端, 测试...
	Name string // 人员名字
}

var ErrMoreNum = errors.New("导入名单中人员数量和分配名额不匹配")
