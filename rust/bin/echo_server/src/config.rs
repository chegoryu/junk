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
            "Failed to load config from '{}' file: file does not exist",
            args.config
        );
    }

    Config::figment().merge(Json::file(args.config))
}

#[cfg(test)]
mod tests {
    use super::*;

    use figment::Jail;
    use rocket::Config;

    #[test]
    fn test_get_rocket_config_success() {
        Jail::expect_with(|jail| {
            jail.create_file(
                "config.json",
                r#"
                {
                    "port": 1234,
                    "workers": 22
                }
                "#,
            )
            .unwrap();

            let config = get_rocket_config().extract::<Config>().unwrap();
            assert_eq!(config.port, 1234);
            assert_eq!(config.workers, 22);

            Ok(())
        });
    }

    #[test]
    #[should_panic(expected = "Failed to load config from 'config.json' file: file does not exist")]
    fn test_get_rocket_config_no_config_file() {
        Jail::expect_with(|_jail| {
            get_rocket_config().extract::<Config>().unwrap();

            Ok(())
        });

        get_rocket_config().extract::<Config>().unwrap();
    }

    #[test]
    #[should_panic(expected = "expected value at line 1 column 2")]
    fn test_get_rocket_config_bad_config_file() {
        Jail::expect_with(|jail| {
            jail.create_file("config.json", "{bad json").unwrap();

            get_rocket_config().extract::<Config>().unwrap();

            Ok(())
        });
    }
}
