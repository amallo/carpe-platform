

#[cfg(test)]

extern crate std;
    

use crate::core::device::gateways::mock_config_storage::MockConfigStorage;
use crate::core::device::gateways::mock_device_id_generator::MockDeviceIdGenerator;
use crate::core::runtime::runtime::Runtime;
use crate::core::runtime::dependencies::Dependencies;

use crate::core::runtime::events::test_event_queue::TestEventQueue;
use crate::core::runtime::events::events::Event;
use crate::core::runtime::state::State;
use crate::core::runtime::command_handler::CommandContext;
#[cfg(test)] 
async fn device_boots_and_becomes_ready() {


   let config_storage = MockConfigStorage::new();
   let device_id_generator = MockDeviceIdGenerator::new();
   device_id_generator.will_generate_device_id("12324");
   let deps = Dependencies::new(&config_storage, &device_id_generator);
   let mut event_queue = TestEventQueue::new();
   let mut state = State::new();
   let mut ctx = CommandContext {
    state: &mut state,
    deps: &deps,
   };
   let mut runtime = Runtime::new(&mut ctx, &mut event_queue);

   runtime.send(Event::PowerOn).await;
   runtime.run_until_idle().await;

  assert!(state.is_ready());
   
}
