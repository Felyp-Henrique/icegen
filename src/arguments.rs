use clap::Parser;

#[derive(Parser)]
#[clap(about, version, author)]
pub struct Icecast {

    #[clap(long, default_value = "1000")]
    pub num_clients: u16,

    #[clap(long, default_value = "1")]
    pub num_sources: i32,

    #[clap(long, default_value = "524288")]
    pub queue: u128,

    #[clap(long, default_value = "30")]
    pub cli_timeout: u16,

    #[clap(long, default_value = "15")]
    pub hdr_timeout: u16,

    #[clap(long, default_value = "10")]
    pub src_timeout: u16,

    #[clap(long, default_value = "65535")]
    pub burst: u128,

    #[clap(long, default_value = "")]
    pub src_pass: String,

    #[clap(long, default_value = "")]
    pub admin: String,

    #[clap(long, default_value = "")]
    pub admin_pass: String,

    #[clap(long, default_value = "localhost")]
    pub host: String,

    #[clap(long, default_value = "8000")]
    pub port: u32,

    #[clap(long)]
    pub relay_on: bool,

    #[clap(long, default_value = "127.0.0.1")]
    pub relay_host: String,

    #[clap(long, default_value = "8001")]
    pub relay_port: u32,

    #[clap(long, default_value = "120")]
    pub relay_update_interval: u32,

    #[clap(long, default_value = "relay")]
    pub relay_user: String,

    #[clap(long, default_value = "hackme")]
    pub relay_password: String,

    #[clap(long, default_value = "1")]
    pub relay_demand: String,
}
