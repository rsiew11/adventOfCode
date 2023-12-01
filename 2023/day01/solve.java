import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.ArrayList;
import java.util.stream.Collectors;

class Solution {

    public static void main(String[] args) throws Exception {
        final var input = Files.readAllLines(Path.of("./input.txt"));

        System.out.println("part1: " + part1(input));
        System.out.println("part2: " + part2(input));
    }

    private static int part1(List<String> input) {
        return input.stream()
            .map(row -> row.chars()
                .mapToObj(c -> (char)c)
                .filter(c -> Character.isDigit(c))
                .findFirst().get())
            .map(c -> Integer.valueOf(String.valueOf(c)) * 10)
            .reduce((c1,c2) -> c1+c2).get() 
            + 
            input.stream()
            .map(row -> row.chars()
                .mapToObj(c -> (char)c)
                .filter(c -> Character.isDigit(c))
                .reduce((c1,c2) -> c2).get())
            .map(c -> Integer.valueOf(String.valueOf(c)))
            .reduce((c1,c2) -> c1+c2).get();
    }

    private static int part2(List<String> input) {
        final var validStrs = new String[]{"one", "two", "three", "four", "five", 
                                           "six", "seven", "eight", "nine"};
        final var leftMostVals = input.stream()
            .map(row -> row.chars()
                .mapToObj(c -> (char)c)
                .filter(c -> Character.isDigit(c))
                .findFirst().get())
            .map(c -> String.valueOf(c))
            .collect(Collectors.toList());
        final var rightMostVals = input.stream()
            .map(row -> row.chars()
                .mapToObj(c -> (char)c)
                .filter(c -> Character.isDigit(c))
                .reduce((c1,c2) -> c2).get())
            .map(c -> String.valueOf(c))
            .collect(Collectors.toList());

        final var lIdx = new ArrayList<Integer>();
        final var rIdx = new ArrayList<Integer>();
        for (int i=0; i<input.size(); i++) {
            lIdx.add(input.get(i).indexOf(leftMostVals.get(i)));
            rIdx.add(input.get(i).lastIndexOf(rightMostVals.get(i)));
        }
        
        for (int i=0; i<input.size(); i++) {
            for (int j=0; j<validStrs.length; j++) {
                final var leftLoc = input.get(i).indexOf(validStrs[j]);
                final var rightLoc = input.get(i).lastIndexOf(validStrs[j]);

                if (leftLoc != -1 && leftLoc < lIdx.get(i)) {
                    lIdx.set(i, leftLoc);
                    leftMostVals.set(i, String.valueOf(j + 1));
                }
                if (rightLoc != -1 && rightLoc > rIdx.get(i)) {
                    rIdx.set(i, rightLoc);
                    rightMostVals.set(i, String.valueOf(j + 1));
                }
            }
        }        

        var res = 0;
        for (int i=0; i<input.size(); i++) {
            res += Integer.parseInt(leftMostVals.get(i) + rightMostVals.get(i));
        }
        return res;

    }

}