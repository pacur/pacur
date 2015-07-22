package resolver

import (
	"container/list"
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"regexp"
	"strings"
)

var (
	keyReg = regexp.MustCompile(`(\$\{[^\}]+\})`)
)

type element struct {
	Key   string
	Val   *string
	Count int
}

type Resolver struct {
	queue *list.List
	data  map[string]string
}

func (r *Resolver) Add(key string, val *string) {
	item := &element{
		Key: key,
		Val: val,
	}
	r.queue.PushBack(item)
}

func (r *Resolver) AddItem(key string, i int, val *string) {
	item := &element{
		Key: fmt.Sprintf("%s[%d]", key, i),
		Val: val,
	}
	r.queue.PushBack(item)
}

func (r *Resolver) AddList(key string, vals []string) {
	for i := 0; i < len(vals); i++ {
		item := &element{
			Key: fmt.Sprintf("%s[%d]", key, i),
			Val: &vals[i],
		}
		r.queue.PushBack(item)
	}
}

func (r *Resolver) resolve(item *element) (status bool, err error) {
	keys := keyReg.FindAllString(*item.Val, -1)

	for _, keyFull := range keys {
		key := keyFull[2 : len(keyFull)-1]

		val, ok := r.data[key]
		if !ok {
			return
		}

		*item.Val = strings.Replace(*item.Val, keyFull, val, 1)
	}

	r.data[item.Key] = *item.Val

	status = true
	return
}

func (r *Resolver) Resolve() (err error) {
	for {
		elem := r.queue.Front()
		if elem == nil {
			return
		}
		item := elem.Value.(*element)
		item.Count += 1

		ok, e := r.resolve(item)
		if e != nil {
			err = e
			return
		}

		if ok {
			r.queue.Remove(elem)
		} else if item.Count > 10 {
			err = errors.Newf("resolver: Failed to resolve '${%s}'", item.Key)
			return
		} else {
			r.queue.PushBack(elem)
		}
	}
	return
}

func New() (reslv *Resolver) {
	reslv = &Resolver{
		queue: list.New(),
		data:  map[string]string{},
	}
	return
}
