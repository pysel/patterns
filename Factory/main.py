'''
Type: Creational 

The idea of Factory pattern is to create a helper function, which
will construct different class types based on user input.

Link: https://www.geeksforgeeks.org/factory-method-for-designing-pattern/?ref=lbp

Benefit of Factory pattern is loose coupling, because you leave only a type selection
for a user, leaving object creation a library's problem.
'''


class ArchLinux:
	def __init__(self):
		self.type = "arch"
	
	def printType(self):
		print(self.type)

class OSX:
	def __init__(self):
		self.type = "OSX"
	
	def printType(self):
		print(self.type)

class Windows:
	def __init__(self):
		self.type = "Windows"
	
	def printType(self):
		print(self.type)

# This is a factory method
def NewOS(type: str):
    if type == "windows":
        return Windows()
    elif type == "osx":
        return OSX()
    elif type == "arch":
        return ArchLinux()
    else:
        raise "type not supported"
		
if __name__ == "__main__":
    myNewOS = NewOS("osx")
    myNewOS.printType()