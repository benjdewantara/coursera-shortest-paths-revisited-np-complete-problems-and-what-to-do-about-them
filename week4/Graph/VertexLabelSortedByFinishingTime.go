package Graph

type VertexLabelSortedByFinishingTime struct {
	VertexLabels  []int
	FinishingTime []int
}

func (v *VertexLabelSortedByFinishingTime) Len() int {
	return len(v.VertexLabels)
}

func (v *VertexLabelSortedByFinishingTime) Less(i, j int) bool {
	return v.FinishingTime[i] < v.FinishingTime[j]
}

func (v *VertexLabelSortedByFinishingTime) Swap(i, j int) {
	v.VertexLabels[i], v.VertexLabels[j] = v.VertexLabels[j], v.VertexLabels[i]
	v.FinishingTime[i], v.FinishingTime[j] = v.FinishingTime[j], v.FinishingTime[i]
}
