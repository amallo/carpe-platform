use crate::core::runtime::command_handler::{CommandHandler, CommandContext};
use crate::core::runtime::events::events::Event;
use crate::core::device::commands::events::DeviceEvent;

pub struct SetupDeviceCommandHandler;

impl SetupDeviceCommandHandler {
    pub fn new() -> Self {
        Self
    }
}

impl<'a> CommandHandler<'a> for SetupDeviceCommandHandler {
    async fn execute(&self, ctx: &'a CommandContext<'_>) -> Event {
        let device_id = ctx.deps.device_id_generator.generate();
        Event::DeviceEvent(DeviceEvent::DeviceHasBeenSetup(device_id))
    }
}