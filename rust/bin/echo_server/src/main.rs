mod config;

mod catchers;
mod handlers;

use crate::config::get_rocket_config;

use crate::catchers::register_catchers;
use crate::handlers::mount_handlers;

use rocket::{Build, Rocket, launch};

#[launch]
fn rocket() -> Rocket<Build> {
    let mut rocket_builder = rocket::custom(get_rocket_config());
    rocket_builder = register_catchers(rocket_builder);
    rocket_builder = mount_handlers(rocket_builder);

    rocket_builder
}
