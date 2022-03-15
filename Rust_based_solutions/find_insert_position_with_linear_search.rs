fn main() {
    let nums = [1,2,3,4];
    let target = 4;
        for index in nums {
            if nums[index] == target {
                println!("{}", index);
                break;
            } else {
                println!("{}", 0);
                break;
            }
    }
}
