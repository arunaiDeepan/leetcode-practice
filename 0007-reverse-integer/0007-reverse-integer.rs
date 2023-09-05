impl Solution {
    pub fn reverse(x: i32) -> i32 {
    const INT_MAX: i32 = i32::MAX;
    const INT_MIN: i32 = i32::MIN;
    let mut rev: i32 = 0;

    let mut is_negative = false;
    let mut x_mut = x;
    if x_mut < 0 {
        is_negative = true;
        x_mut = -x_mut;
    }

    while x_mut != 0 {
        let pop: i32 = x_mut % 10;
        x_mut /= 10;
        if rev > (INT_MAX - pop) / 10 || rev < (INT_MIN + pop) / 10 {
            return 0;
        }
        rev = rev * 10 + pop;
    }
        
    if is_negative {
        rev = -rev;
    }
    rev
    }
}