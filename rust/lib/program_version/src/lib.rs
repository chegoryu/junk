pub fn get_program_version(git_describe: &str) -> String {
    let mut program_version: String = git_describe.to_owned();

    for prefix in ["heads/", "remotes/", "tags/"] {
        if let Some(stripped_program_version) = program_version.strip_prefix(prefix) {
            program_version = stripped_program_version.to_owned();
            break;
        }
    }

    if let Some(stripped_program_version) = program_version.strip_suffix("-dirty") {
        program_version = stripped_program_version.to_owned() + "-modified";
    }

    program_version
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_program_version() {
        assert_eq!(get_program_version("unknown"), "unknown");
        assert_eq!(get_program_version("random"), "random");

        assert_eq!(
            get_program_version("heads/main/something"),
            "main/something"
        );
        assert_eq!(
            get_program_version("remotes/main/something"),
            "main/something"
        );
        assert_eq!(
            get_program_version("tags/v1.0.0/something"),
            "v1.0.0/something"
        );

        assert_eq!(
            get_program_version("heads/tags/remotes/tags/heads/v1.0.0/something"),
            "tags/remotes/tags/heads/v1.0.0/something"
        );
        assert_eq!(
            get_program_version("tags/heads/remotes/heads/tags/v1.0.0/something"),
            "heads/remotes/heads/tags/v1.0.0/something"
        );
        assert_eq!(
            get_program_version("remotes/heads/tags/heads/remotes/v1.0.0/something"),
            "heads/tags/heads/remotes/v1.0.0/something"
        );

        assert_eq!(get_program_version("main-dirty"), "main-modified");
        assert_eq!(get_program_version("main-dirty-x"), "main-dirty-x");
        assert_eq!(get_program_version("heads/main-dirty"), "main-modified");
    }
}
