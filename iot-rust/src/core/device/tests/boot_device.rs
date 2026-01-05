

#[cfg(test)]

extern crate std;
    

use crate::core::device::gateways::mock_config_storage::MockConfigStorage;
use crate::core::device::gateways::config_storage::ConfigStorage;
use crate::core::device::gateways::mock_device_id_generator::MockDeviceIdGenerator;
use crate::core::runtime::runtime::Runtime;
use crate::core::runtime::dependencies::Dependencies;

use crate::core::runtime::events::test_event_queue::TestEventQueue;
use crate::core::runtime::events::events::Event;
use crate::core::runtime::state::State;
use crate::core::runtime::command_handler::CommandContext;

// TestContext encapsule tous les composants nécessaires pour les tests
struct TestContext {
    config_storage: MockConfigStorage,
    device_id_generator: MockDeviceIdGenerator,
    state: State,
    event_queue: TestEventQueue,
}

impl TestContext {
    fn new() -> Self {
        Self {
            config_storage: MockConfigStorage::new(),
            device_id_generator: MockDeviceIdGenerator::new(),
            state: State::new(),
            event_queue: TestEventQueue::new(),
        }
    }
    
    fn with_device_id_generator_will_return(mut self, device_id: &'static str) -> Self {
        self.device_id_generator.will_generate_device_id(device_id);
        self
    }
    
    async fn with_config_storage_has_device_id(mut self, device_id: &'static str) -> Self {
        self.config_storage.save_device_id(device_id).await.ok();
        self
    }
    
    async fn when_power_on(&mut self) {
        let deps = Dependencies::new(&self.config_storage, &self.device_id_generator);
        let mut command_context = CommandContext {
            state: &mut self.state,
            deps: &deps,
        };
        let mut runtime = Runtime::new(&mut command_context, &mut self.event_queue);
        
        runtime.send(Event::PowerOn).await;
        runtime.run_until_idle().await;
    }
    
    fn then_device_id_in_state_is(&self, expected_id: &'static str) {
        assert_eq!(self.state.device_id(), Some(expected_id));
    }
    
    fn then_config_storage_has_device_id(&self, expected_id: &'static str) {
        assert_eq!(*self.config_storage.device_id.borrow(), expected_id);
    }
}

// Helpers Given
fn given_empty_device() -> TestContext {
    TestContext::new()
}

async fn given_stored_device_with_id(device_id: &'static str) -> TestContext {
    TestContext::new().with_config_storage_has_device_id(device_id).await
}
#[cfg(test)]
#[test]
fn setup_device_first_time() {
    use futures::executor::block_on;
    
    block_on(async {
        let mut ctx = given_empty_device()
            .with_device_id_generator_will_return("12324");
        
        ctx.when_power_on().await;
        
        ctx.then_device_id_in_state_is("12324");
        ctx.then_config_storage_has_device_id("12324");
    });
}

#[cfg(test)]
#[test]
fn setup_already_setup_device() {
    use futures::executor::block_on;
    
    block_on(async {
        let mut ctx = given_stored_device_with_id("1234").await;
        
        ctx.when_power_on().await;
        
        ctx.then_device_id_in_state_is("1234");
    });
}
