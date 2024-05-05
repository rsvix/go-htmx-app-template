package snippetshandler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rsvix/go-htmx-app-template/internal/structs"
	"github.com/rsvix/go-htmx-app-template/internal/templates"
	"github.com/rsvix/go-htmx-app-template/internal/utils"
	"gorm.io/gorm"
)

type getSnippetsHandlerParams struct {
	pageTitle string
}

func GetSnippetsHandler() *getSnippetsHandlerParams {
	return &getSnippetsHandlerParams{
		pageTitle: "Snippets",
	}
}

func (h getSnippetsHandlerParams) Serve(c echo.Context) error {
	sessionInfo, err := utils.GetSessionInfo(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	db := c.Get("__db").(*gorm.DB)
	var result []structs.Snippet
	_ = db.Raw("SELECT * FROM snippets WHERE owner = ? OR ispublic = ?;", sessionInfo.Id, "1").Scan(&result)

	var m = make(map[string]bool)
	var langArr []structs.Languages

	for _, v := range result {
		lang := v.Language
		if m[lang] {
			for index := range langArr {
				if langArr[index].Language == lang {
					langArr[index].Count = langArr[index].Count + 1
				}
			}
		} else {
			langArr = append(langArr, structs.Languages{Language: lang, Count: 1})
			m[lang] = true
		}
	}
	log.Println(langArr)
	totalSnippets := len(result)
	log.Println(totalSnippets)

	return templates.SnippetsPage(c, h.pageTitle, sessionInfo.Username, result, langArr, totalSnippets).Render(c.Request().Context(), c.Response())
}
