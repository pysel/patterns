'''
Singletone is a software design pattern used when your application needs to make sure that it has only one copy of
a specific object instantiated. Ex: you only need to have one DB object at any time

Link: https://www.geeksforgeeks.org/singleton-design-pattern/

Implementation in this example is not the best because it does not take into account the problem of asynchrony. 
That is, if two threads enter the <method> at an unfortunate same time, they might both get a distinct object,
hence, the property of having only one object is violated.

The reason for that is that it is somewhat complex to implement synchronized methods in python and I do not
want to spend much time on a Singltone lol.

The best approach can be found in the Link above on Java.
'''

if __name__ == "__main__":
    pass