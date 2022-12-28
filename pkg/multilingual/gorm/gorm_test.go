package gorm

import (
	"context"
	"github.com/go-saas/commerce/pkg/multilingual"
	"github.com/go-saas/kit/pkg/localize"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = db.Debug()

	exitCode := m.Run()
	os.Exit(exitCode)
}

type testEntity struct {
	Code         string `gorm:"type:char(36);primaryKey;"`
	Name         string
	Translations []*testEntityTranslation `gorm:"foreignKey:TestEntityCode;references:Code"`
}

func (t *testEntity) GetTranslations() []interface{} {
	return lo.Map(t.Translations, func(item *testEntityTranslation, _ int) interface{} {
		return item
	})
}

type testEntityTranslation struct {
	multilingual.Embed
	Name           string
	TestEntityCode string
}

func TestDB(t *testing.T) {
	var err error
	err = db.Migrator().DropTable(&testEntity{}, &testEntityTranslation{})
	assert.NoError(t, err)
	err = db.AutoMigrate(&testEntity{}, &testEntityTranslation{})
	assert.NoError(t, err)

	preferredTags := []language.Tag{language.Chinese, language.SimplifiedChinese, language.AmericanEnglish, language.English}
	ctx := localize.NewLanguageTagsContext(context.Background(), preferredTags)

	db = db.WithContext(ctx)

	lan := &testEntity{
		Code: "zh-CN",
		Name: "中文",
		Translations: []*testEntityTranslation{
			{
				Embed: multilingual.Embed{LanguageCode: language.SimplifiedChinese.String()},
				Name:  "中文",
			},
			{
				Embed: multilingual.Embed{LanguageCode: language.Japanese.String()},
				Name:  "中国語",
			},
			{
				Embed: multilingual.Embed{LanguageCode: language.TraditionalChinese.String()},
				Name:  "繁体中文",
			},
			{
				Embed: multilingual.Embed{LanguageCode: language.English.String()},
				Name:  "Chinese",
			},
		},
	}

	err = db.Session(&gorm.Session{FullSaveAssociations: true}).Create(lan).Error
	assert.NoError(t, err)
	var dbLan testEntity

	err = db.Session(&gorm.Session{NewDB: true}).Model(&testEntity{}).Preload(clause.Associations).Find(&dbLan, "code = ?", "zh-CN").Error
	assert.NoError(t, err)
	assert.Equal(t, 4, len(dbLan.Translations))

	err = db.Session(&gorm.Session{NewDB: true}).Model(&testEntity{}).Scopes(PreloadCurrentLanguage()).Find(&dbLan, "code = ?", "zh-CN").Error
	assert.NoError(t, err)
	assert.Equal(t, 3, len(dbLan.Translations))

	trans, ok := multilingual.GetTranslation(&dbLan, language.Chinese)
	assert.True(t, ok)
	assert.NotNil(t, trans)

}
