use crate::core::runtime::events::event_queue::EventQueue;
use crate::core::runtime::events::events::Event;
use crate::core::runtime::command_handler::{CommandContext, CommandHandler};
use crate::core::device::commands::setup_device_command::SetupDeviceCommandHandler;
use crate::core::runtime::events::events::Command;
use crate::core::device::commands::events::DeviceCommand;
use crate::core::device::commands::events::DeviceEvent;
use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::device_id_generator::DeviceIdGenerator;
use heapless::Vec;
pub struct Runtime<'a, Q: EventQueue<Event>, C: ConfigStorage, D: DeviceIdGenerator> {
    queue: &'a mut Q,
    context: &'a mut CommandContext<'a, C, D>,
}   

impl<'a, Q: EventQueue<Event>, C: ConfigStorage, D: DeviceIdGenerator> Runtime<'a, Q, C, D> {
    pub fn new(context: &'a mut CommandContext<'a, C, D>, queue: &'a mut Q) -> Self {
        Self { context, queue }
    }
    pub async fn send(&mut self, event: Event) {
        self.queue.push(event).await;
    }
    pub fn decide(&mut self, event: Event) -> Vec<Command, 1> {
        let mut commands = Vec::new();
        match event {
            Event::PowerOn => {
                commands.push(Command::DeviceCommand(DeviceCommand::SetupDevice)).ok();
            }
            Event::DeviceEvent(DeviceEvent::DeviceHasBeenSetup(device_id)) => {
                self.context.state.set_ready(device_id);
            }
        }
        commands
    }
    pub async fn run_until_idle(&mut self) {
        loop {
            let event = match self.queue.pop().await {
                Some(event) => event,
                None => break,
            };
            let commands = self.decide(event);
            for cmd in commands {
                match cmd {
                    Command::DeviceCommand(DeviceCommand::SetupDevice) => {
                        let handler = SetupDeviceCommandHandler::new();
                        let event = handler.execute(&self.context).await;
                        self.queue.push(event).await;
                    }
                }
            }
        }
    }
}
