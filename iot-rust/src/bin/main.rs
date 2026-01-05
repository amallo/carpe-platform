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
use esp_hal::i2c::master::{Config as I2cConfig, I2c};
#[cfg(feature = "esp32")]
use esp_hal::time::Rate;
#[cfg(feature = "esp32")]
use ssd1306::{I2CDisplayInterface, Ssd1306, prelude::*};
#[cfg(feature = "esp32")]
use embedded_graphics::{
    mono_font::{ascii::FONT_6X10, MonoTextStyleBuilder},
    pixelcolor::BinaryColor,
    prelude::*,
    text::{Baseline, Text},
};

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
    let peripherals = esp_hal::init(config);

    #[cfg(feature = "esp32")]
    {
        // Initialiser I2C : SDA=21, SCL=22
        let i2c_bus = I2c::new(
            peripherals.I2C0,
            I2cConfig::default().with_frequency(Rate::from_khz(400)),
        )
        .unwrap()
        .with_scl(peripherals.GPIO22)
        .with_sda(peripherals.GPIO21);

        // Initialiser l'écran SSD1306 (128x64, adresse 0x3C)
        let interface = I2CDisplayInterface::new(i2c_bus);
        let mut display = Ssd1306::new(interface, DisplaySize128x64, DisplayRotation::Rotate0)
            .into_buffered_graphics_mode();
        display.init().unwrap();

        // Afficher "hello"
        let text_style = MonoTextStyleBuilder::new()
            .font(&FONT_6X10)
            .text_color(BinaryColor::On)
            .build();

        Text::with_baseline("hello", Point::new(0, 20), text_style, Baseline::Top)
            .draw(&mut display)
            .unwrap();

        display.flush().unwrap();

        info!("Message 'hello' affiché sur l'écran OLED");
    }

    loop {
        info!("Hello world!");
        let delay_start = Instant::now();
        while delay_start.elapsed() < Duration::from_millis(500) {}
    }

    // for inspiration have a look at the examples at https://github.com/esp-rs/esp-hal/tree/esp-hal-v1.0.0/examples/src/bin
}
