pub use crate::kernel::entities::*;
pub use crate::kernel::engine::*;
pub use std::fs::File;
pub use std::io::prelude::*;
pub use std::path::Path;

pub mod new {
    use super::*;
    
    pub fn generate<C>(context: C, engine: &mut dyn Engine<C>) {
        let output = engine.get_document(&context);
        let path = Path::new("icecast.xml");
        let mut file  = match File::create(&path) {
            Ok(file) => file,
            Err(err) => {
                panic!("Coundn't create {}: {}", path.display(), err);
            }
        };
        match file.write_all(output.as_bytes()) {
            Ok(_) => {
                println!("Success!");
            },
            Err(err) => {
                panic!("couldn't write to {}: {}", path.display(), err);
            }
        };
    }
}