#[macro_use]
extern crate rocket;

use clap::Parser;
use figment::providers::{Format, Json};
use rocket::config::Config;

#[derive(Parser)]
struct Args {
    #[arg(short, long, default_value = "config.json")]
    config: String,
}

#[get("/ping")]
fn ping() -> &'static str {
    "pong"
}

#[launch]
fn rocket() -> _ {
    let args = Args::parse();
    let config = Config::figment()
        .merge(Json::file(args.config));

    rocket::custom(config).mount("/", routes![ping])
}
