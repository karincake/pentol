package pentol

func Bool(i ...bool) *bool {
	var x bool
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func String(i ...string) *string {
	var x string
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Int(i ...int) *int {
	var x int
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Int8(i ...int8) *int8 {
	var x int8
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Int16(i ...int16) *int16 {
	var x int16
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Int32(i ...int32) *int32 {
	var x int32
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Int64(i ...int64) *int64 {
	var x int64
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func UInt(i ...uint) *uint {
	var x uint
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func UInt8(i ...uint8) *uint8 {
	var x uint8
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func UInt16(i ...uint16) *uint16 {
	var x uint16
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func UInt32(i ...uint32) *uint32 {
	var x uint32
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func UInt64(i ...uint64) *uint64 {
	var x uint64
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Float32(i ...float32) *float32 {
	var x float32
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}

func Float64(i ...float64) *float64 {
	var x float64
	if len(i) > 0 {
		x = i[0]
	}
	return &x
}
