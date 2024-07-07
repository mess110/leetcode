
def permute nums
    result = []

    def backtrack(nums, path, result)
        if nums.length == 0
            result.append(path)
        end

        nums.each_with_index do |e, i|
            new_nums = nums.dup[0...i].concat(nums[i+1..nums.length])
            new_path = path.dup.concat([e])
            backtrack(new_nums, new_path, result)
        end
    end

    backtrack(nums, [], result)

    result
end

p permute([1,2,3])