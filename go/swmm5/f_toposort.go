// Generated from EPA SWMM 5.2.4 (toposort.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var InDegree []int32
var StartPos []int32
var AdjList []int32
var Stack []int32
var First int32
var Last int32
var Examined R.CPtr
var InTree R.CPtr
var LoopLinks []int32
var LoopLinksLast int32

// toposort.c:57
func toposort_sortLinks(sortedLinks []int32) {
	var (
		i   int32
		n_2 int32
	)
	_, _ = i, n_2
	i = 0
	n_2 = 0
	for i = 0; i < Nobjects[LINK]; i++ {
		sortedLinks[i] = i
	}
	if RouteModel == DW {
		checkDummyLinks()
		if ErrorCode != 0 {
			return
		}
		for i = 0; i < Nobjects[NODE]; i++ {
			Node[i].degree = 0
		}
		for i = 0; i < Nobjects[LINK]; i++ {
			n_2 = Link[i].node1
			if Link[i].direction < 0 {
				n_2 = Link[i].node2
			}
			if Node[n_2].type_ == OUTFALL {
				if Link[i].direction < 0 {
					n_2 = Link[i].node1
				} else {
					n_2 = Link[i].node2
				}
				Node[n_2].degree++
			} else {
				Node[n_2].degree++
			}
		}
		return
	}
	if ErrorCode != 0 {
		return
	}
	InDegree = make([]int32, int(uint64(Nobjects[NODE])))
	StartPos = make([]int32, int(uint64(Nobjects[NODE])))
	AdjList = make([]int32, int(uint64(Nobjects[LINK])))
	Stack = make([]int32, int(uint64(Nobjects[NODE])))
	if (((InDegree == nil) || (StartPos == nil)) || (AdjList == nil)) || (Stack == nil) {
		report_writeErrorMsg(ERR_MEMORY, L0)
	} else {
		createAdjList(DIRECTED)
		adjustAdjList()
		for i = 0; i < Nobjects[NODE]; i++ {
			InDegree[i] = 0
		}
		for i = 0; i < Nobjects[LINK]; i++ {
			InDegree[Link[i].node2]++
		}
		n_2 = topoSort(sortedLinks)
	}
	{
		if InDegree != nil {
			_ = InDegree
			InDegree = nil
		}
	}
	{
		if StartPos != nil {
			_ = StartPos
			StartPos = nil
		}
	}
	{
		if AdjList != nil {
			_ = AdjList
			AdjList = nil
		}
	}
	{
		if Stack != nil {
			_ = Stack
			Stack = nil
		}
	}
	if (!(ErrorCode != 0)) && (n_2 != Nobjects[LINK]) {
		report_writeErrorMsg(ERR_LOOP, L0)
		findCycles()
	}
}

// toposort.c:142
func createAdjList(listType int32) {
	var (
		i int32
		j int32
		k int32
	)
	_, _, _ = i, j, k
	i = 0
	j = 0
	k = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		Node[i].degree = 0
	}
	for j = 0; j < Nobjects[LINK]; j++ {
		Node[Link[j].node1].degree++
		if listType == UNDIRECTED {
			Node[Link[j].node2].degree++
		}
	}
	StartPos[0] = 0
	for i = 0; i < (Nobjects[NODE] - 1); i++ {
		StartPos[(i + 1)] = (StartPos[i] + Node[i].degree)
		Node[i].degree = 0
	}
	(Node[(Nobjects[NODE] - 1)]).degree = 0
	for j = 0; j < Nobjects[LINK]; j++ {
		i = Link[j].node1
		k = (StartPos[i] + Node[i].degree)
		AdjList[k] = j
		Node[i].degree++
		if listType == UNDIRECTED {
			i = Link[j].node2
			k = (StartPos[i] + Node[i].degree)
			AdjList[k] = j
			Node[i].degree++
		}
	}
}

