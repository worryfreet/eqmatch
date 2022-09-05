package eqmatch

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func New(filename, sheetName string, num int) (ps Persons, err error) {
	excel, err := excelize.OpenFile(filename)
	if err != nil {
		return
	}
	defer func() {
		if err = excel.Close(); err != nil {
			return
		}
	}()
	rows, err := excel.GetRows(sheetName)
	if err != nil {
		return
	}
	ps.Total = len(rows) - 1
	ps.Num = num
	ps.List = make([]Person, ps.Total)
	ps.Quick = make(map[string]Person, ps.Total)
	ps.Result = make(map[string][]string, ps.Total)
	for i := 1; i < len(rows); i++ {
		tmpPerson := Person{
			Id:   rows[i][0],
			Type: rows[i][1],
			Name: rows[i][2],
		}
		ps.List[i-1] = tmpPerson
		ps.Quick[tmpPerson.Id] = tmpPerson
	}
	return
}

func (ps Persons) Push2Result() error {
	return ps.convert()
}

func (ps Persons) convert() error {
	if ps.Num >= ps.Total {
		return ErrMoreNum
	}
	// 根据不同分类分成不同的List
	typeMap := make(map[string][]Person)
	for i := 0; i < ps.Total; i++ {
		typeMap[ps.List[i].Type] = make([]Person, 0)
	}
	for i := 0; i < ps.Total; i++ {
		typeMap[ps.List[i].Type] = append(typeMap[ps.List[i].Type], ps.List[i])
	}
	for _, list := range typeMap {
		if len(list) <= ps.Num {
			return errors.New("导入名单中人员数量和分配名额不匹配")
		}
	}
	// 核心代码---随机数组, 多协程快速检索
	wg := sync.WaitGroup{}
	wg.Add(len(typeMap))
	randArray := func(list []Person) {
		idxArr := make([]int, len(list))
		for i := 0; i < len(list); i++ {
			idxArr[i] = i
		}
		finalIdxArr := make([]int, len(list))
		for i := 0; i < len(finalIdxArr); i++ {
			rand.Seed(time.Now().UnixNano())
			rd := rand.Intn(len(idxArr))
			finalIdxArr[i] = idxArr[rd]
			idxArr = append(idxArr[:rd], idxArr[rd+1:]...)
		}
		for i := 0; i < len(list); i++ {
			for j := 0; j < ps.Num; j++ {
				ps.Result[list[finalIdxArr[i]].Id] = append(ps.Result[list[finalIdxArr[i]].Id], list[finalIdxArr[(i+j+1)%len(list)]].Id)
			}
		}
		wg.Done()
	}
	for _, list := range typeMap {
		go randArray(list)
	}
	wg.Wait()
	return nil
}

func (ps Persons) Print() {
	for i := 0; i < len(ps.List); i++ {
		fmt.Println(ps.List[i].Name, ": ", ps.Result[ps.List[i].Id])
	}
}

func (ps Persons) Save(filename string) error {
	f := excelize.NewFile()
	sheetName := "Sheet1"
	sheet := f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "ID")
	f.SetCellValue(sheetName, "B1", "姓名")
	for i := 0; i < ps.Num; i++ {
		f.SetCellValue(sheetName, string(rune('C'+i))+"1", "被面人员"+strconv.Itoa(i+1))
	}
	for i := 0; i < len(ps.List); i++ {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(i+2), i+1)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(i+2), ps.List[i].Name)
		for j := 0; j < ps.Num; j++ {
			f.SetCellValue(sheetName, string(rune('C'+j))+strconv.Itoa(i+2), ps.Quick[ps.Result[ps.List[i].Id][j]].Name)
		}
	}
	f.SetActiveSheet(sheet)
	return f.SaveAs(filename)
}
