import 'package:fitness_app/domain/model/active_workout.dart';
import 'package:fitness_app/domain/model/today_exercise.dart';
import 'package:fitness_app/presentation/pages/workout/workout_screen.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const FitnessApp());
}

class FitnessApp extends StatelessWidget {
  const FitnessApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorSchemeSeed: Colors.white,
        fontFamily: 'Lato',
      ),
      initialRoute: '/',
      routes: {
        '/': (context) => HomeScreen(),
        '/workoutScreen': (context) => WorkoutScreen(),
      },
      // home: HomeScreen(),
    );
  }
}

class HomeScreen extends StatelessWidget {
  HomeScreen({super.key});

  final List<ActiveWorkout> activeWorkouts = ActiveWorkout.getActiveWorkouts();
  final List<TodayExersice> todayExersices = TodayExersice.getTodayExersices();

  // void _getActiveWorkouts() {
  //   activeWorkouts = ActiveWorkout.getActiveWorkouts();
  // }

  @override
  Widget build(BuildContext context) {
    // _getActiveWorkouts();
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            _greetUser(),
            _activeWorkoutsList(),
            const SizedBox(height: 24,),
            Padding(
              padding: const EdgeInsets.all(20),
              child: Column(
                children: [
                  const Text(
                    'Today Plan',
                    style: TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                  const SizedBox(height: 18,),
                  ListView.separated(
                    itemCount: todayExersices.length,
                    shrinkWrap: true,
                    physics: NeverScrollableScrollPhysics(),
                    separatorBuilder: (context, index) => SizedBox(height: 18,),
                    itemBuilder: (context, index) {
                      return Container(
                        height: 120,
                        decoration: BoxDecoration(
                          color: Colors.amber.shade200,
                          borderRadius: BorderRadius.circular(23),
                        ),
                        child: Row(
                          children: [
                            Padding(
                              padding: const EdgeInsets.all(10),
                              child: Container(
                                height: 100,
                                width: 100,
                                decoration: BoxDecoration(
                                  color: Colors.white,
                                  borderRadius: BorderRadius.circular(16)
                                ),
                                child: Image.asset(todayExersices[index].imagePath)
                                ),
                            ),
                            Padding(
                              padding: const EdgeInsets.only(left: 12),
                              child: Column(
                                crossAxisAlignment: CrossAxisAlignment.start,
                                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                                children: [
                                  Text(
                                    todayExersices[index].name
                                  ),
                                  Text(
                                    'Amount ${todayExersices[index].amount}'
                                  ),
                                ],
                              ),
                            )
                          ],
                        ),
                      );
                    },
                  ),
                ],
              ),
            )
          ],
        ),
      ),
    );
  }

  Padding _activeWorkoutsList() {
    return Padding(
      padding: const EdgeInsets.only(left: 20),
      child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Padding(
                padding: EdgeInsets.only(left: 4),
                child: Text(
                  'Active Workouts',
                  style: TextStyle(
                    fontSize: 18,
                    fontWeight: FontWeight.w700,
                  ),
                ),
              ),
              const SizedBox(height: 16,),
              Container(
                height: 174,
                // color: Colors.green,
                child: ListView.separated(
                  itemCount: activeWorkouts.length + 1,
                  scrollDirection: Axis.horizontal,
                  separatorBuilder: (context, index) => const SizedBox(width: 20,),
                  itemBuilder: (context, index) {
                    if (index == activeWorkouts.length) {
                      return GestureDetector(
                        onTap: () {
                          Navigator.pushNamed(context, '/workoutScreen');
                        },
                        child: Container(
                          width: 280,
                          decoration: BoxDecoration(
                            color: Colors.blue.shade200,
                            borderRadius: BorderRadius.circular(24),
                          ),
                          child: const Center(
                            child: Text(
                              'Добавить тренировку',
                              style: TextStyle(
                                fontSize: 24,
                                fontWeight: FontWeight.bold,
                              ),
                            ),
                          ),
                        ),
                      );
                    } else {
                      return GestureDetector(
                        onTap: () {print('tap');},
                        child: Container(
                          width: 280,
                          decoration: BoxDecoration(
                            color: Colors.amber.shade200,
                            borderRadius: BorderRadius.circular(24),
                          ),
                          child: Column(
                            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                            children: [
                              Text(
                                activeWorkouts[index].name,
                                style: const TextStyle(
                                  fontSize: 24,
                                  fontWeight: FontWeight.bold,
                                ),
                              ),
                              Text(
                                '${activeWorkouts[index].completedTrainingSessions} / ${activeWorkouts[index].totalTrainingSessions}',
                                style: const TextStyle(
                                  fontSize: 24,
                                  fontWeight: FontWeight.bold,
                                ),
                              ),
                            ],
                          ),
                        ),
                      );
                    }
                  },
                ),
              )
            ],
          ),
    );
  }

  SafeArea _greetUser() {
    return const SafeArea(
      child: Padding(
        padding: EdgeInsets.only(left: 20),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Good Afternoon,',
              style: TextStyle(
                fontSize: 14,
                fontWeight: FontWeight.w600,
              ),
            ),
            Text(
              'User',
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.w800,
              ),
            )
          ]
        )
      ),
    );
  }
}