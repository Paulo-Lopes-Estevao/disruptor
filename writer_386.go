package disruptor

func (this *Writer) Commit(lower, upper int64) {
	this.writerCursor.Store(upper)
}
