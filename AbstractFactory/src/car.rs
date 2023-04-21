#[derive(Debug)]
pub enum Location {
    USA,
    RUSSIA,
}

#[derive(Debug)]
pub enum CarType {
    MICRO,
    RV
}

#[derive(Debug)]
pub struct Car {
    model: CarType, // CarType
    location: Location,
}



pub fn create_car(model: CarType, location: Location) {

}

trait DefaultCar {
    fn set_model(self, model: CarType);
    fn set_location(self, location: Location);
    fn explain_car(&self);
}

impl DefaultCar for Car {
    fn set_model(mut self, model: CarType) {
        self.model = model;
    }
    fn set_location(mut self, location: Location) {
        self.location = location;
    }
    fn explain_car(&self) {
        println!("Default. Model: {:?}, Location: {:?}", self.model, self.location);
    }
}

trait RVCar: DefaultCar {
    fn explain_car(&self);
}

impl RVCar for Car {
    fn explain_car(&self) {
        println!("RV. Model: {:?}, Location: {:?}", self.model, self.location);
    }
}

trait MicroCar: DefaultCar {
    fn explain_car(&self);
}

impl MicroCar for Car {
    fn explain_car(&self) {
        println!("Micro. Model: {:?}, Location: {:?}", self.model, self.location);
    }
}
