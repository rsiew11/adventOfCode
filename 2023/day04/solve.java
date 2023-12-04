import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.ArrayDeque;
import java.util.Set;

class Solution {

    public static void main(String[] args) throws Exception {
        final var input = Files.readAllLines(Path.of("./input.txt"));
        final var data = formatInput(input);
        System.out.println("part1: " + part1(data));
        System.out.println("part2: " + part2(data));
        
    }

    private static List<List<Set<Integer>>> formatInput(List<String> input) {
        return input.stream()
            .map(row -> row.split(":")[1])
            .map(row -> row.trim().split(" \\|"))
            .map(row -> List.of(
                    Stream.of(row[0].split(" "))
                        .map(String::trim)
                        .filter(x -> !x.isEmpty())
                        .map(x -> Integer.parseInt(x))
                        .collect(Collectors.toSet()), 
                    Stream.of(row[1].split(" "))
                        .map(String::trim)
                        .filter(x -> !x.isEmpty())
                        .map(x -> Integer.parseInt(x))
                        .collect(Collectors.toSet())
                )
            )
            .toList();
    }

    private static int part1(List<List<Set<Integer>>> data) {
        var res = 0;
        for (int row=0; row<data.size(); row++) {
            final var winning = data.get(row).get(0);
            final var have = data.get(row).get(1);
            final var count = have.stream()
                .filter(x -> winning.contains(x))
                .count();
            if (count > 0) {
                res += Math.pow(2, count - 1);
            }             
        }
        return res;
    }

    private static int part2(List<List<Set<Integer>>> data) {
        var res = data.size(); // original cards
        final var matches = data.stream()
            .map(game -> (int)game.get(1).stream()    // for each card we have
                .filter(c -> game.get(0).contains(c)) // check how many win
                .count()
            )
            .toList();
        final var work = new ArrayDeque<Integer>();
        IntStream.range(0, data.size()).forEach(i -> work.push(i));

        while (!work.isEmpty()) {
            final var gameId = work.pop();
            final var copiesWon = matches.get(gameId);
            if (copiesWon == 0) {
                continue;
            }
            for (int i=1; i<copiesWon+1; i++) {
                work.push(gameId + i);
            }
            res += copiesWon;
        }

        return res;
    }
    
}
