use core::cell::RefCell;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;
#[cfg(test)]

pub struct MockDeviceIdGenerator<'a> {
    device_id: RefCell<&'a str >,
}

impl<'a> MockDeviceIdGenerator<'a> {
    pub fn new() -> Self {
        Self { device_id: RefCell::new("")  }
    }

    pub fn will_generate_device_id(&self, device_id: &'a str) {
        *self.device_id.borrow_mut() = device_id;
    }
}

impl<'a> DeviceIdGenerator<'a> for MockDeviceIdGenerator<'a> {
    fn generate(&self) -> &'a str {
        *self.device_id.borrow()
    }
}