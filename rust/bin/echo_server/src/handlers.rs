use rocket::{get, routes, Build, Rocket};

#[get("/ping")]
fn ping() -> &'static str {
    "pong\n"
}

pub fn mount_handlers(rocket_builder: Rocket<Build>) -> Rocket<Build> {
    rocket_builder.mount("/", routes![ping])
}

#[cfg(test)]
mod test {
    use super::*;

    use rocket::uri;

    use rocket::http::Status;
    use rocket::local::blocking::Client;

    fn rocket() -> Rocket<Build> {
        mount_handlers(rocket::build())
    }

    #[test]
    fn test_ping() {
        let client = Client::tracked(rocket()).expect("Failed to create rocket client");
        let response = client.get(uri!(ping)).dispatch();

        assert_eq!(response.status(), Status::Ok);
        assert_eq!(response.into_string(), Some("pong\n".to_owned()));
    }
}
