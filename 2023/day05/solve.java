import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.stream.Stream;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.LongStream;
import java.util.Arrays;
import java.util.Comparator;
import java.util.ArrayList;
import java.util.Set;


class Solution {

    record Point(
        long dest,
        long src,
        long rng
    ) {}

    record Pair(
        long start,
        long rng
    ) {}

    public static void main(String[] args) throws Exception {
        final var input = Files.readString(Path.of("./input.txt"));
        final var seeds = getSeeds(input);
        final var maps = getMaps(input);
        
        System.out.println("part1: " + part1(seeds, maps));
        System.out.println("part2: " + part2(seeds, maps));
        
    }

    private static List<Long> getSeeds(String input) {
        return Stream.of(input.split("\n\n")[0].split(":")[1].split(" "))
            .filter(e -> !e.isEmpty())
            .map(e -> Long.parseLong(e))
            .toList();
    }

    private static Point genPoint(String triplet) {
        final var t = triplet.split(" ");
        return new Point(Long.parseLong(t[0]), Long.parseLong(t[1]), Long.parseLong(t[2]));
    }

    private static List<List<Point>> getMaps(String input) {
        return Stream.of(input.split("\n\n")).skip(1)
            .map(m -> Stream.of(m.split("\n")).skip(1)
                .map(line -> genPoint(line))
                .toList()
            )
            .toList();
    }

    private static Long getDestVal(Long val, List<Point> map) {
        final var res = map.stream()
            .filter(p -> p.src <= val  && val < p.src + p.rng)
            .findFirst();
        if (res.isPresent()) {
            final var p = res.get();
            return p.dest + (val - p.src);
        } else {
            return val;
        }
    }

    private static long part1(List<Long> seeds, List<List<Point>> maps) {
        final var locations = new ArrayList<Long>();
        for (var seed : seeds) {
            var val = seed;
            for (var m : maps) {
                val = getDestVal(val, m);
            }
            locations.add(val);
        }
        return locations.stream().min(Comparator.naturalOrder()).get();
    }

    private static long part2(List<Long> seeds, List<List<Point>> maps) {
        final var allSeeds = new ArrayList<Pair>();
        for (int i=0; i<seeds.size(); i += 2) {
            allSeeds.add(new Pair(seeds.get(i), seeds.get(i+1)));
        }

        return allSeeds.parallelStream() // for each (start, range) run on a core
            .map(pair -> LongStream
                .range(pair.start, pair.start + pair.rng)
                    .map(seed -> {
                        var val = seed;
                        for (var m : maps) {
                            val = getDestVal(val, m);
                        }
                        return val;
                    }).min()
            )
            .map(l -> l.getAsLong())
            .min(Comparator.naturalOrder()).get();
    }

}