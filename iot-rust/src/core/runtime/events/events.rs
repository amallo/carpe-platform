use crate::core::device::commands::events::DeviceEvent;
use crate::core::device::commands::events::DeviceCommand;
pub enum Event {
    PowerOn,
    DeviceEvent(DeviceEvent),
}

pub enum Command {
    DeviceCommand(DeviceCommand),
}