# eqmatch
平等概率随机匹配算法

##### 适用场景
假如你的团队需要每个人选两(多)个人玩游戏或者做其他事情, 自己也得保证选两(多)次, 那么此时这个算法就非常适合你.

##### 导入需要excel数据, 导入导出格式在文件模板里有

##### API说明
###### New(filename, sheetName string, num int) (ps Persons, err error)
将excel导入文件中的数据导入到结构体当中并返回

###### Push2Result() error
将原始数据转化为结果数据并存对象结构当中

###### Print()
将结果以控制台的方式打印出来

###### Save(filename string) error
将结果存储到excel导出文件中

##### 注: 此代码可以直接作为库函数适用, 若还有调整算法的地方, 也可以直接clone下来修改, 作为工具包使用.
