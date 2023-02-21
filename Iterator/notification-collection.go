package main

// Implements collection interface
type NotificationCollection struct {
	Collection
	Notifications []Notification
}

func (nc NotificationCollection) createIterator() Iterator {
	return &NotificationIterator{position: 0, NotificationList: nc.Notifications}
}
