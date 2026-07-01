package ui

import (
	"fmt"
	"log"
	"time"

	"github.com/ArmanHZ/git-release-downloader/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) initInputCapture() {
	a.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyTab:
			a.activeFocus.Next(a.app)
			return nil

		case tcell.KeyBacktab:
			a.activeFocus.Prev(a.app)
			return nil
		}

		return event
	})
}

func (a *App) resetReleaseTree() {
	a.releaseView.GetRoot().ClearChildren()
	a.releaseView.SetCurrentNode(a.releaseView.GetRoot())
	a.downloadList = map[utils.Asset][]string{}
}

func (a *App) populateReleaseTree(releases []utils.Release) {
	a.resetReleaseTree()

	root := a.releaseView.GetRoot()

	for _, release := range releases {
		rel := release

		// XXX: Published date or update date? idk
		t, err := time.Parse(time.RFC3339, rel.PublishedAt)
		if err != nil {
			log.Fatal(err)
		}

		formatted := t.Format("2006-01-02")

		releaseNodeText := fmt.Sprintf("%s [gray](%s)[-]", rel.TagName, formatted)
		releaseNode := tview.NewTreeNode(releaseNodeText).
			SetReference(rel)

		root.AddChild(releaseNode)
	}

	a.releaseView.SetSelectedFunc(func(node *tview.TreeNode) {
		ref := node.GetReference()

		if ref == nil {
			return
		}

		// TODO: More readable name
		switch v := ref.(type) {
		case utils.Release:
			// Only populate once
			if len(node.GetChildren()) == 0 {
				space := utils.AssetDigestSpaceCalc(v.Assets)
				for _, asset := range v.Assets {
					asset := asset
					assetNodeText := fmt.Sprintf("%-*s [gray]%s[-]", space, asset.Name, asset.Digest)
					assetNode := tview.NewTreeNode(assetNodeText).
						SetReference(asset)
					node.AddChild(assetNode)
				}
				node.SetExpanded(true)
			} else {
				node.SetExpanded(!node.IsExpanded())
			}

		// Asset selected
		case utils.Asset:
			if _, ok := a.downloadList[v]; ok {
				delete(a.downloadList, v)
				node.SetColor(tcell.ColorWhite)
			} else {
				a.downloadList[v] = []string{v.Name, v.BrowserDownloadURL}
				node.SetColor(tcell.ColorYellow)
			}
		default:
			node.SetExpanded(!node.IsExpanded())
		}
	})

}

func (a *App) urlAction() {
	a.urlInput.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			releases, err := utils.GetReleases(a.urlInput.GetText())
			if err != nil {
				// TODO: Pop-up saying error or something.
			} else {
				a.populateReleaseTree(releases)
				a.activeFocus.index = int(ReleaseView)
				a.app.SetFocus(a.releaseView)
			}
		}
	})

}

func (a *App) downloadButtonAction() {
	a.downloadButton.SetSelectedFunc(func() {
		modal := a.buildDownloadPage()
		a.app.SetRoot(modal, true).SetFocus(modal)
	})
}

func (a *App) bindEvents() {
	a.urlAction()
	a.downloadButtonAction()
}
