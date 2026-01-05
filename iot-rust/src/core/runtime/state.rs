#[derive(PartialEq)]
pub enum StatusValue {
    Initial,
    Ready(&'static str),
}

pub struct State {
    status: StatusValue,
}


impl State {
    pub fn new() -> Self {
        Self { status: StatusValue::Initial }
    }
    pub fn is_ready(&self) -> bool {
        matches!(self.status, StatusValue::Ready(_))
    }

    pub fn set_ready(&mut self, device_id: &'static str) {
        self.status = StatusValue::Ready(device_id);
    }

    pub fn device_id(&self) -> Option<&'static str> {
        match &self.status {
            StatusValue::Ready(id) => Some(id),
            _ => None,
        }
    }

}