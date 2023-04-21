mod abstract_factory;
mod default_factories;
mod car;

fn main() {
    let car_factory: abstract_factory::CarFactory = abstract_factory::CarFactory {};
    let car: Box<dyn car::DefaultCar> = car_factory.create_car(car::CarType::MICRO);
}