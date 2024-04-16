package view

import (
	"fmt"

	"github.com/keidarcy/e1s/internal/utils"
	"github.com/rivo/tview"
)

const (
	footerSelectedItemFmt = " [black:aqua:b] <%s> [-:-:-] "
	footerItemFmt         = " [black:grey:] <%s> [-:-:-] "
)

// View footer struct
type footer struct {
	footerFlex     *tview.Flex
	cluster        *tview.TextView
	service        *tview.TextView
	task           *tview.TextView
	container      *tview.TextView
	taskDefinition *tview.TextView
}

func newFooter() *footer {
	return &footer{
		footerFlex:     tview.NewFlex().SetDirection(tview.FlexColumn),
		cluster:        tview.NewTextView().SetDynamicColors(true).SetText(fmt.Sprintf(footerItemFmt, ClusterKind)),
		service:        tview.NewTextView().SetDynamicColors(true).SetText(fmt.Sprintf(footerItemFmt, ServiceKind)),
		task:           tview.NewTextView().SetDynamicColors(true).SetText(fmt.Sprintf(footerItemFmt, TaskKind)),
		container:      tview.NewTextView().SetDynamicColors(true).SetText(fmt.Sprintf(footerItemFmt, ContainerKind)),
		taskDefinition: tview.NewTextView().SetDynamicColors(true).SetText(fmt.Sprintf(footerItemFmt, TaskDefinitionKind)).SetTextAlign(L),
	}
}
func (v *view) addFooterItems() {
	// left resources
	v.footer.footerFlex.AddItem(v.footer.cluster, 13, 0, false).
		AddItem(v.footer.service, 13, 0, false).
		AddItem(v.footer.task, 10, 0, false).
		AddItem(v.footer.container, 15, 0, false)

	// keep middle space
	v.footer.footerFlex.
		AddItem(tview.NewTextView(), 5, 0, false).
		AddItem(v.footer.taskDefinition, 0, 1, false)

	// right labels
	// name version label
	regionLabel := tview.NewTextView().SetDynamicColors(true).SetText(fmt.Sprintf("[black:yellow:bi] %s ", v.app.Region))
	v.footer.footerFlex.AddItem(regionLabel, len(v.app.Region)+3, 0, false)

	appLabel := fmt.Sprintf("[black:blue:bi] %s:%s ", utils.AppName, utils.AppVersion)
	t := tview.NewTextView().SetTextAlign(L).SetDynamicColors(true).SetText(appLabel)
	v.footer.footerFlex.AddItem(t, 15, 1, false)
}