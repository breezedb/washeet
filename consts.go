package washeet

const (
	constMaxRowCount           int64   = 1048576
	constMaxRow                int64   = constMaxRowCount - 1
	constMaxColCount           int64   = 16384
	constMaxCol                int64   = constMaxColCount - 1
	constDefaultCellWidth      float64 = 90.0
	constDefaultCellHeight     float64 = 25.0
	constSheetPaintQueueLength int     = 50

	SELECTION_FILL_COLOR string = "rgba(100, 156, 244, 0.2)"
)

const (
	sheetPaintWholeSheet sheetPaintType = iota
	sheetPaintCell
	sheetPaintCellRange
	sheetPaintSelection
)

const (
	AlignLeft TextAlignType = iota
	AlignCenter
	AlignRight
)
