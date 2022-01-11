use crate::kernel::entities::*;

pub struct IcecastTera {
    path: String,
    num_clients: u16,
    num_sources: i32,
    queue: u128,
    cli_timeout: u16,
    hdr_timeout: u16,
    src_timeout: u16,
    burst: u128,
    src_pass: String,
    admin: String,
    admin_pass: String,
    host: String,
    port: u32,
    relay_on: bool,
    relay_host: String,
    relay_port: u32,
    relay_update_interval: u32,
    relay_user: String,
    relay_password: String,
    relay_demand: String,
}

impl IcecastTera {
    fn new(path: String, icecast: &dyn Icecast) -> Self {
        IcecastTera{
            path: path.clone(),
            num_clients: icecast.get_num_clients(),
            num_sources: icecast.get_num_sources(),
            queue: icecast.get_queue(),
            cli_timeout: icecast.get_cli_timeout(),
            hdr_timeout: icecast.get_hdr_timeout(),
            src_timeout: icecast.get_src_timeout(),
            burst: icecast.get_burst(),
            src_pass: icecast.get_src_pass(),
            admin: icecast.get_admin(),
            admin_pass: icecast.get_admin_pass(),
            host: icecast.get_host(),
            port: icecast.get_port(),
            relay_on: icecast.get_relay_on(),
            relay_host: icecast.get_relay_host(),
            relay_port: icecast.get_relay_port(),
            relay_update_interval: icecast.get_relay_update_interval(),
            relay_user: icecast.get_relay_user(),
            relay_password: icecast.get_relay_password(),
            relay_demand: icecast.get_relay_demand(),
        }
    }
}

impl Icecast for IcecastTera {

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