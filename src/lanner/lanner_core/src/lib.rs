/// Returns an `f64` that represents the sum of all numbers in the vector.
/// ```rust
/// add(vec![10.0, 10.1]) // returns 20.1
/// ```
pub fn add(nums: Vec<f64>) -> f64 {
    let mut sum: f64 = 0.0;
    for number in nums.iter() {
        sum += number;
    }

    sum
}

/// Returns an `f64` that represents the difference between the first `f64` in the `vector` and the following vector items.
/// ```rust
/// subtract(vec![2.0, 1.0, 1.0]) // returns 0, because 2 - 1 - 1 is 0
/// ```
/// Since `2.0` is the first item, it is the number that will be the starting value; that is, each consecutive number will subtract from it after the previous number has, thus creating the difference, `0.0`.
pub fn subtract(nums: Vec<f64>) -> f64 {
    let mut diff = nums[0];
    let mut sub_iter = nums.iter();
    
    sub_iter.next();
    for number in sub_iter {
        diff -= number;
    }

    diff
}

pub fn multiply(nums: Vec<f64>) -> f64 {
    let mut product = nums[0];
    let mut mult_iter = nums.iter();
    
    mult_iter.next();
    for factor in mult_iter {
        product *= factor;
    }

    product
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
    #[test]
    fn check_multiply() {
        assert_eq!(multiply(vec![2.0, 2.0]), 4.0);
    }
}
