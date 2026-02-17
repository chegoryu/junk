use build_info::PROGRAM_VERSION;
use const_format::formatcp;
use itertools::join;
use rocket::async_trait;
use rocket::http::Header;
use rocket::request::{FromRequest, Outcome};
use rocket::{Build, Request, Rocket, get, routes};

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
        join(
            request_headers
                .headers
                .iter()
                .map(|header| header.name().to_string().to_lowercase() + ": " + header.value()),
            "\n",
        ) + "\n"
    }
}

pub fn mount_handlers(rocket_builder: Rocket<Build>) -> Rocket<Build> {
    rocket_builder.mount("/", routes![ping, version, headers])
}

#[cfg(test)]
mod tests {
    use super::*;

    use const_format::concatcp;
    use rocket::http::Status;
    use rocket::local::blocking::Client;
    use rocket::uri;

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
        struct Test<'r> {
            headers: Vec<Header<'r>>,
            expected_body: &'r str,
        }

        let tests = [
            Test {
                headers: [].to_vec(),
                expected_body: "",
            },
            Test {
                headers: [Header::new("Header", "Value")].to_vec(),
                expected_body: "header: Value\n",
            },
            Test {
                headers: [
                    Header::new("Header", "Value2"),
                    Header::new("Header", "Value1"),
                    Header::new("Header", "Value3"),
                ]
                .to_vec(),
                expected_body: concatcp!(
                    "header: Value1\n",
                    "header: Value2\n",
                    "header: Value3\n",
                ),
            },
            Test {
                headers: [
                    Header::new("Header2", "Header2Value2"),
                    Header::new("Header2", "Header2Value1"),
                    Header::new("Header2", "Header2Value3"),
                    Header::new("Header1", "Header1Value2"),
                    Header::new("Header1", "Header1Value1"),
                    Header::new("ZYXLastHeader", "AAFirstValue"),
                    Header::new("OtherHeader", "SomeValue"),
                ]
                .to_vec(),
                expected_body: concatcp!(
                    "header1: Header1Value1\n",
                    "header1: Header1Value2\n",
                    "header2: Header2Value1\n",
                    "header2: Header2Value2\n",
                    "header2: Header2Value3\n",
                    "otherheader: SomeValue\n",
                    "zyxlastheader: AAFirstValue\n",
                ),
            },
            Test {
                headers: [Header::new("host", "somehost.com:1234")].to_vec(),
                expected_body: "host: somehost.com:1234\n",
            },
            Test {
                headers: [
                    Header::new("aHeader", "aHeaderValue"),
                    Header::new("BHeader", "BHeaderValue"),
                    Header::new("cHeader", "cHeaderValue"),
                    Header::new("DHeader", "DHeaderValue"),
                ]
                .to_vec(),
                expected_body: concatcp!(
                    "aheader: aHeaderValue\n",
                    "bheader: BHeaderValue\n",
                    "cheader: cHeaderValue\n",
                    "dheader: DHeaderValue\n",
                ),
            },
            Test {
                headers: [
                    Header::new("Header", "aHeaderValue"),
                    Header::new("Header", "BHeaderValue"),
                    Header::new("Header", "cHeaderValue"),
                    Header::new("Header", "DHeaderValue"),
                ]
                .to_vec(),
                expected_body: concatcp!(
                    "header: BHeaderValue\n",
                    "header: DHeaderValue\n",
                    "header: aHeaderValue\n",
                    "header: cHeaderValue\n",
                ),
            },
        ];

        let client = get_client();

        for test in tests {
            let mut local_request = client.get(uri!(headers));
            for header in test.headers {
                local_request = local_request.header(header)
            }
            let response = local_request.dispatch();
            assert_eq!(response.status(), Status::Ok);
            assert_eq!(response.into_string(), Some(test.expected_body.to_owned()));
        }
    }
}
