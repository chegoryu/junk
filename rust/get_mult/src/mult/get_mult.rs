use std::ops::Mul;

pub fn get_mult<T: Mul<Output = T>>(a: T, b: T) -> T {
    a * b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_mult() {
        assert_eq!(get_mult(0, 0), 0);
        assert_eq!(get_mult(1, 2), 2);
        assert_eq!(get_mult(-7, 11), -77);
        assert_eq!(get_mult(-5, -1), 5);
        assert_eq!(get_mult(234, 434), 101556);
    }
}
