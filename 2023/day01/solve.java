import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Collectors;

class Solution {


    public static void main(String[] args) throws Exception {
        final var input = Files.readAllLines(Path.of("./input.txt"));

        System.out.println("part1: " + part1(input));
        System.out.println("part2: " + part2(input));
    }

    private static int part1(List<String> input) {
        final var l = input.stream()
            .map(row -> row.chars()
                .mapToObj(c -> (char)c)
                .filter(c -> Character.isDigit(c))
                .findFirst().get())
            .collect(Collectors.toList());
        final var r = input.stream()
            .map(row -> row.chars()
                .mapToObj(c -> (char)c)
                .filter(c -> Character.isDigit(c))
                .reduce((c1,c2) -> c2).get())
            .collect(Collectors.toList());
      
        var res = 0;
        for (int i=0; i<l.size(); i++) {
            res += Integer.parseInt(
                String.valueOf(l.get(i)) + String.valueOf(r.get(i))
            );
        }
        return res;
    }

    private static int part2(List<String> input) {
        return 0;

    }

}