pub fn add(nums: Vec<f64>) -> f64 {
    let mut sum: f64 = 0.0;
    for number in nums.iter() {
        sum += number;
    }
    sum
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn check_add() {
        assert_eq!(add(vec![2.2, 2.2]), 4.4);
    }
}
