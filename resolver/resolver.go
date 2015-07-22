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

func (r *Resolver) resolve(item *element) (err error) {
	keys := keyReg.FindAllString(*item.Val, -1)

	for _, keyFull := range keys {
		key := keyFull[2 : len(keyFull)-1]

		val, ok := r.data[key]
		if !ok {
			err = &ResolveError{
				errors.Newf(`resolver: Failed to resolve '%s' in '%s="%s"'`,
					keyFull, item.Key, *item.Val),
			}
			return
		}

		*item.Val = strings.Replace(*item.Val, keyFull, val, 1)
	}

	r.data[item.Key] = *item.Val

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

		err = r.resolve(item)
		if err != nil {
			if item.Count > 10 {
				return
			} else {
				err = nil
				r.queue.PushBack(elem)
			}
		} else {
			r.queue.Remove(elem)
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
