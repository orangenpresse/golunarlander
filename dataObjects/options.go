package dataObjects

type Options struct {
	DebugMode    bool
	BufferWidth  int32
	BufferHeight int32
}

func (o *Options) SetDimensions(width int, height int) {
	o.BufferWidth = int32(width)
	o.BufferHeight = int32(height)
}
