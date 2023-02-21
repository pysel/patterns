package main

type NotificationIterator struct {
	Iterator
	position         int
	NotificationList []Notification
}

func (ni *NotificationIterator) next() interface{} {
	notification := ni.NotificationList[ni.position]
	ni.position++
	return notification
}

func (ni NotificationIterator) hasNext() bool {
	if ni.position <= len(ni.NotificationList)-1 {
		return true
	}
	return false
}
