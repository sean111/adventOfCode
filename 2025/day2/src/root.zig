const std = @import("std");
pub fn stringToInt(input: []const u8) !u64 {
    return std.fmt.parseInt(u64, input, 10);
}

pub fn intToString(allocator: std.mem.Allocator, input: u64) ![]u8 {
    return std.fmt.allocPrint(allocator,"{d}", .{input});

}

pub fn checkNumber(input: []const u8) bool {
    const strLen = input.len;
    if (@mod(strLen,2) == 0) {
        const split = (strLen / 2);
        const first = input[0..split];
        const second = input[split..];
        return !std.mem.eql(u8, first, second);
    }
    return true;
}

pub fn checkNumber2(input: []const u8) bool {
    const length = input.len;
    if (length < 2) {
        return true;
    }

    var tmpLen: usize = 1;

    while (tmpLen <= length / 2) : (tmpLen += 1) {

        if (length % tmpLen != 0) {
            continue;
        }

        const pattern = input[0..tmpLen];
        var isRepeating = true;

        var incr: usize = tmpLen;

        while (incr < length): (incr += tmpLen) {
            if (!std.mem.eql(u8, input[incr .. incr + tmpLen], pattern)) {
                isRepeating = false;
                break;
            }

        }

        if (isRepeating) {
            return false;
        }

    }
    return true;
}