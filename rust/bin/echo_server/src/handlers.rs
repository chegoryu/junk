use rocket::get;

#[get("/ping")]
pub fn ping() -> &'static str {
    "pong"
}
