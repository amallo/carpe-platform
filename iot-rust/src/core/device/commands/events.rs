
pub enum DeviceEvent {
    DeviceHasBeenSetup(&'static str),
}

pub enum DeviceCommand {
    SetupDevice,
}