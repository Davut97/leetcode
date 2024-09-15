function twoSum(nums: number[], target: number): number[] {
  for (let index = 0; index < nums.length; index++) {
    for (let j = index + 1; j < nums.length; j++) {
      console.log(index, j);
      if (nums[index] + nums[j] == target) {
        return [index, j];
      }
    }
  }
}
function twoSumHasMap(nums: number[], target: number): number[] {
  const numberMap = new Map();

  for (let index = 0; index < nums.length; index++) {
    console.log(numberMap);
    const value = nums[index];
    const difference = target - value;
    if (numberMap.has(difference)) {
      return [numberMap.get(difference), index];
    }

    numberMap.set(value, index);
  }
  return [];
}
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
console.log(twoSumHasMap([2, 7, 11, 15], 17));
//console.log(twoSumHasMap([3, 2, 4], 6));
