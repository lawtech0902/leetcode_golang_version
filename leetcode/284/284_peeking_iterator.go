/*
__author__ = 'lawtech'
__date__ = '2020/04/16 2:51 下午'
*/

package _284

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

// random
type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	return false
}

func (this *Iterator) next() int {
	return 0
}

type PeekingIterator struct {
	nextIter *Iterator
	val      interface{}
	has      bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	if iter.hasNext() {
		tmp := iter.next()
		return &PeekingIterator{
			nextIter: iter,
			val:      tmp,
			has:      true,
		}
	} else {
		return &PeekingIterator{
			nextIter: iter,
			val:      nil,
			has:      false,
		}
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.has
}

func (this *PeekingIterator) next() int {
	if this.has {
		tmp := this.val
		this.has = this.nextIter.hasNext()
		this.val = this.nextIter.next()
		return tmp.(int)
	} else {
		return this.val.(int)
	}
}

func (this *PeekingIterator) peek() int {
	return this.val.(int)
}
