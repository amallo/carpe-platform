use crate::core::runtime::events::events::Event;
use crate::core::runtime::state::State;
use crate::core::runtime::dependencies::Dependencies;

pub struct CommandContext<'a> {
    pub state: &'a mut State,
    pub deps: &'a Dependencies<'a>,
}

pub trait CommandHandler {
    async fn execute(&self, ctx: &mut CommandContext<'_>) -> Event;
}