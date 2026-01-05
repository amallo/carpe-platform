pub trait DeviceIdGenerator {
    fn generate(&self) -> &'static str;
}