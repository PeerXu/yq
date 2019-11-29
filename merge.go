package main

import mergo "github.com/PeerXu/mergo"

func merge(dst interface{}, src interface{}, overwrite bool, append bool, deepcopy bool) error {
	if overwrite {
		return mergo.Merge(dst, src, mergo.WithOverride)
	} else if append {
		return mergo.Merge(dst, src, mergo.WithAppendSlice)
	} else if deepcopy {
		return mergo.Merge(dst, src, mergo.WithSliceDeepCopy)
	}
	return mergo.Merge(dst, src)
}
