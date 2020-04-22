package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type TarWord struct {
	Id   int    `orm:"column(id);pk"`
	Word string `orm:"column(word)"`
}

func (t *TarWord) TableName() string {
	return "trie_nodes"
}

func init() {
	orm.RegisterModel(new(TarWord))
}

func AddAllWords(AllWords []string) (successNum int64, err error) {
	o := orm.NewOrm()
	NewWords := []TarWord{}
	for _, word := range AllWords {
		w := TarWord{Word: word}
		NewWords = append(NewWords, w)
	}
	successNum, err = o.InsertMulti(len(NewWords), NewWords)
	return successNum, nil
}

func GetAllWords() (res []*TarWord, err error) {
	o := orm.NewOrm()
	var nodes []*TarWord
	num, err := o.QueryTable("trie_nodes").All(&nodes)
	fmt.Println(num)
	return nodes, nil
}
