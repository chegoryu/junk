pub const GIT_DESCRIBE: &str = "unknown";
pub const VERSION: &str = "unknown";

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_git_describe() {
        assert!(!GIT_DESCRIBE.is_empty())
    }

    #[test]
    fn test_version() {
        assert!(!VERSION.is_empty())
    }
}
