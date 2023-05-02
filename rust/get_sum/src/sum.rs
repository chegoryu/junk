use std::ops::Add;

pub fn get_sum<T: Add<Output = T>>(a: T, b: T) -> T {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_sum() {
        assert_eq!(get_sum(0, 0), 0);
        assert_eq!(get_sum(1, 2), 3);
        assert_eq!(get_sum(-7, 11), 4);
        assert_eq!(get_sum(-5, -1), -6);
            assert_eq!(get_sum(234, 434), 668);
    }
}
