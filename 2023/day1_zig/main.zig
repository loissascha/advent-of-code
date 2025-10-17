const std = @import("std");
const print = std.debug.print;

fn readFile(alloc: std.mem.Allocator, path: []const u8) ![]u8 {
    const cwd = std.fs.cwd();
    return cwd.readFileAlloc(alloc, path, 4096);
}

pub fn main() !void {
    const filename = "test_input.txt";

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    const fileContents = try readFile(alloc, filename);
    defer alloc.free(fileContents);

    print("{s}", .{fileContents});
}
