package response

func mapper[T any, U any](ss []T, fn func(T) U) (ss2 []U) {
	if ss == nil {
		return []U{}
	}

	ss2 = make([]U, len(ss))
	for i, s := range ss {
		ss2[i] = fn(s)
	}

	return
}
