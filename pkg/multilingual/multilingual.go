package multilingual

import (
	"github.com/samber/lo"
	"golang.org/x/text/language"
)

// Embed translation table should embed this struct
type Embed struct {
	LanguageCode string
}

type HasLanguageCode interface {
	GetLanguageCode() string
}

func (e *Embed) GetLanguageCode() string {
	return e.LanguageCode
}

// Multilingual Multilingual
type Multilingual interface {
	GetTranslations() []interface{}
}

func GetTranslation(w Multilingual, tags ...language.Tag) (interface{}, bool) {
	var keys []language.Tag
	m := lo.SliceToMap(w.GetTranslations(), func(item interface{}) (string, interface{}) {
		l := language.Make(item.(HasLanguageCode).GetLanguageCode())
		keys = append(keys, l)
		return l.String(), item
	})

	matcher := language.NewMatcher(keys)
	t, index, _ := matcher.Match(tags...)
	var defaultT interface{}
	if t == language.Und {
		return defaultT, false
	}
	//t is incorrect https://github.com/golang/go/issues/24211#issuecomment-378479543
	t = keys[index]
	trans, ok := m[t.String()]
	if !ok {
		return defaultT, false
	}
	return trans, true
}

type Language struct {
	Code         string `gorm:"type:char(36);primaryKey;"`
	Name         string
	Translations []*LanguageTranslation
}

func (l *Language) GetTranslations() []interface{} {
	return lo.Map(l.Translations, func(item *LanguageTranslation, _ int) interface{} {
		return item
	})
}

type LanguageTranslation struct {
	Embed
	Name string
}
