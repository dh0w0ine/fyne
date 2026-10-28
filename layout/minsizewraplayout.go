package layout

import (
	"fyne.io/fyne/v2"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*minSizeWrapLayout)(nil)

type minSizeWrapLayout struct {
	origLayout fyne.Layout
	minSize    fyne.Size
}

func NewMinSizeWrapLayout(baseLayout fyne.Layout, minSize fyne.Size) fyne.Layout {
	return &minSizeWrapLayout{origLayout: baseLayout, minSize: minSize}
}

func (m *minSizeWrapLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	m.origLayout.Layout(objects, size)
}

func (m *minSizeWrapLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	// Content minimum size
	contentMin := m.origLayout.MinSize(objects)
	// Returns whichever is larger: content minimum size or the specified min size
	return contentMin.Max(m.minSize)
}
