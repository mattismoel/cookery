package model

type Unit string

type Ingredient struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount,omitempty"`
	Unit   Unit    `json:"unit,omitempty"`
	Note   string  `json:"note,omitempty"`
}

type SubRecipe struct {
	Title        string        `json:"title,omitempty"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
}

type Instruction struct {
	Text     string `json:"text"`
	ImageUrl string `json:"imageUrl,omitempty"`
}

type Recipe struct {
	Id           int64       `json:"id"`
	Title        string      `json:"title"`
	Author       User        `json:"author"`
	BannerUrl    string      `json:"bannerUrl"`
	CookMinutes  int         `json:"cookMinutes"`
	TotalMinutes int         `json:"totalMinutes"`
	SubRecipes   []SubRecipe `json:"subRecipes"`
}
