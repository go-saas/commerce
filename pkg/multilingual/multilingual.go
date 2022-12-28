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

// Multilingual empty interface
type Multilingual[T HasLanguageCode] interface {
	GetTranslations() []T
}

func GetTranslation[T HasLanguageCode](w Multilingual[T], tags ...language.Tag) (T, bool) {
	var keys []language.Tag
	m := lo.SliceToMap(w.GetTranslations(), func(item T) (string, T) {
		l := language.Make(item.GetLanguageCode())
		keys = append(keys, l)
		return l.String(), item
	})

	matcher := language.NewMatcher(keys)
	t, index, _ := matcher.Match(tags...)
	var defaultT T
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

func (l *Language) GetTranslations() []*LanguageTranslation {
	return l.Translations
}

type LanguageTranslation struct {
	Embed
	Name string
}
