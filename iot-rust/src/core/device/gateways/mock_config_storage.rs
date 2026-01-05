use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::config_storage::ConfigStorageError;
use core::cell::RefCell;

pub struct MockConfigStorage {
    pub device_id: RefCell<&'static str >,
}

impl MockConfigStorage {
    pub fn new() -> Self {
       Self { device_id: RefCell::new("")  }
    }
}

impl<'a> ConfigStorage for MockConfigStorage {
    fn save_device_id(&self, device_id: &'static str) -> Result<(), ConfigStorageError> {
        *self.device_id.borrow_mut() = device_id;
        Ok(())
    }
    fn get_device_id(&self) -> Result<&'static str, ConfigStorageError> {
        Ok(&self.device_id.borrow())
    }
}
