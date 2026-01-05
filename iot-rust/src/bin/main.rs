#![no_std]
#![no_main]
#![deny(
    clippy::mem_forget,
    reason = "mem::forget is generally not safe to do with esp_hal types, especially those \
    holding buffers for the duration of a data transfer."
)]

#[cfg(feature = "esp32")]
extern crate alloc;

use esp_hal::clock::CpuClock;
use esp_hal::main;
use esp_hal::time::{Duration, Instant};
use log::info;

#[cfg(feature = "esp32")]
use linked_list_allocator::LockedHeap;

#[cfg(feature = "esp32")]
#[global_allocator]
static HEAP: LockedHeap = LockedHeap::empty();

#[cfg(feature = "esp32")]
fn init_heap() {
    const HEAP_SIZE: usize = 32 * 1024; // 32KB heap
    static mut HEAP_MEM: [u8; HEAP_SIZE] = [0; HEAP_SIZE];
    unsafe {
        // Utiliser addr_of_mut! pour obtenir un pointeur sans créer de référence
        let heap_ptr = core::ptr::addr_of_mut!(HEAP_MEM) as *mut u8;
        HEAP.lock().init(heap_ptr, HEAP_SIZE);
    }
}

#[panic_handler]
fn panic(_: &core::panic::PanicInfo) -> ! {
    loop {}
}

// This creates a default app-descriptor required by the esp-idf bootloader.
// For more information see: <https://docs.espressif.com/projects/esp-idf/en/stable/esp32/api-reference/system/app_image_format.html#application-description>
esp_bootloader_esp_idf::esp_app_desc!();

#[main]
fn main() -> ! {
    // generator version: 1.0.1

    #[cfg(feature = "esp32")]
    init_heap();

    esp_println::logger::init_logger_from_env();

    let config = esp_hal::Config::default().with_cpu_clock(CpuClock::max());
    let _peripherals = esp_hal::init(config);

    loop {
        info!("Hello world!");
        let delay_start = Instant::now();
        while delay_start.elapsed() < Duration::from_millis(500) {}
    }

    // for inspiration have a look at the examples at https://github.com/esp-rs/esp-hal/tree/esp-hal-v1.0.0/examples/src/bin
}
