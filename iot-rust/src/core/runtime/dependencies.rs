use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;

pub struct Dependencies<'a> {
    pub config_store: &'a dyn ConfigStorage,
    pub device_id_generator: &'a dyn DeviceIdGenerator,
}

impl<'a> Dependencies<'a>  {
    pub fn new(config_store: &'a dyn ConfigStorage, device_id_generator: &'a dyn DeviceIdGenerator) -> Self {
        Self { config_store, device_id_generator }
    }
}
