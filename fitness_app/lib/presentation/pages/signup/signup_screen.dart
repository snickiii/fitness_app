import 'package:flutter/material.dart';

class SignUpScreen extends StatefulWidget {
  const SignUpScreen({super.key});

  @override
  State<SignUpScreen> createState() => _SignUpScreenState();
}

class _SignUpScreenState extends State<SignUpScreen> {
  final usernameController = TextEditingController();
  final passwordController = TextEditingController();
  final firstNameController = TextEditingController();
  final lastNameController = TextEditingController();
  final emailController = TextEditingController();

  void handleRegister() {
    final username = usernameController.text.trim();
    final password = passwordController.text.trim();
    final firstName = firstNameController.text.trim();
    final lastName = lastNameController.text.trim();
    final email = emailController.text.trim();

    if (username.isEmpty || password.isEmpty) {
      showError('Username and password are required');
    } else if (email.isNotEmpty && !isValidEmail(email)) {
      showError('Please enter a valid email address');
    } else {
      print('Registration details:');
      print('Username: $username');
      print('Password: $password');
      print('First Name: $firstName');
      print('Last Name: $lastName');
      print('Email: $email');
      Navigator.pop(context);
    }
  }

  bool isValidEmail(String email) {
    final emailRegex = RegExp(
      r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$',
    );
    return emailRegex.hasMatch(email);
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
      appBar: AppBar(
        title: const Text("Sign Up"),
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(20),
          child: Column(
            children: [
              textField('Username', usernameController),
              const SizedBox(height: 20),
              textField('Password', passwordController, obscureText: true),
              const SizedBox(height: 20),
              textField('First Name', firstNameController),
              const SizedBox(height: 20),
              textField('Last Name', lastNameController),
              const SizedBox(height: 20),
              textField('Email', emailController),
              const SizedBox(height: 20),
              GestureDetector(
                onTap: handleRegister,
                child: Container(
                  height: 50,
                  width: 200,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(12),
                    color: Color(0xFFBBF246),
                  ),
                  child: const Align(
                    alignment: Alignment.center,
                    child: Text(
                      "Register",
                      style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.w700,
                      ),
                    ),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  TextField textField(String hint, TextEditingController controller,
      {bool obscureText = false}) {
    return TextField(
      controller: controller,
      decoration: InputDecoration(
        border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
        hintText: hint,
      ),
      obscureText: obscureText,
    );
  }
}
