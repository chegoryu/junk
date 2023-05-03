mod config;
mod handlers;

use crate::config::get_rocket_config;
use crate::handlers::mount_handlers;

use rocket::{launch, Build, Rocket};

#[launch]
fn rocket() -> Rocket<Build> {
    let mut rocket_builder = rocket::custom(get_rocket_config());
    rocket_builder = mount_handlers(rocket_builder);

    rocket_builder
}
