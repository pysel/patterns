'''
Observer Pattern is useful when exists a dependency between a single object (subject) and many listeners (observers)
aka one to many relationship. Observers can "subscribe" and "unsubscribe" to/from messages from a subject (which sometimes
leads to memory leaks - Labsed Listener Problem). 

When data changes on the subject, all the observers automatically get notified (ex: weather forecast)
'''
# TODO
