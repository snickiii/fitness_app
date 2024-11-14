// import 'package:dio/dio.dart';
// import 'package:json_annotation/json_annotation.dart';
// import 'package:retrofit/retrofit.dart';

// part 'example.g.dart';

// @RestApi(baseUrl: 'https://5d42a6e2bc64f90014a56ca0.mockapi.io/api/v1/')
// abstract class RestClient {
//   factory RestClient(Dio dio, {String? baseUrl}) = _RestClient;

//   @GET('/tasks')
//   Future<List<Task>> getTasks();
// }

// @JsonSerializable()
// class User {
//   final int id;
//   final String name;
//   final String subname;
//   final String username;
//   final String password;
//   final String mail;
//   final String role;
//   final String createdAt;

//   User({
//     required this.id,
//     required this.name,
//     required this.subname,
//     required this.username,
//     required this.password,
//     required this.mail,
//     required this.role,
//     required this.createdAt
//   });

//   factory User.fromJson(Map<String, dynamic> json) => _$UserFromJson(json);

//   Map<String, dynamic> toJson() => _$UserToJson(this);
// }