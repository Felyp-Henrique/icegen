pub trait Icecast {
    fn get_num_clients(&self) -> u16;
    fn get_num_sources(&self) -> i32;
    fn get_queue(&self) -> u128;
    fn get_cli_timeout(&self) -> u16;
    fn get_hdr_timeout(&self) -> u16;
    fn get_src_timeout(&self) -> u16;
    fn get_burst(&self) -> u128;
    fn get_src_pass(&self) -> String;
    fn get_admin(&self) -> String;
    fn get_admin_pass(&self) -> String;
    fn get_host(&self) -> String;
    fn get_port(&self) -> u32;
    fn get_relay_on(&self) -> bool;
    fn get_relay_host(&self) -> String;
    fn get_relay_port(&self) -> u32;
    fn get_relay_update_interval(&self) -> u32;
    fn get_relay_user(&self) -> String;
    fn get_relay_password(&self) -> String;
    fn get_relay_demand(&self) -> String;
}