// toposort.c:193
func adjustAdjList() {
	var (
		i int32
		j int32
		k int32
		m int32
	)
	_, _, _, _ = i, j, k, m
	i = 0
	j = 0
	k = 0
	m = 0
	for i = 0; i < Nobjects[NODE]; i++ {
		if Node[i].type_ != DIVIDER {
			continue
		}
		if Node[i].degree != 2 {
			continue
		}
		k = Node[i].subIndex
		m = StartPos[i]
		j = AdjList[m]
		if j == Divider[k].link {
			AdjList[m] = AdjList[(m + 1)]
			AdjList[(m + 1)] = j
		}
	}
}

// toposort.c:225
func topoSort(sortedLinks []int32) int32 {
	var (
		i   int32
		j   int32
		k   int32
		n_2 int32
		i1  int32
		i2  int32
		k1  int32
		k2  int32
	)
	_, _, _, _, _, _, _, _ = i, j, k, n_2, i1, i2, k1, k2
	i = 0
	j = 0
	k = 0
	n_2 = 0
	i1 = 0
	i2 = 0
	k1 = 0
	k2 = 0
	First = 0
	Last = (-1)
	for i = 0; i < Nobjects[NODE]; i++ {
		if InDegree[i] == 0 {
			Last++
			Stack[Last] = i
		}
	}
	n_2 = 0
	for First <= Last {
		i1 = Stack[First]
		k1 = StartPos[i1]
		k2 = (k1 + Node[i1].degree)
		for k = k1; k < k2; k++ {
			j = AdjList[k]
			sortedLinks[n_2] = j
			n_2++
			i2 = Link[j].node2
			InDegree[i2]--
			if InDegree[i2] == 0 {
				Last++
				Stack[Last] = i2
			}
		}
		First++
	}
	return n_2
}

// toposort.c:284
func findCycles() {
	var (
		i int32
	)
	_ = i
	i = 0
	AdjList = make([]int32, int((2 * uint64(Nobjects[LINK]))))
	StartPos = make([]int32, int(uint64(Nobjects[NODE])))
	Stack = make([]int32, int(uint64(Nobjects[NODE])))
	Examined = R.Buf(int(uint64(Nobjects[NODE])))
	InTree = R.Buf(int(uint64(Nobjects[LINK])))
	LoopLinks = make([]int32, int(uint64(Nobjects[LINK])))
	if (((((StartPos != nil) && (AdjList != nil)) && (Stack != nil)) && (!Examined.IsNull())) && (!InTree.IsNull())) && (LoopLinks != nil) {
		createAdjList(UNDIRECTED)
		for i = 0; i < Nobjects[NODE]; i++ {
			Examined.S(i, 0)
		}
		for i = 0; i < Nobjects[LINK]; i++ {
			InTree.S(i, 0)
		}
		for i = 0; i < Nobjects[NODE]; i++ {
			if Examined.G(i) != 0 {
				continue
			}
			Last = (-1)
			findSpanningTree(i)
		}
	}
	{
		if StartPos != nil {
			_ = StartPos
			StartPos = nil
		}
	}
	{
		if AdjList != nil {
			_ = AdjList
			AdjList = nil
		}
	}
	{
		if Stack != nil {
			_ = Stack
			Stack = nil
		}
	}
	{
		if !Examined.IsNull() {
			_ = Examined
			Examined = (R.CPtr{})
		}
	}
	{
		if !InTree.IsNull() {
			_ = InTree
			InTree = (R.CPtr{})
		}
	}
	{
		if LoopLinks != nil {
			_ = LoopLinks
			LoopLinks = nil
		}
	}
}

// toposort.c:330
func findSpanningTree(startNode int32) {
	var (
		nextNode int32
		j        int32
		k        int32
		m        int32
	)
	_, _, _, _ = nextNode, j, k, m
	nextNode = 0
	j = 0
	k = 0
	m = 0
	for m = StartPos[startNode]; m < (StartPos[startNode] + Node[startNode].degree); m++ {
		k = AdjList[m]
		if Link[k].node1 == startNode {
			j = Link[k].node2
		} else {
			j = Link[k].node1
		}
		if InTree.G(k) == 0 {
			if Examined.G(j) != 0 {
				InTree.S(k, 2)
				evalLoop(k)
			} else {
				Examined.S(j, 1)
				Last++
				Stack[Last] = j
				InTree.S(k, 1)
			}
		}
	}
	if Last >= 0 {
		nextNode = Stack[Last]
		Last--
		findSpanningTree(nextNode)
	}
}

