package ui

import (
	"fmt"
	"grd/utils"
	"log"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type NodeType int

const (
	ReleaseNode NodeType = iota
	AssetNode
)

type NodeData struct {
	Type    NodeType
	Release *utils.Release
	Asset   *utils.Asset
}

func (a *App) populateReleaseTree(releases []utils.Release) {
	// Reset the tree
	a.releaseView.GetRoot().ClearChildren()
	a.releaseView.SetCurrentNode(a.releaseView.GetRoot())

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
			node.SetColor(tcell.ColorYellow)
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
			}

			a.populateReleaseTree(releases)
		}
	})

}

func (a *App) bindEvents() {
	a.urlAction()
}
