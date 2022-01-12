pub trait Engine <C> {
    fn get_context(&self) -> C;
    fn get_document(&self, context: &C) -> String;
}