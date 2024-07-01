# @param {Integer[]} nums
# @return {Integer}
def first_missing_positive(nums)
    nums << 0
    nums.reject! { |e| e.negative? }
    nums.sort!
    nums.uniq!
    i = 0
    while nums[i] + 1 == nums[i+1]
        i += 1
    end
    i+1
end

# def first_missing_positive(nums)
#     max = nums.max
#     return ((1..(max+1)).to_a - nums)[0]
# end

# def first_missing_positive(nums)
#     i = 1
#     while nums.include? i
#         i += 1
#     end
#     i
# end

fail if first_missing_positive([1,2,0]) != 3
fail if first_missing_positive([3,4,-1,1]) != 2
fail if first_missing_positive([7,8,9,11,12]) != 1