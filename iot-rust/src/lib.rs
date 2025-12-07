#![no_std]

// Exemple de fonction testable (ne dépend pas d'ESP32)
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

// Exemple d'abstraction pour code testable avec std/no_std
// Le code métier utilise des traits, les implémentations ESP32 sont dans le binaire
pub trait Logger {
    fn log(&self, message: &str);
}

// Code métier testable (ne dépend pas d'ESP32)
pub struct MessageProcessor {
    logger: Option<&'static dyn Logger>,
}

impl MessageProcessor {
    pub fn new(logger: Option<&'static dyn Logger>) -> Self {
        Self { logger }
    }

    pub fn process<'a>(&self, message: &'a str) -> Result<&'a str, &'static str> {
        if message.is_empty() {
            return Err("Message vide");
        }
        if let Some(logger) = self.logger {
            logger.log(message);
        }
        Ok(message)
    }
}

#[cfg(test)]
mod tests {
    // Les tests peuvent utiliser std même si le crate est no_std
    extern crate std;
    
    use super::*;

    // Mock pour les tests
    struct MockLogger {
        logged_messages: std::vec::Vec<std::string::String>,
    }

    impl MockLogger {
        fn new() -> Self {
            Self {
                logged_messages: std::vec::Vec::new(),
            }
        }

        fn get_messages(&self) -> &[std::string::String] {
            &self.logged_messages
        }
    }

    impl Logger for MockLogger {
        fn log(&self, message: &str) {
            // Note: Dans un vrai mock, on utiliserait RefCell ou autre pour mutabilité
            // Ici c'est juste pour l'exemple
            std::println!("Mock log: {}", message);
        }
    }

    #[test]
    fn test_add() {
        assert_eq!(add(2, 2), 4);
        assert_eq!(add(5, 3), 8);
    }

    #[test]
    fn test_message_processor_success() {
        let processor = MessageProcessor::new(None);
        assert_eq!(processor.process("Hello"), Ok("Hello"));
    }

    #[test]
    fn test_message_processor_empty() {
        let processor = MessageProcessor::new(None);
        assert_eq!(processor.process(""), Err("Message vide"));
    }

    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
