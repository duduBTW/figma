package components

import (
	"github.com/dudubtw/figma/app"
	ds "github.com/dudubtw/figma/design-system"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TabsIntance struct {
	selected string
	items    []TabItemProps

	ClickedId string
}

type TabItemProps struct {
	Id       string
	IsIcon   bool
	IconName app.IconName
	Text     string
}

func Tabs() *TabsIntance {
	return &TabsIntance{
		items: []TabItemProps{},
	}
}

func (tabsInstance *TabsIntance) Selected(id string) *TabsIntance {
	tabsInstance.selected = id
	return tabsInstance
}

func (tabsInstance *TabsIntance) AddIcon(id string, iconName app.IconName) *TabsIntance {
	tabsInstance.items = append(tabsInstance.items, TabItemProps{IsIcon: true, IconName: iconName, Id: id})
	return tabsInstance
}

func (tabsInstance *TabsIntance) Add(id string, text string, iconName app.IconName) *TabsIntance {
	tabsInstance.items = append(tabsInstance.items, TabItemProps{IsIcon: false, IconName: iconName, Text: text, Id: id})
	return tabsInstance
}

func (tabsInstance *TabsIntance) Component(onClicked func(string)) app.Component {
	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			Direction(app.DIRECTION_ROW).
			Gap(ds.SPACING_2).
			PositionRect(rect)

		for _, props := range tabsInstance.items {
			layout.Add(tabsInstance.TabItem(props, onClicked))
		}

		return layout.Draw, layout.Size.Width, layout.Size.Height
	}
}

func (tabsInstance *TabsIntance) TabItem(props TabItemProps, onClicked func(string)) app.Component {
	isSelected := tabsInstance.selected == props.Id

	return func(rect rl.Rectangle) (func(), float32, float32) {
		layout := app.
			NewLayout().
			PositionRect(rect).
			Padding(app.NewPadding().All(ds.SPACING_2)).
			Direction(app.DIRECTION_ROW).
			Gap(ds.SPACING_2).
			Height(32)

		if !props.IsIcon {
			iconContrain := app.ChildSize{
				SizeType: app.SIZE_ABSOLUTE,
				Value:    app.ICON_WIDTH,
			}

			layout.
				Width(200,
					iconContrain,
					app.ChildSize{
						SizeType: app.SIZE_WEIGHT,
						Value:    1,
					},
					iconContrain,
				).
				Add(Icon(props.IconName)).
				Add(TabText(props.Text)).
				Add(Icon(app.ICON_X))
		} else {
			layout.Add(Icon(props.IconName))
		}

		bgRect := rl.NewRectangle(rect.X, rect.Y, layout.Size.Width, layout.Size.Height)
		interactable := app.NewInteractable("tab-item-" + props.Id)
		if !isSelected {
			clicked := interactable.Event(rl.GetMousePosition(), bgRect)
			if clicked {
				tabsInstance.ClickedId = props.Id
				onClicked(props.Id)
			}
		}

		return func() {
			DrawRectangleRoundedLinePixels(
				rl.NewRectangle(bgRect.X+1, bgRect.Y+1, bgRect.Width-2, bgRect.Height-2),
				ds.RADII_SM,
				1,
				ds.T2_BORDER,
			)

			bgColor := rl.NewColor(0, 0, 0, 0)
			if isSelected {
				bgColor = ds.T2_COLOR_SURFACE_LIGHT
			} else {
				switch interactable.State() {
				case app.STATE_HOT:
					bgColor = ds.COLOR_GRAY_300
				case app.STATE_ACTIVE:
					bgColor = ds.COLOR_GRAY_400
				}
			}

			DrawRectangleRoundedPixels(
				bgRect,
				ds.RADII_SM,
				bgColor,
			)

			layout.Draw()
		}, layout.Size.Width, layout.Size.Height
	}
}

func TabText(text string) app.Component {
	return Typography(text, ds.FONT_SIZE_LG, ds.FONT_WEIGHT_MEDIUM, ds.T2_COLOR_CONTENT)
}
