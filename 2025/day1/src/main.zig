const std = @import("std");

const Dial = struct { max: i32, min: i32, current: i32, counter: u32 = 0, range: i32 = 0 };

pub fn main() !void {
    const result: u32 = try start("data/data.txt");
    std.debug.print("Password: {d}\n", .{result});
}

fn start(file: []const u8) !u32 {
    var dial = Dial{ .current = 50, .max = 99, .min = 0 };

    dial.range = dial.max - dial.min + 1;

    const cwd = std.fs.cwd();
    const data = try cwd.openFile(file, .{ .mode = .read_only });
    defer data.close();

    var buffer: [4096]u8 = undefined;
    var reader = data.reader(&buffer);
    var readerInterface = &reader.interface;

    while (true) {
        const line = try readerInterface.takeDelimiter('\n') orelse break;
        const trimmedLine = std.mem.trimRight(u8, line, "\r\n");
        const direction = trimmedLine[0];
        const moves = try std.fmt.parseInt(i32, trimmedLine[1..], 10);

        if (direction == 'R') {
            turnRight(moves, &dial);
        } else if (direction == 'L') {
            turnLeft(moves, &dial);
        }
    }
    return dial.counter;
}

fn turnLeft(moves: i32, dial: *Dial) void {
    std.debug.print("L => Current: {any}, Moves: {any}\n", .{ dial.current, moves });
    var temp = @abs(@divFloor((dial.current - moves), dial.range));
    if (dial.current == 0 and temp > 0) {
        temp -= 1;
    }
    std.debug.print("\tTemp: {any}\n", .{temp});
    if (temp > 0) {
        dial.counter += temp;
    }

    dial.current = dial.min + @mod(dial.current - dial.min - moves, dial.range);
    std.debug.print("\tNew Current: {any}\n", .{dial.current});
    if (dial.current == 0) {
        std.debug.print("\tLanded on 0\n", .{});
        dial.counter += 1;
    }
    std.debug.print("\t\tCurrent counter: {d}\n", .{dial.counter});
}

fn turnRight(moves: i32, dial: *Dial) void {
    std.debug.print("R => Current: {any}, Moves: {any}\n", .{ dial.current, moves });
    const temp = @abs(@divFloor((dial.current + moves - 1), dial.range));
    std.debug.print("\tTemp: {any}\n", .{temp});
    if (temp > 0) {
        dial.counter += temp;
    }

    dial.current = dial.min + @mod(dial.current - dial.min + moves, dial.range);
    std.debug.print("\tNew Current: {any}\n", .{dial.current});
    if (dial.current == 0) {
        dial.counter += 1;
    }
    std.debug.print("\t\tCurrent counter: {d}\n", .{dial.counter});
}

test "check that the example data passes" {
    const result = try start("data/test.txt");
    try std.testing.expectEqual(@as(u32, 6), result);
}
