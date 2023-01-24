

fn heightmap() -> Vec<Vec<u8>> {
    include_str!("../input.txt")
        .lines()
        .map(|line| line.as_bytes().to_vec())
        .collect::<Vec<_>>()
}

fn find_pos(heightmap: &Vec<Vec<u8>>, loc: u8) -> (usize, usize) {
    for (r, row) in heightmap.iter().enumerate() {
        for (c, val) in row.iter().enumerate() {
            if *val == loc {
                return (r,c);
            }
        }
    }
    unreachable!();
}

fn main() {
    let heightmap = heightmap();
    let (sx, sy) = find_pos(&heightmap, b'S');
    let (ex, ey) = find_pos(&heightmap, b'E');
    println!("{} - {}",sx,sy);
    println!("{} - {}",ex,ey);

}
