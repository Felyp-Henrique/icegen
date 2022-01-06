mod arguments;

use clap::Parser;
use arguments::Icecast;

fn main() {
    let args = Icecast::parse();

    println!("{}", args.port);
}
