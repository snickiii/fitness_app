import 'package:dio/dio.dart';
import 'package:fitness_app/data/userService/user_service.dart';
import 'package:fitness_app/domain/model/user/user.dart';
import 'package:flutter/material.dart';

class WorkoutScreen extends StatefulWidget {

  const WorkoutScreen({super.key});

  @override
  State<WorkoutScreen> createState() => _WorkoutScreenState();
}

class _WorkoutScreenState extends State<WorkoutScreen> {
  late ApiService apiService;
  late Future<User> user;
  final TextEditingController emailController = TextEditingController();
  User? selectedUser;

  @override
  void initState() {
    super.initState();
    Dio dio = Dio();
    apiService = ApiService(dio);

    user = apiService.getUser();
  }

  Future<void> updateEmail() async {
    if (emailController.text.isNotEmpty) {
      // Создаем объект с обновленным email
      User updatedUser = User(
        id: selectedUser!.id,
        username: selectedUser!.username,
        password: selectedUser!.password,
        name: selectedUser!.name,
        subname: selectedUser!.subname,
        email: emailController.text,  // Новый email
      );
      
      try {
        // Отправляем обновленный email на сервер
        User user = await apiService.updateUserEmail(updatedUser);
        setState(() {
          // Обновляем информацию о пользователе на экране
          selectedUser = user;
        });
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Email updated successfully")));
      } catch (e) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Failed to update email")));
      }
    } else {
      ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text("Please enter a valid email")));
    }
  }

  void _showEmailDialog(User user) {
    showDialog(
      context: context,
      builder: (context) {
        return AlertDialog(
          title: Text("Update Email for ${user.name}"),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              TextField(
                controller: emailController,
                decoration: InputDecoration(labelText: "New Email"),
              ),
            ],
          ),
          actions: [
            TextButton(
              onPressed: () {
                Navigator.pop(context);
              },
              child: Text("Cancel"),
            ),
            ElevatedButton(
              onPressed: () {
                updateEmail();  // Обновляем email
                Navigator.pop(context);
              },
              child: Text("Update"),
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color(0xFFFFFFFF),
      appBar: AppBar(
        backgroundColor: Color(0xFF192126),
        title: Text(
          'Тренировка',
          style: TextStyle(
            color: Color(0xFFFFFFFF),
            fontSize: 16,
            fontWeight: FontWeight.w500,
          ),
        ),
        centerTitle: true,
        leading: MaterialButton(
          onPressed: () {
            Navigator.pop(context);
          },
          child: Icon(
            Icons.arrow_back,
            color: Color(0xFFFFFFFF),
          ),
        ),
      ),
      body: FutureBuilder<User>(
        future: user,
        builder: (context, snapshot) {
          if (snapshot.connectionState == ConnectionState.waiting) {
            // print('loading');
            return Center(child: CircularProgressIndicator());
          } else if (snapshot.hasError) {
            // print('error');
            return Center(child: Text('Error: ${snapshot.error}'));
          } else if (snapshot.hasData) {
            User user = snapshot.data!;
            // return Center(child: Text(user.name),);
            print(user.id);
            print(user.username);
            print(user.password);
            print(user.name);
            print(user.subname);
            print(user.email);
            return Container(
              child: ListTile(
                title: Text(user.name),
                  subtitle: Text(user.email),
                  leading: CircleAvatar(
                    child: Text(user.name),
                  ),
                  onTap: () {
                    setState(() {
                      selectedUser = user;
                      emailController.text = user.email; // Заполняем поле ввода текущим email
                    });
                    _showEmailDialog(user);
                  },
              ),
            );
          } else {
            return Center(child: Text('No data available'));
          }
        },
      )
    );
  }
}