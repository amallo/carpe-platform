use crate::core::runtime::events::event_queue::EventQueue;
use crate::core::runtime::events::events::Event;
use crate::core::runtime::dependencies::Dependencies;
use crate::core::runtime::state::State;
use crate::core::runtime::command_handler::{CommandContext, CommandHandler};
use crate::core::device::commands::setup_device_command::SetupDeviceCommandHandler;
use crate::core::runtime::events::events::Command;
use crate::core::device::commands::events::DeviceCommand;
use heapless::Vec;
pub struct Runtime<'a, Q: EventQueue<Event>> {
    queue: &'a mut Q,
    context: &'a mut CommandContext<'a>,
}   

impl<'a, Q: EventQueue<Event>> Runtime<'a, Q> {
    pub fn new(context: &'a mut CommandContext<'a>, queue: &'a mut Q) -> Self {
        Self { context, queue }
    }
    pub async fn send(&mut self, event: Event) {
        self.queue.push(event).await;
    }
    pub fn decide(&mut self, event: Event) -> Vec<Command, 1> {
        let mut commands = Vec::new();
        match event {
            Event::PowerOn => {
                // Seulement générer SetupDevice pour PowerOn
                commands.push(Command::DeviceCommand(DeviceCommand::SetupDevice)).ok();
            }
            _ => {
                // Pour les autres événements, ne rien faire
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
                        let event = handler.execute(&mut self.context).await;
                        self.queue.push(event).await;
                    }
                }
            }
        }
    }
}
