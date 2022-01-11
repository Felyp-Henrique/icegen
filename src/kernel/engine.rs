pub trait Engine <C, R> {
    fn get_context(&self) -> C;
    fn get_document(&self, context: &C) -> R;
}