pub use crate::kernel::entities::*;
pub use crate::kernel::engine::*;
pub use tera;

use std::process::exit;

pub struct EngineTera {
    icecast: dyn Icecast,
}

impl Engine<tera::Context, String> for EngineTera {
    fn get_context(&self) -> tera::Context {
        let mut context = tera::Context::new();
        context.insert("num_clients", &self.icecast.get_num_clients());
        context.insert("num_sources", &self.icecast.get_num_sources());
        context.insert("queue", &self.icecast.get_queue());
        context.insert("cli_timeout", &self.icecast.get_cli_timeout());
        context.insert("hdr_timeout", &self.icecast.get_hdr_timeout());
        context.insert("src_timeout", &self.icecast.get_src_timeout());
        context.insert("burst", &self.icecast.get_burst());
        context.insert("src_pass", &self.icecast.get_src_pass());
        context.insert("admin", &self.icecast.get_admin());
        context.insert("admin_pass", &self.icecast.get_admin_pass());
        context.insert("host", &self.icecast.get_host());
        context.insert("port", &self.icecast.get_port());
        context.insert("relay_on", &self.icecast.get_relay_on());
        context.insert("relay_host", &self.icecast.get_relay_host());
        context.insert("relay_port", &self.icecast.get_relay_port());
        context.insert("relay_update_interval", &self.icecast.get_relay_update_interval());
        context.insert("relay_user", &self.icecast.get_relay_user());
        context.insert("relay_password", &self.icecast.get_relay_password());
        context.insert("relay_demand", &self.icecast.get_relay_demand());
        context
    }

    fn get_document(&self, context: &tera::Context) -> String {
        let tera = match tera::Tera::new(&self.icecast.get_path()) {
            Ok(engine) => engine,
            Err(err) => {
                println!("Parsing error(s): {}", err);
                exit(1);
            }
        };
        let result = match tera.render("icecast.xml", context) {
            Ok(document) => document,
            Err(err) => {
                println!("Render error(s): {}", err);
                exit(1);
            }
        };
        result
    }
}
