pub const GIT_DESCRIBE: &str = env!("GIT_DESCRIBE");
pub const PROGRAM_VERSION: &str = env!("PROGRAM_VERSION");

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[allow(clippy::const_is_empty)]
    fn test_git_describe() {
        assert!(!GIT_DESCRIBE.is_empty());
    }

    #[test]
    #[allow(clippy::const_is_empty)]
    fn test_program_version() {
        assert!(!PROGRAM_VERSION.is_empty());
    }
}
