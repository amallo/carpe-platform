
/**
 * @brief TestEventQueue is a queue of events for testing purposes
 */

use crate::core::runtime::events::event_queue::EventQueue;
use crate::core::runtime::events::events::Event;

extern crate std;
use std::collections::VecDeque;

#[cfg(test)]

pub struct TestEventQueue {
    events: VecDeque<Event>,
}

impl TestEventQueue {
    pub fn new() -> Self {
        Self { events: VecDeque::new() }
    }
}

impl EventQueue<Event> for TestEventQueue {
    async fn push(&mut self, event: Event) {
        self.events.push_back(event);
    }
    async fn pop(&mut self) -> Option<Event> {
        self.events.pop_front()
    }
}