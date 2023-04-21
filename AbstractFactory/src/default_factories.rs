mod car;

pub struct USFactory {}
pub struct RussiaFactory {}

trait Factory {
    fn create_car(&self, model: CarType) -> car::Car;
}

impl Factory for USFactory {
    fn create_car(&self, model: CarType) -> car::Car {
        car::create_car(model, car::Location::USA)
    }
}

impl Factory for RussiaFactory {
    fn create_car(&self, model: CarType) -> car::Car {
        car::create_car(model, car::Location::RUSSIA)
    }
}