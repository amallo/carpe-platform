use crate::core::runtime::events::events::Event;
use crate::core::runtime::state::State;
use crate::core::runtime::dependencies::Dependencies;
use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;

pub struct CommandContext<'a, C: ConfigStorage, D: DeviceIdGenerator> {
    pub state: &'a mut State,
    pub deps: &'a Dependencies<'a, C, D>,
}

pub trait CommandHandler<'a, C: ConfigStorage, D: DeviceIdGenerator> {
    async fn execute(&self, ctx: &'a CommandContext<'_, C, D>) -> Event;
}