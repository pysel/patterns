'''

Strategy pattern 

Type: behavioral ( abstracts common class abilities into other classes )

Strategy pattern encapsulates common algorithm behaviors and the actual behavior is defined at runtime.

In this example we will have cars sharing same methods: drive, display, getting power (gas or elictricity).
Drive and display methods will be the same across all cars (just driving forward and just saying something like "I am a car"),
    hence, we define a parent class for every car. (1)

Problems arise when we need to define getting power method for every new car, because (example) these methods will be different for electric, gasoline and disel cars.
For these problems we will use a strategy pattern (interface), which ensapsulates common behaviors (2)

'''

# (2)
# Lets all think this is an interface, python does not support normal interfaces
# This interface IS a strategy pattern, as it encapsulates common behavior (getting power in this example)
class GetPowerInterface:
    def __init__(self) -> None:
        self.chargedPercent = 0
        
    def getPower(self):
        '''
        I am unimplemented method
        '''
        pass

# (1)
# This class is the parent of every car possible
# (because every car can drive and they do it pretty much the same way)
class Car:
    def __init__(self, powerBehavior: GetPowerInterface) -> None:
        self.powerBehavior = powerBehavior.getPower
        
    def drive(self):
        print("I am driving")
        
    def display(self):
        print("I am a car")
        
    def getPower(self):
        self.powerBehavior()
        

# Gas cars
class Gas(GetPowerInterface):
    def __init__(self) -> None:
        super().__init__()
    
    def getPower(self):
        print("\nUsing pump, pumping gas . . .")
        self.chargedPercent = 100
        print("This gas car has successfully charged, charge percent: ", self.chargedPercent, "%\n")

# Electric cars
class Electricity(GetPowerInterface):
    def __init__(self) -> None:
        super().__init__()
        
    def getPower(self):
        print("\nUsing power outlet, charging . . .")
        self.chargedPercent = 100
        print("This electric vehicle has successfully charged, charge percent: ", self.chargedPercent, "%\n")

        
if __name__ == "__main__":
    # collect behaviors into variables (for simplicity)
    gasBehavior = Gas()
    electricBehavior = Electricity()
    
    # Defining (in runtime) what type of `getPower` objects will have
    tesla = Car(electricBehavior)
    mercedes = Car(gasBehavior)
    
    # run methods
    tesla.getPower()
    mercedes.getPower()
    
    