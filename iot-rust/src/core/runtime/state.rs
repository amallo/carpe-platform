#[derive(PartialEq)]
pub enum StatusValue {
    Initial,
    Ready,
}

pub struct State {
    status: StatusValue,
}


impl State {
    pub fn new() -> Self {
        Self { status: StatusValue::Initial }
    }
    pub fn is_ready(&self) -> bool {
        self.status == StatusValue::Ready
    }

    pub fn set_ready(&mut self) {
        self.status = StatusValue::Ready;
    }

}