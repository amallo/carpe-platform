#[derive(Debug)]
#[derive(PartialEq)]
pub enum ConfigStorageError {
    SaveError,
}

pub trait ConfigStorage<'a> {
    fn save_device_id(&self, device_id: &'a str) -> Result<(), ConfigStorageError>;
    fn get_device_id(&self) -> Result<&'a str, ConfigStorageError>;
}
