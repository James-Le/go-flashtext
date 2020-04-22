package controllers

import (
	"encoding/json"
	"flash-text-api/models"
	"fmt"

	"github.com/astaxie/beego"
)

//TextController ...
type TextController struct {
	beego.Controller
	trie *models.Trie
}

//QueryWords ...
type QueryWords struct {
	Word []string `json:"word"`
}

// AddAllWords ...
func (tc *TextController) AddAllWords() {

	var queryWords QueryWords

	data := tc.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &queryWords)
	if err != nil {
		fmt.Println("invalid json format!")
	}
	successNums, err := models.AddAllWords(queryWords.Word)
	tc.Data["json"] = successNums
	tc.ServeJSON()
}

// ListWords ...
func (tc *TextController) ListWords() {
	res, err := models.GetAllWords()
	if err != nil {
		tc.Data["json"] = err.Error()
	} else {
		tc.Data["json"] = res
	}
	tc.ServeJSON()
}

//Parse ...
func (tc *TextController) Parse() {
	if tc.trie == nil {
		tc.trie = models.InitTrie()
		res, err := models.GetAllWords()
		if err == nil {
			for _, node := range res {
				tc.trie.Insert(node.Word)
			}
		}
	}

	var queryWords QueryWords

	data := tc.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &queryWords)
	if err != nil {
		fmt.Println("invalid json format!")
	}
	var tarWords []string
	for _, word := range queryWords.Word {
		if tc.trie.Search(word) == true {
			tarWords = append(tarWords, word)
		}
	}
	tc.Data["json"] = tarWords
	tc.ServeJSON()
}
