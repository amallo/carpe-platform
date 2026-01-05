

pub trait EventQueue<E> {
    async fn push(&mut self, event: E);
    async fn pop(&mut self) -> Option<E>;
}