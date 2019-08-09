package util

import (
	"context"
	"github.com/searKing/golang/go/error/exception"
	"github.com/searKing/golang/go/util/class"
	"github.com/searKing/golang/go/util/function/consumer"
	"github.com/searKing/golang/go/util/object"
)

/**
 * An iterator over a collection.  {@code Iterator} takes the place of
 * {@link Enumeration} in the Java Collections Framework.  Iterators
 * differ from enumerations in two ways:
 *
 * <ul>
 *      <li> Iterators allow the caller to remove elements from the
 *           underlying collection during the iteration with well-defined
 *           semantics.
 *      <li> Method names have been improved.
 * </ul>
 *
 * <p>This interface is a member of the
 * <a href="{@docRoot}/java/util/package-summary.html#CollectionsFramework">
 * Java Collections Framework</a>.
 *
 * @apiNote
 * An {@link Enumeration} can be converted into an {@code Iterator} by
 * using the {@link Enumeration#asIterator} method.
 *
 * @param <E> the type of elements returned by this iterator
 *
 * @author  Josh Bloch
 * @see Collection
 * @see ListIterator
 * @see Iterable
 * @since 1.2
 */

type Iterator interface {
	/**
	 * Returns {@code true} if the iteration has more elements.
	 * (In other words, returns {@code true} if {@link #next} would
	 * return an element rather than throwing an exception.)
	 *
	 * @return {@code true} if the iteration has more elements
	 */
	HasNext() bool
	/**
	 * Returns the next element in the iteration.
	 *
	 * @return the next element in the iteration
	 * @throws NoSuchElementException if the iteration has no more elements
	 */
	Next() interface{}

	/**
	 * Removes from the underlying collection the last element returned
	 * by this iterator (optional operation).  This method can be called
	 * only once per call to {@link #next}.
	 * <p>
	 * The behavior of an iterator is unspecified if the underlying collection
	 * is modified while the iteration is in progress in any way other than by
	 * calling this method, unless an overriding class has specified a
	 * concurrent modification policy.
	 * <p>
	 * The behavior of an iterator is unspecified if this method is called
	 * after a call to the {@link #forEachRemaining forEachRemaining} method.
	 *
	 * @implSpec
	 * The default implementation throws an instance of
	 * {@link UnsupportedOperationException} and performs no other action.
	 *
	 * @throws UnsupportedOperationException if the {@code remove}
	 *         operation is not supported by this iterator
	 *
	 * @throws IllegalStateException if the {@code next} method has not
	 *         yet been called, or the {@code remove} method has already
	 *         been called after the last call to the {@code next}
	 *         method
	 */
	Remove()

	/**
	 * Performs the given action for each remaining element until all elements
	 * have been processed or the action throws an exception.  Actions are
	 * performed in the order of iteration, if that order is specified.
	 * Exceptions thrown by the action are relayed to the caller.
	 * <p>
	 * The behavior of an iterator is unspecified if the action modifies the
	 * collection in any way (even by calling the {@link #remove remove} method
	 * or other mutator methods of {@code Iterator} subtypes),
	 * unless an overriding class has specified a concurrent modification policy.
	 * <p>
	 * Subsequent behavior of an iterator is unspecified if the action throws an
	 * exception.
	 *
	 * @implSpec
	 * <p>The default implementation behaves as if:
	 * <pre>{@code
	 *     while (hasNext())
	 *         action.accept(next());
	 * }</pre>
	 *
	 * @param action The action to be performed for each element
	 * @throws NullPointerException if the specified action is null
	 * @since 1.8
	 */
	ForEachRemaining(ctx context.Context, action consumer.Consumer)
}

type AbstractIteratorClass struct {
	class.Class
}

func (iter *AbstractIteratorClass) HasNext() bool {
	panic(exception.NewIllegalStateException1("called wrong HasNext method"))
}

func (iter *AbstractIteratorClass) Next() interface{} {
	panic(exception.NewIllegalStateException1("called wrong Next method"))
}

func (iter *AbstractIteratorClass) Remove() {
	panic(exception.NewUnsupportedOperationException1("remove"))
}

func (iter *AbstractIteratorClass) ForEachRemaining(ctx context.Context, action consumer.Consumer) {
	object.RequireNonNull(action)
	for iter.GetDerivedElse(iter).(Iterator).HasNext() {
		select {
		case <-ctx.Done():
			return
		default:
		}
		action.Accept(iter.GetDerivedElse(iter).(Iterator).Next())
	}
}