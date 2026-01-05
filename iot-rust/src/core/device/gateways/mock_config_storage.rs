use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::config_storage::ConfigStorageError;
use core::cell::RefCell;

pub struct MockConfigStorage<'a> {
    pub device_id: RefCell<&'a str >,
}

impl<'a> MockConfigStorage<'a> {
    pub fn new() -> Self {
       Self { device_id: RefCell::new("")  }
    }
}

impl<'a> ConfigStorage<'a> for MockConfigStorage<'a> {
    fn save_device_id(&self, device_id: &'a str) -> Result<(), ConfigStorageError> {
        *self.device_id.borrow_mut() = device_id;
        Ok(())
    }
    fn get_device_id(&self) -> Result<&'a str, ConfigStorageError> {
        Ok(&self.device_id.borrow())
    }
}
