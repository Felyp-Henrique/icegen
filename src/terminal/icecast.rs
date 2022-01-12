pub use crate::terminal::new::*;
pub use clap::{ Parser, Subcommand };

#[derive(Subcommand)]
pub enum Commands {
    New(NewTerminal)
}

#[derive(Parser)]
#[clap(about, version, author)]
pub struct IcecastTerminal {

    #[clap(subcommand)]
    pub command: Commands,
}