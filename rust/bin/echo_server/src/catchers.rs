use rocket::http::Status;
use rocket::{Build, Request, Rocket, catch, catchers};

#[catch(default)]
fn default(status: Status, req: &Request) -> String {
    format!("{} ({} {})\n", status, req.method(), req.uri())
}

pub fn register_catchers(rocket_builder: Rocket<Build>) -> Rocket<Build> {
    rocket_builder.register("/", catchers![default])
}

#[cfg(test)]
mod tests {
    use super::*;

    use rocket::http::Status;
    use rocket::local::blocking::Client;

    fn get_client() -> Client {
        Client::tracked(register_catchers(rocket::build())).expect("Failed to create rocket client")
    }

    #[test]
    fn test_default() {
        let client = get_client();

        {
            let response = client.get("/this_route_does_not_exist").dispatch();
            assert_eq!(response.status(), Status::NotFound);
            assert_eq!(
                response.into_string(),
                Some("404 Not Found (GET /this_route_does_not_exist)\n".to_owned())
            );
        }

        {
            let response = client.post("/this_route_does_not_exist").dispatch();
            assert_eq!(response.status(), Status::NotFound);
            assert_eq!(
                response.into_string(),
                Some("404 Not Found (POST /this_route_does_not_exist)\n".to_owned())
            );
        }
    }
}
