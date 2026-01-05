#[derive(Debug)]
#[derive(PartialEq)]
pub enum ConfigStorageError {
    SaveError,
}

pub trait ConfigStorage {
    fn save_device_id(&self, device_id: &'static str) -> Result<(), ConfigStorageError>;
    fn get_device_id(&self) -> Result<&'static str, ConfigStorageError>;
}
