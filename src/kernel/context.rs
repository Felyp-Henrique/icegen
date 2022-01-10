use tera;
use serde::ser::Serialize;
pub trait Context <T> {
    fn get_data(&self) -> &T;
    fn insert<K, V>(&mut self, k: K, v: V) where V: Serialize;
}

pub struct TeraContenxt {
    data: tera::Context,
}

impl TeraContenxt {
    pub fn new() -> Self {
        TeraContenxt {
            data: tera::Context::new(),
        }
    }
}

impl Context<tera::Context> for TeraContenxt {
    fn get_data(&self) -> &tera::Context {
        &self.data
    }

    fn insert<K, V>(&mut self, k: K, v: V) where V: Serialize {
    }
}