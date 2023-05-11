from abc import ABC, abstractmethod

class HouseBuildingTemplate(ABC):
    def buildHouse(self): # Template method, non-overridable (at least should be)
        self.buildFoundation()
        self.buildPillars()
        self.buildWalls()
        self.buildWindows()
        print("House is built.")
    
    @abstractmethod
    def buildFoundation(self): pass
    
    @abstractmethod
    def buildPillars(self): pass
    
    @abstractmethod
    def buildWalls(self): pass
    
    @abstractmethod
    def buildWindows(self): pass

# Child subclasses that override abstract methods:

class WoodenHouse(HouseBuildingTemplate):
    def buildFoundation(self):
        print("Building foundation with wood, iron rods and cement.")
    
    def buildPillars(self):
        print("Building pillars with Wood coating.")
    
    def buildWalls(self):
        print("Building walls with Wood coating.")
    
    def buildWindows(self):
        print("Building windows with Wood coating.")

class StoneHouse(HouseBuildingTemplate):
    def buildFoundation(self):
        print("Building foundation with cement, iron rods and sand.")
    
    def buildPillars(self):
        print("Building pillars with cement coating.")
    
    def buildWalls(self):
        print("Building walls with cement coating.")
    
    def buildWindows(self):
        print("Building windows with cement coating.")
    
if __name__ == "__main__":
    wood = WoodenHouse()
    wood.buildHouse()
    
    print()
    
    stone = StoneHouse()
    stone.buildHouse()