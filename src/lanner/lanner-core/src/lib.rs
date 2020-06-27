pub fn add(nums: Vec<f64>) -> f64 {
    let mut sum: f64 = 0.0;
    for number in nums.iter() {
        sum += number;
    }
    sum
}

pub fn subtract(nums: Vec<f64>) -> f64 {
    let mut diff = nums[0];
    let mut subiter = nums.iter();
    
    subiter.next();
    for number in subiter {
        diff -= number;
    }
    diff
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn check_add() {
        assert_eq!(add(vec![10.0000025, 10.00001]), 20.0000125);
        assert_eq!(add(vec![2.2, 2.2]), 4.4);
    }
    #[test]
    fn check_subtract() {
        assert_eq!(subtract(vec![2.2, 2.2]), 0.0);
    }
}
