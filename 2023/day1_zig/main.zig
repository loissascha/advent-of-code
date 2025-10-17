const std = @import("std");
const print = std.debug.print;

fn readFile(alloc: std.mem.Allocator, path: []const u8) ![]const u8 {
    const cwd = std.fs.cwd();
    return cwd.readFileAlloc(alloc, path, 4096);
}

fn getFileLines(alloc: std.mem.Allocator, path: []const u8) !std.ArrayList([]const u8) {
    const fileContents = try readFile(alloc, path);
    defer alloc.free(fileContents);
    var lines = try std.ArrayList([]const u8).initCapacity(alloc, 16);

    var splits = std.mem.splitAny(u8, fileContents, "\n");
    while (splits.next()) |chunk| {
        const lineCopy = try alloc.dupe(u8, chunk);
        try lines.append(alloc, lineCopy);
    }

    return lines;
}

pub fn main() !void {
    const filename = "test_input.txt";

    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const alloc = gpa.allocator();

    var lines = try getFileLines(alloc, filename);
    defer {
        for (lines.items) |line| alloc.free(line);
        lines.deinit(alloc);
    }

    for (lines.items) |line| {
        print("Line: {s}\n", .{line});
    }
}
