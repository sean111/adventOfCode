const std = @import("std");
const day2 = @import("day2");

const TestInputs = struct {
    stringValue: []const u8 = "",
    intValue: u64 = 0,
    result: bool = false,
};

const inputFile = "data/input.txt";
const groupingDelimiter = ',';
const rangeDelimiter = '-';
var total: u64 = 0;
var total2: u64 = 0;

pub fn main() !void {
    const fileAllocator = std.heap.page_allocator;

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();

    // const cwd = std.fs.cwd();
    // const file = try cwd.openFile("data/test.txt", .{.mode = .read_only});
    // defer file.close();

    const data = try std.fs.cwd().readFileAlloc(fileAllocator, inputFile, std.math.maxInt(usize));
    defer fileAllocator.free(data);

    // std.debug.print("Contents:\n\n{s}\n", .{data});
    var splits = std.mem.splitScalar(u8, data, groupingDelimiter);

    while (splits.next()) |split| {
        var first: []const u8 = "";
        var second: []const u8 = "";
        // std.debug.print("Split: {s}\n", .{split});
        var temp = std.mem.splitScalar(u8, split, rangeDelimiter);
        while(temp.next()) |number| {
            if (std.mem.eql(u8,first,"")) {
                first = number;
            } else {
                second = number;
            }
        }
        // std.debug.print("First: {s}, Second: {s}\n", .{first, second});

        const firstInt = try day2.stringToInt(first);
        const secondInt = try day2.stringToInt(second);

        if (!day2.checkNumber(first)) {
            std.debug.print("Part 1 Invalid ID: {s}", .{first});
            total+=firstInt;
        }

        if (!day2.checkNumber2(first)) {
            std.debug.print("Part2 Invalid ID: {s}\n", .{first});
            total2+=firstInt;
        }

        if (!day2.checkNumber(second)) {
            std.debug.print("Part 1 Invalid ID: {s}", .{second});
            total+=secondInt;
        }

        if (!day2.checkNumber2(second)) {
            std.debug.print("Part 2 Invalid ID: {s}", .{second});
            total2+=secondInt;
        }



        for (firstInt+1..secondInt) |i| {
            // std.debug.print("\tI: {d}\n", .{i});
            const iString = try day2.intToString(allocator, i);
            defer allocator.free(iString);

            if (!day2.checkNumber(iString)) {
                std.debug.print("Part1 Invalid ID: {s}\n", .{iString});
                total+=i;
            }

            if (!day2.checkNumber2(iString)) {
                std.debug.print("Part2 Invalid ID: {s}\n", .{iString});
                total2+=i;
            }
        }

    }

    std.debug.print("Total: {d}\n", .{total});
    std.debug.print("Total2: {d}\n", .{total2});

}



test "test checkNumber" {
    const tests = [_]TestInputs{
        .{.stringValue = "112112", .result = false},
        .{.stringValue = "123456", .result = true},
        .{.stringValue = "119900229933399119900229933399", .result = false},
        .{.stringValue = "824824824", .result = true},
    };

    for (tests) |testItem| {
        const check = day2.checkNumber(testItem.stringValue);
        try std.testing.expectEqual(testItem.result, check);
    }
}

test "test stringToInt" {
    const tests = [_]TestInputs{
        .{.stringValue = "112", .intValue = 112},
        .{.stringValue = "1227775554", .intValue = 1227775554}
    };

    for (tests) |testItem| {
        const check = try day2.stringToInt(testItem.stringValue);
        try std.testing.expectEqual(testItem.intValue, check);
    }
}

test "test intoToString" {
    const tests = [_]TestInputs{
        .{.intValue = 112, .stringValue = "112"},
        .{.intValue = 100111222, .stringValue = "100111222"},
        .{.intValue = 1227775554, .stringValue = "1227775554"}
    };
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();
    defer _ = gpa.deinit();
    for (tests) |testItem| {
        const check = try day2.intToString(allocator, testItem.intValue);
        defer allocator.free(check);
        try std.testing.expectEqualStrings(testItem.stringValue, check);
    }
}

test "test checkNumber2" {
    const tests = [_]TestInputs{
        .{.stringValue = "112112", .result = false},
        .{.stringValue = "123456", .result = true},
        .{.stringValue = "824824824", .result = false},
        .{.stringValue = "565656", .result = false},
        .{.stringValue = "2121212121", .result = false},
        .{.stringValue = "999", .result = false},
        .{.stringValue = "11", .result = false},
    };

    for (tests) |testItem| {
        const check = day2.checkNumber2(testItem.stringValue);
        try std.testing.expectEqual(testItem.result, check);
    }
}