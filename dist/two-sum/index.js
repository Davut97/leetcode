function twoSum(nums, target) {
    for (let index = 0; index < nums.length; index++) {
        for (let j = index + 1; j < nums.length - 1; j++) {
            console.log(index, j);
            if (nums[index] + nums[j] == target) {
                return [index, j];
            }
        }
    }
}
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
//console.log(twoSum([2, 7, 11, 15], 9));
console.log(twoSum([3, 2, 4], 6));
