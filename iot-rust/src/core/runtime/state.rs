#[derive(PartialEq)]
pub enum StateValue {
    Initial,
    Ready,
}

pub struct State {
    value: StateValue,
}


impl State {
    pub fn new() -> Self {
        Self { value: StateValue::Initial }
    }
    pub fn is_ready(&self) -> bool {
        self.value == StateValue::Ready
    }

}