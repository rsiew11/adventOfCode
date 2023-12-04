import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.ArrayList;

class Solution {

    public static void main(String[] args) throws Exception {
        final var input = Files.readAllLines(Path.of("./input.txt"));
        final var games = formatGames(input);
        System.out.println("part1: " + part1(games));
        System.out.println("part2: " + part2(games));
    }

    private static List<ArrayList<Integer>> formatGames(List<String> input) {
        final var games = new ArrayList<ArrayList<Integer>>();

        for (int i=0; i<input.size(); i++) {
            final var allRounds = input.get(i).split(":")[1];
            games.add(new ArrayList<>(List.of(0,0,0)));
            for (var round : allRounds.split(";")) {
                for (var cube : round.split(",")) {
                    final var idxs = new int[]{cube.indexOf("red"), 
                                               cube.indexOf("green"), 
                                               cube.indexOf("blue")};
                    for (int j=0; j<idxs.length; j++) {
                        if (idxs[j] != -1) {
                            games.get(i).set(j, Math.max(
                                games.get(i).get(j), Integer.valueOf(cube.substring(1, idxs[j]-1))
                            ));
                        }
                    }
                }
            }
        }
        return games;
    }

    private static int part1(List<ArrayList<Integer>> games) {
        var res = 0;
        for (int i=0; i<games.size(); i++) {
            final var game = games.get(i);
            if (game.get(0) <= 12 && game.get(1) <= 13 && game.get(2) <= 14) {
                res += i + 1;
            }
        }
        return res;
    }

    private static int part2(List<ArrayList<Integer>> games) {
        return games.stream()
            .map(rgb -> rgb.stream().reduce(1, (c1,c2) -> c1*c2))
            .reduce((c1,c2) -> c1+c2)
            .get();
    }
    
}
