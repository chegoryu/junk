mod config;
mod handlers;

use crate::config::get_rocket_config;

use rocket::{launch, routes};

#[launch]
fn rocket() -> _ {
    rocket::custom(get_rocket_config()).mount("/", routes![handlers::ping])
}
