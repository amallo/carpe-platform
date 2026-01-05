use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;

pub struct Dependencies<'a, C: ConfigStorage, D: DeviceIdGenerator> {
    pub config_store: &'a C,
    pub device_id_generator: &'a D,
}

impl<'a, C: ConfigStorage, D: DeviceIdGenerator> Dependencies<'a, C, D> {
    pub fn new(config_store: &'a C, device_id_generator: &'a D) -> Self {
        Self { config_store, device_id_generator }
    }
}
