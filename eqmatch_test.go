package eqmatch

import (
	"fmt"
	"testing"
)

func TestEqMatch(t *testing.T) {
	persons, err := New("导入名单模板.xlsx", "Sheet1", 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err = persons.Push2Result(); err != nil {
		fmt.Println(err)
		return
	}
	persons.Print()
	if err = persons.Save("导出名单模板.xlsx"); err != nil {
		fmt.Println(err)
		return
	}
}
