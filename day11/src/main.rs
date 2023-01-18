struct Monkey {
    items: Vec<usize>,
    operation: Box<dyn Fn(usize) -> usize>,
    test: usize,
    pass: usize,
    fail: usize,
    inspection_count: usize,
}


fn main() {
    let monkeys: Vec<_> = include_str!("../input.txt")
        .split("\r\n\r\n")
        .map(|monkey| {
            let fields: Vec<_> = monkey.lines()
                .map(|line| line.split(": ") // separate key : val
                    .last() // value of each field
                    .unwrap())
                .collect();
            Monkey {
                items: fields[1].split(", ").map(|item| item.parse().unwrap()).collect(),
                operation: {
                    let func: Vec<_> = fields[2].split("= ").last().unwrap().split(" ").collect();
                    let (operator, right_hand) = (func[1], func[2]);
                    if right_hand == "old" {
                        Box::new(|old| old * old)
                    } else {
                        match (operator, right_hand.parse::<usize>().unwrap()) {
                            ("+", rh) => Box::new(move |old| old + rh),
                            ("*", rh) => Box::new(move |old| old * rh),
                            _ => unreachable!(),
                        }
                    }
                },
                test: fields[3].split(" ").last().unwrap().parse().unwrap(),
                pass: fields[4].split(" ").last().unwrap().parse().unwrap(),
                fail: fields[5].split(" ").last().unwrap().parse().unwrap(),
                inspection_count: 0,
            }
        })
        .collect();
    for monkey in monkeys.iter(){
        println!("\n");
        println!("{:?}", monkey.items);
        println!("{}",monkey.test);
        println!("{}",monkey.pass);
        println!("{}",monkey.fail);
        println!("{}",monkey.inspection_count);
    }

}
