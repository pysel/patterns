use super::car;

pub struct USFactory {}
pub struct RussiaFactory {}

pub trait Factory {
    fn create_car(&self, model: car::CarType) -> Box<dyn car::DefaultCar>;
}

impl Factory for USFactory {
    fn create_car(&self, model: car::CarType) -> Box<dyn car::DefaultCar> {
        car::create_car(model, car::Location::USA)
    }
}

impl Factory for RussiaFactory {
    fn create_car(&self, model: car::CarType) -> Box<dyn car::DefaultCar> {
        car::create_car(model, car::Location::RUSSIA)
    }
}