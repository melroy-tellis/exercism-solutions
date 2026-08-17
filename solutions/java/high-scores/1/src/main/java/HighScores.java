import java.util.List;
import java.util.Collections;
import java.util.Comparator;
import java.util.stream.Collectors;

class HighScores {
    private final List<Integer> highScores;
    private final Integer latest;
    private final Integer best;
    private final List<Integer> topThree;
    
    public HighScores(List<Integer> highScores) {
        this.highScores = highScores;
        this.latest = highScores.get(highScores.size()-1);

        final var sortedScores = highScores.stream()
                                .sorted(Comparator.reverseOrder())
                                .collect(Collectors.toList());
        this.best = sortedScores.get(0);
        this.topThree = sortedScores.subList(0, Math.min(3, sortedScores.size()));
    }

    List<Integer> scores() {
        return highScores;
    }

    Integer latest() {
        return latest;
        
    }

    Integer personalBest() {
        return best;
        
    }

    List<Integer> personalTopThree() {
        return topThree;
    }

}
