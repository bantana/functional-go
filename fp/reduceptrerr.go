package fp

// ReduceIntPtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*int, *int)
//	B. list of type []*int
// 	C. initializer (optional of type int)
//
// Returns:
//	single value, error: (*int, error)

func ReduceIntPtrErr(f func(*int, *int) (*int, error), list []*int, initializer ...int) (*int, error) {
	var initVal int
	var init *int = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceIntPtrErr(f, list[1:], *r)
}

// ReduceInt64PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*int64, *int64)
//	B. list of type []*int64
// 	C. initializer (optional of type int64)
//
// Returns:
//	single value, error: (*int64, error)

func ReduceInt64PtrErr(f func(*int64, *int64) (*int64, error), list []*int64, initializer ...int64) (*int64, error) {
	var initVal int64
	var init *int64 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceInt64PtrErr(f, list[1:], *r)
}

// ReduceInt32PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*int32, *int32)
//	B. list of type []*int32
// 	C. initializer (optional of type int32)
//
// Returns:
//	single value, error: (*int32, error)

func ReduceInt32PtrErr(f func(*int32, *int32) (*int32, error), list []*int32, initializer ...int32) (*int32, error) {
	var initVal int32
	var init *int32 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceInt32PtrErr(f, list[1:], *r)
}

// ReduceInt16PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*int16, *int16)
//	B. list of type []*int16
// 	C. initializer (optional of type int16)
//
// Returns:
//	single value, error: (*int16, error)

func ReduceInt16PtrErr(f func(*int16, *int16) (*int16, error), list []*int16, initializer ...int16) (*int16, error) {
	var initVal int16
	var init *int16 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceInt16PtrErr(f, list[1:], *r)
}

// ReduceInt8PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*int8, *int8)
//	B. list of type []*int8
// 	C. initializer (optional of type int8)
//
// Returns:
//	single value, error: (*int8, error)

func ReduceInt8PtrErr(f func(*int8, *int8) (*int8, error), list []*int8, initializer ...int8) (*int8, error) {
	var initVal int8
	var init *int8 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceInt8PtrErr(f, list[1:], *r)
}

// ReduceUintPtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*uint, *uint)
//	B. list of type []*uint
// 	C. initializer (optional of type uint)
//
// Returns:
//	single value, error: (*uint, error)

func ReduceUintPtrErr(f func(*uint, *uint) (*uint, error), list []*uint, initializer ...uint) (*uint, error) {
	var initVal uint
	var init *uint = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceUintPtrErr(f, list[1:], *r)
}

// ReduceUint64PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*uint64, *uint64)
//	B. list of type []*uint64
// 	C. initializer (optional of type uint64)
//
// Returns:
//	single value, error: (*uint64, error)

func ReduceUint64PtrErr(f func(*uint64, *uint64) (*uint64, error), list []*uint64, initializer ...uint64) (*uint64, error) {
	var initVal uint64
	var init *uint64 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceUint64PtrErr(f, list[1:], *r)
}

// ReduceUint32PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*uint32, *uint32)
//	B. list of type []*uint32
// 	C. initializer (optional of type uint32)
//
// Returns:
//	single value, error: (*uint32, error)

func ReduceUint32PtrErr(f func(*uint32, *uint32) (*uint32, error), list []*uint32, initializer ...uint32) (*uint32, error) {
	var initVal uint32
	var init *uint32 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceUint32PtrErr(f, list[1:], *r)
}

// ReduceUint16PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*uint16, *uint16)
//	B. list of type []*uint16
// 	C. initializer (optional of type uint16)
//
// Returns:
//	single value, error: (*uint16, error)

func ReduceUint16PtrErr(f func(*uint16, *uint16) (*uint16, error), list []*uint16, initializer ...uint16) (*uint16, error) {
	var initVal uint16
	var init *uint16 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceUint16PtrErr(f, list[1:], *r)
}

// ReduceUint8PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*uint8, *uint8)
//	B. list of type []*uint8
// 	C. initializer (optional of type uint8)
//
// Returns:
//	single value, error: (*uint8, error)

func ReduceUint8PtrErr(f func(*uint8, *uint8) (*uint8, error), list []*uint8, initializer ...uint8) (*uint8, error) {
	var initVal uint8
	var init *uint8 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceUint8PtrErr(f, list[1:], *r)
}

// ReduceStrPtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*string, *string)
//	B. list of type []*string
// 	C. initializer (optional of type string)
//
// Returns:
//	single value, error: (*string, error)

func ReduceStrPtrErr(f func(*string, *string) (*string, error), list []*string, initializer ...string) (*string, error) {
	var initVal string
	var init *string = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceStrPtrErr(f, list[1:], *r)
}

// ReduceFloat32PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*float32, *float32)
//	B. list of type []*float32
// 	C. initializer (optional of type float32)
//
// Returns:
//	single value, error: (*float32, error)

func ReduceFloat32PtrErr(f func(*float32, *float32) (*float32, error), list []*float32, initializer ...float32) (*float32, error) {
	var initVal float32
	var init *float32 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceFloat32PtrErr(f, list[1:], *r)
}

// ReduceFloat64PtrErr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments (*float64, *float64)
//	B. list of type []*float64
// 	C. initializer (optional of type float64)
//
// Returns:
//	single value, error: (*float64, error)

func ReduceFloat64PtrErr(f func(*float64, *float64) (*float64, error), list []*float64, initializer ...float64) (*float64, error) {
	var initVal float64
	var init *float64 = &initVal
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0], nil
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init, nil
	}

	r, err := f(init, list[0])
	if err != nil {
		return nil, err
	}
	return ReduceFloat64PtrErr(f, list[1:], *r)
}
