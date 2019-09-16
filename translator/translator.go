package translator

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

type Translator interface {
	RegisterUnmarshalFunc(format string, unmarshalFunc i18n.UnmarshalFunc)
	LoadMessageFile(path string) (*i18n.MessageFile, error)
	Localize(config *i18n.LocalizeConfig) string
	SetLanguage(lang string)
	GetLanguage() string
	TranslationFilesPath(path string)
}

var (
	Tr Translator
)

func Enable() {
	Tr = newTranslator()
}

func newTranslator() *translator {
	bundle := i18n.NewBundle(language.English)
	translate := &translator{bundle: bundle}
	translate.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	translate.TranslationFilesPath("translations")
	translate.SetLanguage("en")
	return translate
}

type translator struct {
	bundle           *i18n.Bundle
	localizer        *i18n.Localizer
	translationsPath string
}

func (t *translator) RegisterUnmarshalFunc(format string, unmarshalFunc i18n.UnmarshalFunc) {
	t.bundle.RegisterUnmarshalFunc(format, unmarshalFunc)
}

func (t *translator) TranslationFilesPath(fPath string) {
	var files []string

	dir, err := filepath.Abs(".")

	if err != nil {
		fmt.Println(err)
	}

	t.translationsPath = fmt.Sprintf("%s/%s", dir, fPath)

	if _, err := os.Stat(t.translationsPath); !os.IsNotExist(err) {
		if err != nil {
			fmt.Println(err)
		}

		if err := filepath.Walk(t.translationsPath, func(path string, info os.FileInfo, err error) error {
			fileInfo, err := os.Stat(path)

			if err != nil {
				fmt.Println(err)
			}

			if fileInfo.Mode().IsRegular() {
				files = append(files, path)
			}
			return nil
		}); err != nil {
			log.Println(err)
		}
	} else {
		fmt.Println(fmt.Errorf("directory %s does not exist", t.translationsPath))
	}

	for _, file := range files {
		if _, err := t.LoadMessageFile(file); err != nil {
			fmt.Println(err)
		}
	}
}

func (t *translator) SetLanguage(lang string) {
	fmt.Println(lang)
	t.localizer = i18n.NewLocalizer(t.bundle, lang)
}

func (t *translator) GetLanguage() string {
	return t.bundle.LanguageTags()[0].String()
}

func (t *translator) LoadMessageFile(path string) (*i18n.MessageFile, error) {
	absPath, err := filepath.Abs(path)

	if err != nil {
		return nil, err
	}

	return t.bundle.LoadMessageFile(absPath)
}

func (t *translator) Localize(config *i18n.LocalizeConfig) string {
	str, err := t.localizer.Localize(config)

	if err != nil {
		fmt.Println(err)
	}

	return str
}

func T(key string, data interface{}, count interface{}) string {
	return Tr.Localize(&i18n.LocalizeConfig{
		MessageID:    key,
		TemplateData: data,
		PluralCount:  count,
	})
}
