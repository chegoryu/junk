#[macro_use]
extern crate rocket;

#[get("/ping")]
fn ping() -> &'static str {
    "pong"
}

#[launch]
fn rocket() -> _ {
    let rocket = rocket::build();

    rocket.mount("/", routes![ping])
}
