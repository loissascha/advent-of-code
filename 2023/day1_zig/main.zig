const std = @import("std");
const print = std.debug.print;

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    const filename = "test_input.txt";

    const fileContents = try read_file(filename);

    print("{s}", .{fileContents});
}

fn read_file(filename: []const u8, alloc: Allocator) ![]u8 {
    const cwd = std.fs.cwd();
    const fileContents = try cwd.readFileAlloc(alloc, filename, 4096);
    defer alloc.free(fileContents);

    return fileContents;
}
