mod terminal;
mod kernel;
mod documents;

use terminal::icecast::*;

fn main() {
    let args: IcecastTerminal = IcecastTerminal::parse();

    println!("{}", args.get_port());
}
