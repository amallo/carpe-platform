pub trait DeviceIdGenerator<'a> {
    fn generate(&self) -> &'a str;
}