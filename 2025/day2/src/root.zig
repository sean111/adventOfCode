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
        std.debug.print("\tCheck: {s} | {s}\n", .{first, second});
        return !std.mem.eql(u8, first, second);
    }
    return true;
}