use clap::Parser;

use figment::providers::{Format, Json};
use rocket::config::Config;
use rocket::figment::Figment;

use std::path::Path;

#[derive(Parser)]
struct Args {
    #[arg(short, long, default_value = "config.json")]
    config: String,
}

pub fn get_rocket_config() -> Figment {
    let args = Args::parse();
    if !Path::new(&args.config).exists() {
        panic!(
            "Failed to load config from '{}': file does not exist",
            args.config
        );
    }

    Config::figment().merge(Json::file(args.config))
}
