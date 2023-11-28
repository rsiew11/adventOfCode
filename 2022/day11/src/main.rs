
struct Monkey {
    items: Vec<usize>,
    operation: Box<dyn Fn(usize) -> usize>,
    test: usize,
    pass: usize,
    fail: usize,
    inspection_count: usize,
}

fn gimme_monkeys() -> Vec<Monkey> {
    include_str!("../input.txt")
        .split("\n\n")
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
        .collect()
}

fn part1(mut monkeys: Vec<Monkey>) {
    for _ in 0..20 {
        for i in 0..monkeys.len() {
            for j in 0..monkeys[i].items.len() {
                let worry = ((monkeys[i].operation)(monkeys[i].items[j])) / 3;
                let receiver = if worry % monkeys[i].test == 0 {monkeys[i].pass} else {monkeys[i].fail};
                monkeys[receiver].items.push(worry);
                monkeys[i].inspection_count += 1;
            } 
            monkeys[i].items.clear();
        }
    }
    monkeys.sort_by_key(|m| m.inspection_count);
    println!("{}",monkeys.iter().rev().take(2).map(|m| m.inspection_count).product::<usize>());
}

fn part2(mut monkeys: Vec<Monkey>) {
    let magic: usize = monkeys.iter().map(|m| m.test).product();
    for _ in 0..10000 {
        for i in 0..monkeys.len() {
            for j in 0..monkeys[i].items.len() {
                let worry = ((monkeys[i].operation)(monkeys[i].items[j])) % magic;
                let receiver = if worry % monkeys[i].test == 0 {monkeys[i].pass} else {monkeys[i].fail};
                monkeys[receiver].items.push(worry);
                monkeys[i].inspection_count += 1;
            } 
            monkeys[i].items.clear();
        }
    }
    monkeys.sort_by_key(|m| m.inspection_count);
    println!("{}",monkeys.iter().rev().take(2).map(|m| m.inspection_count).product::<usize>());
}

fn main() {
    part1(gimme_monkeys());
    part2(gimme_monkeys());
}
