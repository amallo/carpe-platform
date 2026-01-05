use crate::core::runtime::command_handler::{CommandHandler, CommandContext};
use crate::core::runtime::events::events::Event;
use crate::core::device::commands::events::DeviceEvent;

pub struct SetupDeviceCommandHandler;

impl SetupDeviceCommandHandler {
    pub fn new() -> Self {
        Self
    }
}

impl CommandHandler for SetupDeviceCommandHandler {
    async fn execute(&self, ctx: &mut CommandContext<'_>) -> Event {
        Event::DeviceEvent(DeviceEvent::DeviceHasBeenSetup("DEV_001"))
    }
}