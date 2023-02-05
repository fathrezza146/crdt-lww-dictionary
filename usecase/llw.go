package usecase

import (
	"sort"
	"time"
)

type LWW struct {
	dict map[string]LWWValue
}

type LWWValue struct {
	Value     string
	Timestamp time.Time
}

type LWWList struct {
	List []LWWResp "json:'list'"
}

type LWWResp struct {
	Key       string "json:'element'"
	Value     string "json:'value'"
	Timestamp string "json:'timestamp'"
}

func InitLWW() *LWW {
	return &LWW{
		dict: make(map[string]LWWValue),
	}
}

// remove a dictionary, set to default value
func (l *LWW) Remove(key string) {
	delete(l.dict, key)
}

// Merge 2 keys together. both keys set value to whoever edited last
// If both have the same, will bias towards set 1
// if one of them doesnt have value, it will set to the one with value
func (l *LWW) Merge(key1 string, key2 string) []LWWResp {
	val1 := l.dict[key1]
	val2 := l.dict[key2]

	var mergedata LWWValue
	if val1.Timestamp.Equal(val2.Timestamp) || val1.Timestamp.After(val2.Timestamp) {
		mergedata = LWWValue{
			Value:     val1.Value,
			Timestamp: val1.Timestamp,
		}
		l.dict[key2] = mergedata
	} else {
		mergedata = LWWValue{
			Value:     val2.Value,
			Timestamp: val2.Timestamp,
		}
		l.dict[key1] = mergedata
	}

	var resp []LWWResp

	resp = append(resp, LWWResp{
		Key:       key1,
		Value:     mergedata.Value,
		Timestamp: mergedata.Timestamp.Format(time.RFC1123),
	}, LWWResp{
		Key:       key2,
		Value:     mergedata.Value,
		Timestamp: mergedata.Timestamp.Format(time.RFC1123),
	})

	return resp

}

// lookup a dictionary based on key, no dictionary returns empty state
func (l *LWW) Lookup(key string) LWWResp {
	val := l.dict[key]

	resp := LWWResp{
		Key:       key,
		Value:     val.Value,
		Timestamp: val.Timestamp.Format(time.RFC1123),
	}

	return resp

}

// get all key & value of dictionaries.
// sort by alphabetic
func (l *LWW) List() []LWWResp {
	resp := []LWWResp{}

	for key, value := range l.dict {
		resp = append(resp, LWWResp{
			Key:       key,
			Value:     value.Value,
			Timestamp: value.Timestamp.Format(time.RFC1123),
		})
	}

	sort.Slice(resp, func(i, j int) bool { return resp[i].Key < resp[j].Key })
	return resp
}

// add and update are in the same function, since its always update to the latest addition
func (l *LWW) Add(key, value string) {
	now := time.Now()
	structValue := LWWValue{
		Value:     value,
		Timestamp: now,
	}

	l.dict[key] = structValue
}
