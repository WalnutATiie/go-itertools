package channel

import "reflect"

func takeSliceArg(arg interface{}) (out []interface{}, ok bool) {
	slice, success := takeArg(arg, reflect.Slice)
	if !success {
		ok = false
		return
	}
	c := slice.Len()
	out = make([]interface{}, c)
	for i := 0; i < c; i++ {
		out[i] = slice.Index(i).Interface()
	}
	return out, true
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}

func takeMapArg(arg interface{}) (out map[string]interface{}, ok bool) {
	smap, success := takeArg(arg, reflect.Map)
	if !success {
		ok = false
		return
	}
	out = make(map[string]interface{}, 0)
	for _, i := range smap.MapKeys() {
		out[i.String()] = smap.MapIndex(i).Interface()
	}
	return out, true
}
