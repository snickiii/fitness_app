class ActiveWorkout {
  final String name;
  final String completedTrainingSessions;
  final String totalTrainingSessions;

  ActiveWorkout({
    required this.name,
    required this.completedTrainingSessions,
    required this.totalTrainingSessions,
  });

  static List<ActiveWorkout> getActiveWorkouts() {
    List<ActiveWorkout> activeWorkouts = [];

    activeWorkouts.add(
      ActiveWorkout(name: 'Workout1', completedTrainingSessions: "10", totalTrainingSessions: "100")
    );

    activeWorkouts.add(
      ActiveWorkout(name: 'Workout2', completedTrainingSessions: "20", totalTrainingSessions: "100")
    );

    activeWorkouts.add(
      ActiveWorkout(name: 'Workout3', completedTrainingSessions: "30", totalTrainingSessions: "100")
    );

    return activeWorkouts;
  }
}