// std

// crates
use clap::{App, Arg, SubCommand};
use lanner::evaluate;

// local

fn main() {
    let app = App::new("lanner")
        .version(env!("CARGO_PKG_VERSION"))
        .author(env!("CARGO_PKG_AUTHORS"))
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .subcommand(
            SubCommand::with_name("evaluate")
                .about("Evaluate an expression")
                .arg(Arg::with_name("INPUT").takes_value(true).required(true)),
        )
        .get_matches();

    if let Some(v) = app.subcommand_matches("evaluate") {
        print!("{}", evaluate(v.value_of("INPUT").unwrap()).unwrap());
    }
}
