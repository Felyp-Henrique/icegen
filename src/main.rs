mod terminal;
mod kernel;
mod models;
mod usecases;

use terminal::icecast::*;
use kernel::tera::*;
use models::icecast::*;
use usecases::new::new;

fn main() {
    let args: IcecastTerminal = IcecastTerminal::parse();

    match args.command {
        Commands::New(data) => {
            let icecast = IcecastTera::new(data.get_path(), &data);
        }
    };
}
