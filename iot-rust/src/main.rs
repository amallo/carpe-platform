#![no_std]
#![no_main]

use esp_idf_sys as _;

#[no_mangle]
pub extern "C" fn app_main() {
    esp_idf_sys::esp_println::esp_println!("Hello, world!");
}
