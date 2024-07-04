
type CharHeap []CharFreq

type CharFreq struct {
	char  byte
	count int
}

func (h CharHeap) Len() int            { return len(h) }
func (h CharHeap) Less(i, j int) bool  { return h[i].char > h[j].char }
func (h CharHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *CharHeap) Push(x interface{}) { *h = append(*h, x.(CharFreq)) }
func (h *CharHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func repeatLimitedString(s string, repeatLimit int) string {
    freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	h := &CharHeap{}
	heap.Init(h)
	for char, count := range freq {
		heap.Push(h, CharFreq{char: char, count: count})
	}

	result := []byte{}

	for h.Len() > 0 {
		current := heap.Pop(h).(CharFreq)

		if len(result) > 0 && result[len(result)-1] == current.char {
			if h.Len() == 0 {
				break
			}

			next := heap.Pop(h).(CharFreq)
			result = append(result, next.char)
			next.count--

			if next.count > 0 {
				heap.Push(h, next)
			}
			heap.Push(h, current)
		} else {
			count := min(current.count, repeatLimit)
			for i := 0; i < count; i++ {
				result = append(result, current.char)
			}
			current.count -= count

			if current.count > 0 {
				heap.Push(h, current)
			}
		}
	}

	return string(result)
}