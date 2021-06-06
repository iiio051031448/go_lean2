package parser

import (
	"regexp"
	"strconv"
	"stu/http/crawler/engine"
	"stu/http/crawler/model"
)

//var genderRe = regexp.MustCompile(`<span class="label">([^<]+)</span><span field="">([^>]+)</span>`)
var profileRe1 = regexp.MustCompile(`<td><span class="label">([^<]+)</span>([^<]+)</td>`)
var profileRe2 = regexp.MustCompile(`<td><span class="label">([^<]+)</span><span field="">([^<]+)</span></td>`)

func ParserProfile(contents []byte) engine.ParseResult {
	//var profileMap map[string]string
	profileMap := make(map[string]string)

	matchs1 := profileRe1.FindAllSubmatch(contents, -1)
	for _, m := range matchs1 {
		profileMap[string(m[1])] = string(m[2])
	}
	matchs2 := profileRe2.FindAllSubmatch(contents, -1)
	for _, m := range matchs2 {
		profileMap[string(m[1])] = string(m[2])
	}

	ageNum := func(ageString string) int {
		a := []rune(ageString)
		a = a[:len(a)-1]
		an, err := strconv.Atoi(string(a))
		if err != nil {
			return 0
		}
		return an
	}
	hNum := func(ageString string) int {
		a := ageString
		a = a[:len(a)-2]
		an, err := strconv.Atoi(a)
		if err != nil {
			return 0
		}
		return an
	}

	profile := model.Profile{
		Name:       "",
		Gender:     profileMap["性别："],
		Age:        ageNum(profileMap["年龄："]),
		Height:     hNum(profileMap["身高："]),
		Weight:     hNum(profileMap["体重："]),
		Income:     profileMap["月收入："],
		Marriage:   profileMap["婚况："],
		Education:  profileMap["性别："],
		Occupation: profileMap["学历："],
		HuKou:      profileMap["籍贯："],
		Xinzou:     profileMap["星座："],
		House:      profileMap["住房条件："],
		Car:        profileMap["是否购车："],
	}

	return engine.ParseResult{Items: []interface{}{profile}}
}