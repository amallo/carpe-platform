use crate::core::runtime::command_handler::{CommandHandler, CommandContext};
use crate::core::runtime::events::events::Event;
use crate::core::device::commands::events::DeviceEvent;
use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;

pub struct SetupDeviceCommandHandler;

impl SetupDeviceCommandHandler {
    pub fn new() -> Self {
        Self
    }
}

impl<'a, C: ConfigStorage, D: DeviceIdGenerator> CommandHandler<'a, C, D> for SetupDeviceCommandHandler {
    async fn execute(&self, ctx: &'a CommandContext<'_, C, D>) -> Event {
        let device_id = match ctx.deps.config_store.get_device_id().await {
            Ok(existing_id) if !existing_id.is_empty() => existing_id,
            _ => {
                let new_device_id = ctx.deps.device_id_generator.generate();
                ctx.deps.config_store.save_device_id(new_device_id).await.ok();
                new_device_id
            }
        };
        Event::DeviceEvent(DeviceEvent::DeviceHasBeenSetup(device_id))
    }
}