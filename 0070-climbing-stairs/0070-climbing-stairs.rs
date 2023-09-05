impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
    let mut prev = 1;
    let mut prev2 = 0;
    let mut current = 0;
    for _ in 0..n {
        current = prev + prev2;
        prev2 = prev;
        prev = current;
    }
    prev
    }
}