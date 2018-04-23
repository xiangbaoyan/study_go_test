package parser

import (
	"github.com/xiangbaoyan/study_go_test/crawler/engine"
	"github.com/xiangbaoyan/study_go_test/crawler/modal"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var houseRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var carRe = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

//<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">HuCL</a>
var guessRes = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)>([^<]+)`)

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([0-9]+)`)

func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := modal.Profile{}
	profile.Name = name

	age, e := strconv.Atoi(string(extractString(contents, ageRe)))
	if e != nil {
		profile.Age = age
	}
	profile.Gender = extractString(contents, genderRe)
	profile.Marriage = extractString(contents, marriageRe)
	height, e := strconv.Atoi(string(extractString(contents, heightRe)))
	if e == nil {
		profile.Height = height
	}
	weight, e := strconv.Atoi(string(extractString(contents, weightRe)))
	if e == nil {
		profile.Weight = weight
	}

	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Xinzuo = extractString(contents, xinzuoRe)
	profile.House = extractString(contents, houseRe)
	profile.Car = extractString(contents, carRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	matches := guessRes.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		name := string(m[2])
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:       url,
			ParseFunc: ProfileParser(name, url),
		})
	}
	return result

}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	} else {
		return ""
	}

}

func ProfileParser(name string, url string) engine.ParseFunc {
	return func(c []byte) engine.ParseResult {
		return ParseProfile(c, url, name)
	}
}
