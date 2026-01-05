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
        let device_id = match ctx.deps.config_store.get_device_id() {
            Ok(existing_id) if !existing_id.is_empty() => existing_id,
            _ => {
                let new_device_id = ctx.deps.device_id_generator.generate();
                ctx.deps.config_store.save_device_id(new_device_id).ok();
                new_device_id
            }
        };
        Event::DeviceEvent(DeviceEvent::DeviceHasBeenSetup(device_id))
    }
}