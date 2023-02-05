package usecase

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

const (
	key1 = "testkey01"
	key2 = "testkey02"

	value1 = "testvalue01"
	value2 = "testvalue02"
)

var (
	now       = time.Now()
	timetest1 = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	timetest2 = timetest1.Add(time.Hour)
)

func TestLWW_Remove(t *testing.T) {
	// init test
	dict := make(map[string]LWWValue)
	dict[key1] = LWWValue{
		Value:     value1,
		Timestamp: now,
	}
	dict[key2] = LWWValue{
		Value:     value1,
		Timestamp: now,
	}
	type fields struct {
		dict map[string]LWWValue
		mu   sync.RWMutex
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "remove a dicitonary",
			fields: fields{dict: dict},
			args:   args{key: key1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LWW{
				dict: tt.fields.dict,
			}
			l.Remove(tt.args.key)
			if dict[key1].Value != "" && dict[key1].Timestamp != now {
				t.Errorf("Key1 is not deleted, value = %v, want %v", dict[key1].Value, "string empty")
			}
			if dict[key2].Value == "" && dict[key2].Timestamp == now {
				t.Errorf("Key2 deleted, value = %v, want %v", dict[key1].Value, "string empty")
			}
		})
	}
}

func TestLWW_MergeSameTime(t *testing.T) {
	dict := make(map[string]LWWValue)
	dict[key1] = LWWValue{
		Value:     value1,
		Timestamp: timetest1,
	}
	dict[key2] = LWWValue{
		Value:     value2,
		Timestamp: timetest2,
	}
	type fields struct {
		dict map[string]LWWValue
	}
	type args struct {
		key1 string
		key2 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []LWWResp
	}{
		{
			name:   "merge different timestamp",
			fields: fields{dict: dict},
			args:   args{key1: key1, key2: key2},
			want: []LWWResp{
				{
					Key:       key1,
					Value:     value2,
					Timestamp: timetest2.Format(time.RFC1123),
				},
				{
					Key:       key2,
					Value:     value2,
					Timestamp: timetest2.Format(time.RFC1123),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LWW{
				dict: tt.fields.dict,
			}
			if got := l.Merge(tt.args.key1, tt.args.key2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LWW.Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestLWW_MergeDiffTime(t *testing.T) {
	dict := make(map[string]LWWValue)
	dict[key1] = LWWValue{
		Value:     value1,
		Timestamp: timetest1,
	}
	dict[key2] = LWWValue{
		Value:     value2,
		Timestamp: timetest2,
	}
	type fields struct {
		dict map[string]LWWValue
	}
	type args struct {
		key1 string
		key2 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []LWWResp
	}{
		{
			name:   "merge different timestamp",
			fields: fields{dict: dict},
			args:   args{key1: key1, key2: key2},
			want: []LWWResp{
				{
					Key:       key1,
					Value:     value2,
					Timestamp: timetest2.Format(time.RFC1123),
				},
				{
					Key:       key2,
					Value:     value2,
					Timestamp: timetest2.Format(time.RFC1123),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LWW{
				dict: tt.fields.dict,
			}
			if got := l.Merge(tt.args.key1, tt.args.key2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LWW.Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLWW_Lookup(t *testing.T) {
	dict := make(map[string]LWWValue)
	dict[key1] = LWWValue{
		Value:     value1,
		Timestamp: timetest1,
	}
	type fields struct {
		dict map[string]LWWValue
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   LWWResp
	}{
		{
			name:   "success lookup key",
			fields: fields{dict: dict},
			args:   args{key: key1},
			want: LWWResp{
				Key:       key1,
				Value:     value1,
				Timestamp: timetest1.Format(time.RFC1123),
			},
		},
		{
			name:   "lookup key with empty value",
			fields: fields{dict: dict},
			args:   args{key: key2},
			want: LWWResp{
				Key:       key2,
				Value:     "",
				Timestamp: time.Time{}.Format(time.RFC1123),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LWW{
				dict: tt.fields.dict,
			}
			if got := l.Lookup(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LWW.Lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLWW_List(t *testing.T) {
	type fields struct {
		dict map[string]LWWValue
	}
	dict1 := make(map[string]LWWValue)
	dict2 := make(map[string]LWWValue)
	dict1[key1] = LWWValue{
		Value:     value1,
		Timestamp: timetest1,
	}
	dict1[key2] = LWWValue{
		Value:     value2,
		Timestamp: timetest2,
	}
	tests := []struct {
		name   string
		fields fields
		want   []LWWResp
	}{
		{
			name:   "success get list dictionary",
			fields: fields{dict: dict1},
			want: []LWWResp{
				{
					Key:       key1,
					Value:     value1,
					Timestamp: timetest1.Format(time.RFC1123),
				},
				{
					Key:       key2,
					Value:     value2,
					Timestamp: timetest2.Format(time.RFC1123),
				},
			},
		},
		{
			name:   "success dictionary empty",
			fields: fields{dict: dict2},
			want:   []LWWResp{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LWW{
				dict: tt.fields.dict,
			}
			if got := l.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LWW.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLWW_Add(t *testing.T) {
	// test will be see matching value, not timestamp
	dict := make(map[string]LWWValue)
	type fields struct {
		dict map[string]LWWValue
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "add a dicitonary",
			fields: fields{dict: dict},
			args:   args{key: key1, value: value1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LWW{
				dict: tt.fields.dict,
			}
			l.Add(tt.args.key, tt.args.value)
			if dict[key1].Value != value1 {
				t.Errorf("LWW.List() wants %v, got Value %v", value1, dict[key1].Value)

			}
		})
	}
}
