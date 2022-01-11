pub use crate::kernel::entities::*;
pub use clap::Parser;

#[derive(Parser)]
#[clap(about, version, author)]
pub struct IcecastTerminal {

    #[clap(long)]
    path: String,

    #[clap(long, default_value = "1000")]
    num_clients: u16,

    #[clap(long, default_value = "1")]
    num_sources: i32,

    #[clap(long, default_value = "524288")]
    queue: u128,

    #[clap(long, default_value = "30")]
    cli_timeout: u16,

    #[clap(long, default_value = "15")]
    hdr_timeout: u16,

    #[clap(long, default_value = "10")]
    src_timeout: u16,

    #[clap(long, default_value = "65535")]
    burst: u128,

    #[clap(long, default_value = "hackme")]
    src_pass: String,

    #[clap(long, default_value = "admin")]
    admin: String,

    #[clap(long, default_value = "hackme")]
    admin_pass: String,

    #[clap(long, default_value = "localhost")]
    host: String,

    #[clap(long, default_value = "8000")]
    port: u32,

    #[clap(long)]
    relay_on: bool,

    #[clap(long, default_value = "127.0.0.1")]
    relay_host: String,

    #[clap(long, default_value = "8001")]
    relay_port: u32,

    #[clap(long, default_value = "120")]
    relay_update_interval: u32,

    #[clap(long, default_value = "relay")]
    relay_user: String,

    #[clap(long, default_value = "hackme")]
    relay_password: String,

    #[clap(long, default_value = "1")]
    relay_demand: String,
}

impl Icecast for IcecastTerminal {

    fn get_path(&self) -> String {
        self.path.clone()
    }

    fn get_num_clients(&self) -> u16 {
        self.num_clients
    }

    fn get_num_sources(&self) -> i32 {
        self.num_sources
    }

    fn get_queue(&self) -> u128 {
        self.queue
    }

    fn get_cli_timeout(&self) -> u16 {
        self.cli_timeout
    }

    fn get_hdr_timeout(&self) -> u16 {
        self.hdr_timeout
    }

    fn get_src_timeout(&self) -> u16 {
        self.src_timeout
    }

    fn get_burst(&self) -> u128 {
        self.burst
    }

    fn get_src_pass(&self) -> String {
        self.src_pass.clone()
    }

    fn get_admin(&self) -> String {
        self.admin.clone()
    }

    fn get_admin_pass(&self) -> String {
        self.admin_pass.clone()
    }

    fn get_host(&self) -> String {
        self.host.clone()
    }

    fn get_port(&self) -> u32 {
        self.port
    }

    fn get_relay_on(&self) -> bool {
        self.relay_on
    }

    fn get_relay_host(&self) -> String {
        self.relay_host.clone()
    }

    fn get_relay_port(&self) -> u32 {
        self.relay_port
    }

    fn get_relay_update_interval(&self) -> u32 {
        self.relay_update_interval
    }

    fn get_relay_user(&self) -> String {
        self.relay_user.clone()
    }

    fn get_relay_password(&self) -> String {
        self.relay_password.clone()
    }

    fn get_relay_demand(&self) -> String {
        self.relay_demand.clone()
    }
}