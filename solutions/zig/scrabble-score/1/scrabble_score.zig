const std = @import("std");
pub fn score(s: []const u8) u32 {
    //                            a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q ,r,s,t,u,v,w,x,y,z
    const points: [26]u8 = [26]u8{ 1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10 };
    var total: u32 = 0;
    for (s) |ch| {
        if (std.ascii.isAlphabetic(ch)) {
            total += points[std.ascii.toLower(ch) - 'a'];
        }
    }
    return total;
}
