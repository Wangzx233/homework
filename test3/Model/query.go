package Model

import (
	"fmt"
	"sort"
)

func QueryRemarks(remarks string) (Records) {
	row,err:=DB.Query("select * from records where remarks like CONCAT('%',?,'%');",remarks)
	if err!=nil{
		fmt.Println("查询备注失败")
	}
	defer row.Close()
	var records Records
	time:=""
	for row.Next() {
		err = row.Scan(&records.Id,&records.FromUid,&records.ToUid,&records.Money,&records.Remarks,&time)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return records
		}
	}
	return records
}
type re []Records
func QueryUid(uid string) []Records {
	row,err:=DB.Query("select * from records where fromUid=?;",uid)
	if err!=nil{
		fmt.Println("查询备注失败")
	}
	defer row.Close()
	var records Records
	var re re
	for row.Next() {
		err = row.Scan(&records.Id,&records.FromUid,&records.ToUid,&records.Money,&records.Remarks,&records.Createtime)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
		}
		re=append(re,records)
	}
	//按照金钱排序
	sort.Sort(re)
	return re
}
func (p re) Swap(i, j int)   { p[i], p[j] = p[j], p[i] }
func (p re) Len() int      { return len(p) }
func (p re) Less(i, j int) bool { return  p[i].Money<p[j].Money }