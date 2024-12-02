import 'package:fitness_app/presentation/pages/food_diary/food_diary_screen.dart';
import 'package:fitness_app/presentation/pages/home/home_screen.dart';
import 'package:fitness_app/presentation/pages/login/login_screen.dart';
import 'package:fitness_app/presentation/pages/workout/workout_screen.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final isAuthorized = await _checkAuthorization();
  runApp(FitnessApp(initialRoute: isAuthorized ? '/home' : '/foodDiaryScreen'));
}

class FitnessApp extends StatelessWidget {
  final String initialRoute;

  const FitnessApp({required this.initialRoute, super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorSchemeSeed: Colors.white,
        fontFamily: 'Lato',
      ),
      initialRoute: initialRoute,
      routes: {
        '/login': (context) => const LoginScreen(),
        '/home': (context) => HomeScreen(),
        '/workoutScreen': (context) => WorkoutScreen(),
        '/foodDiaryScreen': (context) => FoodDiaryScreen(),
      },
      // home: HomeScreen(),
    );
  }
}

Future<bool> _checkAuthorization() async {
  final prefs = await SharedPreferences.getInstance();
  final token = prefs.getString('token');
  return token != null && token.isNotEmpty;
}