// toposort.c:385
func evalLoop(startLink int32) {
	var (
		i       int32
		i1      int32
		i2      int32
		j       int32
		kount   int32
		isCycle int32
	)
	_, _, _, _, _, _ = i, i1, i2, j, kount, isCycle
	i = 0
	i1 = 0
	i2 = 0
	j = 0
	kount = 0
	isCycle = 0
	LoopLinksLast = 0
	LoopLinks[0] = startLink
	i1 = Link[startLink].node1
	i2 = Link[startLink].node2
	if !(traceLoop(i1, i2, startLink) != 0) {
		return
	}
	isCycle = 1
	j = LoopLinks[0]
	i2 = Link[j].node2
	for i = 1; i <= LoopLinksLast; i++ {
		j = LoopLinks[i]
		i1 = Link[j].node1
		if i1 == i2 {
			i2 = Link[j].node2
		} else {
			isCycle = 0
			break
		}
	}
	if isCycle != 0 {
		kount = 0
		for i = 0; i <= LoopLinksLast; i++ {
			if (kount % 5) == 0 {
				R.Fprintf(Frpt.file, L276)
			}
			kount++
			R.Fprintf(Frpt.file, L651, Link[LoopLinks[i]].ID)
			if i < LoopLinksLast {
				R.Fprintf(Frpt.file, L652)
			}
		}
	}
}

// toposort.c:440
func traceLoop(i1 int32, i2 int32, k1 int32) int32 {
	var (
		j int32
		k int32
		m int32
	)
	_, _, _ = j, k, m
	j = 0
	k = 0
	m = 0
	if i1 == i2 {
		return 1
	}
	for m = StartPos[i1]; m < (StartPos[i1] + Node[i1].degree); m++ {
		k = AdjList[m]
		if (k == k1) || (InTree.G(k) != 1) {
			continue
		}
		if Link[k].node1 == i1 {
			j = Link[k].node2
		} else {
			j = Link[k].node1
		}
		if traceLoop(j, i2, k) != 0 {
			LoopLinksLast++
			LoopLinks[LoopLinksLast] = k
			return 1
		}
	}
	return 0
}

// toposort.c:482
func checkDummyLinks() {
	var (
		i      int32
		j      int32
		marked []int32
	)
	_, _, _ = i, j, marked
	i = 0
	j = 0
	marked = nil
	marked = make([]int32, int(uint64(Nobjects[NODE])))
	if marked == nil {
		report_writeErrorMsg(ERR_MEMORY, L0)
		return
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		j = Link[i].node2
		if Link[i].direction < 0 {
			j = Link[i].node1
		}
		if ((Link[i].type_ == CONDUIT) && (Link[i].xsect.type_ == DUMMY)) || ((Link[i].type_ == PUMP) && (Pump[Link[i].subIndex].type_ == IDEAL_PUMP)) {
			if marked[j] == 0 {
				marked[j] = 1
			}
		} else {
			marked[j] = (-1)
		}
	}
	for i = 0; i < Nobjects[LINK]; i++ {
		if ((Link[i].type_ == CONDUIT) && (Link[i].xsect.type_ == DUMMY)) || ((Link[i].type_ == PUMP) && (Pump[Link[i].subIndex].type_ == IDEAL_PUMP)) {
			j = Link[i].node1
			if marked[j] > 0 {
				report_writeErrorMsg(ERR_DUMMY_LINK, Node[j].ID)
			}
		}
	}
	{
		if marked != nil {
			_ = marked
			marked = nil
		}
	}
}
