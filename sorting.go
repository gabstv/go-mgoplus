package mgoplus

type SortCollectionStatsByIndexSizeDesc []CollectionStats

func (c SortCollectionStatsByIndexSizeDesc) Len() int {
	return len(c)
}

func (c SortCollectionStatsByIndexSizeDesc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortCollectionStatsByIndexSizeDesc) Less(i, j int) bool {
	return c[i].TotalIndexSize > c[j].TotalIndexSize
}

type SortCollectionStatsByIndexSizeAsc []CollectionStats

func (c SortCollectionStatsByIndexSizeAsc) Len() int {
	return len(c)
}

func (c SortCollectionStatsByIndexSizeAsc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortCollectionStatsByIndexSizeAsc) Less(i, j int) bool {
	return c[i].TotalIndexSize < c[j].TotalIndexSize
}

type SortCollectionStatsBySizeDesc []CollectionStats

func (c SortCollectionStatsBySizeDesc) Len() int {
	return len(c)
}

func (c SortCollectionStatsBySizeDesc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortCollectionStatsBySizeDesc) Less(i, j int) bool {
	return c[i].Size > c[j].Size
}

type SortCollectionStatsBySizeAsc []CollectionStats

func (c SortCollectionStatsBySizeAsc) Len() int {
	return len(c)
}

func (c SortCollectionStatsBySizeAsc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c SortCollectionStatsBySizeAsc) Less(i, j int) bool {
	return c[i].Size < c[j].Size
}
