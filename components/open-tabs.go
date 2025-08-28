package components

import (
	"fmt"

	"github.com/dudubtw/figma/app"
)

func selectedPageId() app.Page {
	if app.Apk.SelectedPage != app.PAGE_WORKPLACE {
		return app.Apk.SelectedPage
	}

	return app.PAGE_WORKPLACE + "_" + app.Apk.Workplace.Id
}

func OpenTabs() app.Component {
	tabs :=
		Tabs().
			Selected(selectedPageId()).
			AddIcon(app.PAGE_HOME, app.ICON_HOUSE)

	if app.Apk.SelectedPage == app.PAGE_WORKPLACE {
		tabs.Add(selectedPageId(), app.Apk.Workplace.Id, app.ICON_FILE)
	}
	return tabs.Component(func(clickedId string) {
		fmt.Println("clickd!", clickedId)
		if clickedId == app.PAGE_HOME {
			app.Apk.Navigate(app.PAGE_HOME)
		}
	})
}
