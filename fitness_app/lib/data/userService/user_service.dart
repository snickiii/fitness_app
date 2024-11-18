import 'package:fitness_app/domain/model/user/user.dart';
import 'package:retrofit/retrofit.dart';
import 'package:dio/dio.dart';

part 'user_service.g.dart';

@RestApi(baseUrl: "http://95.181.151.141:5000/")
abstract class ApiService {
  factory ApiService(Dio dio, {String baseUrl}) = _ApiService;

  @GET("/user")
  Future<User> getUser();

  // Метод для обновления email пользователя
  @PUT("/user/update")
  Future<User> updateUserEmail(@Body() User user);
}
