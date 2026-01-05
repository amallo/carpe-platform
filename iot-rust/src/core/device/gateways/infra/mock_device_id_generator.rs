use core::cell::RefCell;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;
#[cfg(test)]

pub struct MockDeviceIdGenerator {
    device_id: RefCell<&'static str >,
}

impl MockDeviceIdGenerator {
    pub fn new() -> Self {
        Self { device_id: RefCell::new("")  }
    }

    pub fn will_generate_device_id(&self, device_id: &'static str) {
        *self.device_id.borrow_mut() = device_id;
    }
}

impl DeviceIdGenerator for MockDeviceIdGenerator {
    fn generate(&self) -> &'static str {
        *self.device_id.borrow()
    }
}

