extern crate tera;

use tera::Tera;
use tera::Context;

pub struct Document {
    context: Context,
}

impl Document {
    pub fn get_context(&self) -> &Context {
        &self.context;
    }
}