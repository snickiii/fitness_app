import 'dart:math';

import 'package:flutter/material.dart';
import 'package:fitness_app/presentation/pages/signup/signup_screen.dart';
import 'package:shared_preferences/shared_preferences.dart';

class LoginScreen extends StatefulWidget {
  const LoginScreen({super.key});

  @override
  State<LoginScreen> createState() => _LoginScreenState();
}

class _LoginScreenState extends State<LoginScreen> {
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();

  void handleSignIn() async {
    final username = usernameController.text.trim();
    final password = passwordController.text.trim();

    if (username == 'user' && password == 'user') {
    // Генерация случайного токена
    final token = _generateToken();

    // Сохранение токена
    final prefs = await SharedPreferences.getInstance();
    await prefs.setString('token', token);

    // Переход на HomeScreen
    Navigator.pushReplacementNamed(context, '/home');
    } else {
      showError('Invalid username or password');
    }
  }

  String _generateToken() {
  const chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  final rand = Random();
  return List.generate(32, (index) => chars[rand.nextInt(chars.length)]).join();
}

  void handleSignUp() {
    Navigator.push(
      context,
      MaterialPageRoute(builder: (context) => const SignUpScreen()),
    );
  }

  void showError(String message) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(
        content: Text(message),
        backgroundColor: Colors.red,
      ),
    );
  }
  
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.only(left: 20, right: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(height: 200,),
              loginText(),
              loginDescription(),
              sizedBox20(),
              usernameInputController(),
              sizedBox20(),
              passwordInputController(),
              sizedBox20(),
              signButtons()
            ],
          ),
        ),
      ),
    );
  }

  SizedBox sizedBox20() => const SizedBox(height: 20,);

  Row signButtons() {
    return Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              GestureDetector(
                onTap: handleSignIn,
                child: Container(
                  height: 50,
                  width: 100,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(12),
                    color: Color(0xFFBBF246),
                  ),
                  child: const Align(
                    alignment: Alignment.center,
                    child: Text(
                      "Sign in",
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w700,
                      ),
                    ),
                  ),
                ),
              ),
              GestureDetector(
                onTap: handleSignUp,
                child: Container(
                  height: 50,
                  width: 100,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(12),
                    color: Color(0xFFA48AED),
                  ),
                  child: const Align(
                    alignment: Alignment.center,
                    child: Text(
                      "Sign up",
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w700,
                      ),
                    ),
                  ),
                ),
              ),
            ],
          );
  }

  TextField passwordInputController() {
    return TextField(
            controller: passwordController,
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(12)
              ),
              hintText: 'password'
            ),
            obscureText: true,
          );
  }

  TextField usernameInputController() {
    return TextField(
            controller: usernameController,
            decoration: InputDecoration(
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(12)
              ),
              hintText: 'username'
            )
          );
  }

  Text loginDescription() {
    return const Text(
            'Sign in to continue',
            style: TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.w500,
                    ),
          );
  }

  Text loginText() {
    return const Text(
            "Login",
            style: TextStyle(
                      fontSize: 35,
                      fontWeight: FontWeight.w700,
                    ),
          );
  }
}