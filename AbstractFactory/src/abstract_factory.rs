use super::car;
use super::default_factories;

pub struct CarFactory;

impl default_factories::Factory for CarFactory {
    fn create_car(&self, model: car::CarType) -> Box<dyn car::DefaultCar> {
        // imagine getting data, for example, from some GPS device
        let location: car::Location = car::Location::USA;  

        match location {
            car::Location::RUSSIA => {
                let factory: default_factories::RussiaFactory = default_factories::RussiaFactory {};
                factory.create_car(model)
            },
            car::Location::USA => {
                let factory: default_factories::USFactory = default_factories::USFactory {};
                factory.create_car(model)
            },
        }
    }
}