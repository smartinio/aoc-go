package utils

import (
	"iter"
	"maps"
)

type Set[E comparable] struct {
	items map[E]struct{}
}

func NewSet[E comparable](initial ...E) *Set[E] {
	this := &Set[E]{items: map[E]struct{}{}}
	for _, item := range initial {
		this.Add(item)
	}
	return this
}

func (this *Set[E]) Values() iter.Seq[E] {
	return maps.Keys(this.items)
}

func (this *Set[E]) Size() int {
	return len(this.items)
}

func (this *Set[E]) Has(value E) bool {
	_, ok := this.items[value]
	return ok
}

func (this *Set[E]) Add(values ...E) *Set[E] {
	for _, value := range values {
		this.items[value] = struct{}{}
	}
	return this
}

func (this *Set[E]) Remove(value E) *Set[E] {
	delete(this.items, value)
	return this
}

func (this *Set[E]) Clear() *Set[E] {
	this.items = map[E]struct{}{}
	return this
}

func (this *Set[E]) Union(other Set[E]) Set[E] {
	new := Set[E]{}
	for a := range this.items {
		new.Add(a)
	}
	for b := range other.items {
		new.Add(b)
	}
	return new
}

func (this *Set[E]) Intersection(other Set[E]) Set[E] {
	new := Set[E]{}
	for a := range this.items {
		if other.Has(a) {
			new.Add(a)
		}
	}
	for b := range other.items {
		if this.Has(b) {
			new.Add(b)
		}
	}
	return new
}
