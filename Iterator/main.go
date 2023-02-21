package main

import "fmt"

/*
Iterator pattern

Type: behavioral

Idea: abstract iteration over an object via iterator interface that has `.createIterator()` method that handles the iteration over a list of objects.

As an example we have a notification object. We also have a collection of notifications (basically, a list).
We implement iterator using an `Iterator` interface to implement it with `NotificationIterator` for iterating over a list of notifications
`Collection` interface is used to specify a collection (list of objects) that is able to create its own iterator (1)
*/

// Interface for iterating over an object
type Iterator interface {
	hasNext() bool
	next() interface{}
}

// (1)
type Collection interface {
	createIterator() Iterator
}

func main() {
	listOfNotifications := []Notification{
		{Noti: "Hello"},
		{Noti: "World"},
		{Noti: "!"},
	}
	collection := NotificationCollection{Notifications: listOfNotifications}
	iter := collection.createIterator()

	fmt.Println("These are current notifications:")
	for iter.hasNext() {
		notification := iter.next().(Notification)
		fmt.Println(notification.Noti)
	}
}
