use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;
pub struct BootDeviceUseCase<'a> {
    pub config_store: &'a dyn ConfigStorage<'a>,
    pub device_id_generator: &'a dyn DeviceIdGenerator<'a>,
}

#[derive(Debug)]
pub struct Success<'a> {
    pub device_id: &'a str ,
}

#[derive(Debug)]
pub enum Err {
    None,
}

impl<'a>  BootDeviceUseCase<'a>  {
    pub fn new(config_store: &'a impl ConfigStorage<'a>, device_id_generator: &'a impl DeviceIdGenerator<'a>) -> Self {
        Self { config_store, device_id_generator }
    }

    pub fn execute(&self) -> Result<(), Err> {
        let _ = self.config_store.save_device_id(self.device_id_generator.generate());
        Ok(())
    }
}