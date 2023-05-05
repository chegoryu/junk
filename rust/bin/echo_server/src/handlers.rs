use build_info::PROGRAM_VERSION;
use const_format::formatcp;
use rocket::async_trait;
use rocket::http::Header;
use rocket::request::{FromRequest, Outcome};
use rocket::{get, routes, Build, Request, Rocket};

struct RequestHeaders<'r> {
    headers: Vec<Header<'r>>,
}

#[async_trait]
impl<'r> FromRequest<'r> for RequestHeaders<'r> {
    type Error = ();

    async fn from_request(request: &'r Request<'_>) -> Outcome<RequestHeaders<'r>, ()> {
        let mut request_headers = RequestHeaders {
            headers: request.headers().iter().collect(),
        };
        request_headers
            .headers
            .sort_by(|a, b| a.name().cmp(b.name()).then(a.value().cmp(b.value())));

        Outcome::Success(request_headers)
    }
}

#[get("/ping")]
fn ping() -> &'static str {
    "pong\n"
}

#[get("/version")]
fn version() -> &'static str {
    formatcp!("{}\n", PROGRAM_VERSION)
}

#[get("/headers")]
fn headers(request_headers: RequestHeaders) -> String {
    if request_headers.headers.is_empty() {
        String::new()
    } else {
        request_headers
            .headers
            .iter()
            .map(|header| header.to_string())
            .collect::<Vec<String>>()
            .join("\n")
            + "\n"
    }
}

pub fn mount_handlers(rocket_builder: Rocket<Build>) -> Rocket<Build> {
    rocket_builder.mount("/", routes![ping, version, headers])
}

#[cfg(test)]
mod tests {
    use super::*;

    use rocket::uri;

    use rocket::http::Status;
    use rocket::local::blocking::Client;

    fn get_client() -> Client {
        Client::tracked(mount_handlers(rocket::build())).expect("Failed to create rocket client")
    }

    #[test]
    fn test_ping() {
        let client = get_client();

        let response = client.get(uri!(ping)).dispatch();
        assert_eq!(response.status(), Status::Ok);
        assert_eq!(response.into_string(), Some("pong\n".to_owned()));
    }

    #[test]
    fn test_version() {
        let client = get_client();

        let response = client.get(uri!(version)).dispatch();
        assert_eq!(response.status(), Status::Ok);
        assert_eq!(
            response.into_string(),
            Some(formatcp!("{}\n", PROGRAM_VERSION).to_owned())
        );
    }

    #[test]
    fn test_headers() {
        let client = get_client();

        let response = client.get(uri!(headers)).dispatch();
        assert_eq!(response.status(), Status::Ok);
        assert_eq!(response.into_string(), Some("".to_owned()));
        // TODO: More tests.
    }
}
