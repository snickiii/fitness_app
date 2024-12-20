import 'package:flutter/material.dart';

class FoodDiaryScreen extends StatefulWidget {
  const FoodDiaryScreen({super.key});

  @override
  State<FoodDiaryScreen> createState() => _FoodDiaryScreenState();
}

class _FoodDiaryScreenState extends State<FoodDiaryScreen> {
  final Map<String, List<Map<String, dynamic>>> foodDiary = {};

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Food Diary'),
        backgroundColor: Colors.green,
      ),
      body: ListView(
        children: foodDiary.keys.map((day) {
          return _buildDayBlock(day, foodDiary[day]!);
        }).toList(),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: _openAddFoodDialog,
        child: const Icon(Icons.add),
      ),
    );
  }

  Widget _buildDayBlock(String day, List<Map<String, dynamic>> meals) {
    return Padding(
      padding: const EdgeInsets.all(8.0),
      child: Card(
        elevation: 4,
        child: ExpansionTile(
          title: Text(
            day,
            style: const TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
          ),
          children: meals.map((meal) {
            return ListTile(
              title: Text(meal['mealName']),
              subtitle: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: (meal['foods'] as List<Map<String, dynamic>>).map((food) {
                  return Text('${food['name']} - ${food['quantity']}');
                }).toList(),
              ),
            );
          }).toList(),
        ),
      ),
    );
  }

  void _openAddFoodDialog() {
    showDialog(
      context: context,
      builder: (context) {
        return const AddFoodDialog();
      },
    ).then((result) {
      if (result != null) {
        setState(() {
          final today = DateTime.now().toString().split(' ')[0];
          if (!foodDiary.containsKey(today)) {
            foodDiary[today] = [];
          }
          foodDiary[today]!.add(result);
        });
        print('Food diary updated: $foodDiary');
      }
    });
  }
}

class AddFoodDialog extends StatefulWidget {
  const AddFoodDialog({super.key});

  @override
  State<AddFoodDialog> createState() => _AddFoodDialogState();
}

class _AddFoodDialogState extends State<AddFoodDialog> {
  final TextEditingController mealNameController = TextEditingController();
  final TextEditingController searchController = TextEditingController();
  final List<Map<String, dynamic>> searchResults = [];
  final List<Map<String, dynamic>> selectedFoods = [];

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text('Add Food Record'),
      content: SingleChildScrollView(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            TextField(
              controller: mealNameController,
              decoration: const InputDecoration(
                labelText: 'Meal Name',
                border: OutlineInputBorder(),
              ),
            ),
            const SizedBox(height: 16),
            TextField(
              controller: searchController,
              onChanged: _searchFood,
              decoration: const InputDecoration(
                labelText: 'Search Food',
                border: OutlineInputBorder(),
              ),
            ),
            const SizedBox(height: 16),
            ...searchResults.map((food) {
              return ListTile(
                title: Text(food['name']),
                trailing: IconButton(
                  icon: const Icon(Icons.add),
                  onPressed: () {
                    setState(() {
                      selectedFoods.add({...food, 'quantity': 1});
                    });
                  },
                ),
              );
            }).toList(),
            const Divider(),
            const Text('Selected Foods:'),
            ...selectedFoods.map((food) {
              return ListTile(
                title: Text(food['name']),
                subtitle: Text('Quantity: ${food['quantity']}'),
                trailing: Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    IconButton(
                      icon: const Icon(Icons.remove),
                      onPressed: () {
                        setState(() {
                          if (food['quantity'] > 1) {
                            food['quantity'] -= 1;
                          } else {
                            selectedFoods.remove(food);
                          }
                        });
                      },
                    ),
                    IconButton(
                      icon: const Icon(Icons.add),
                      onPressed: () {
                        setState(() {
                          food['quantity'] += 1;
                        });
                      },
                    ),
                  ],
                ),
              );
            }).toList(),
          ],
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.of(context).pop(),
          child: const Text('Cancel'),
        ),
        TextButton(
          onPressed: () {
            final mealName = mealNameController.text.trim();
            if (mealName.isEmpty || selectedFoods.isEmpty) {
              print('Meal name or selected foods cannot be empty');
              return;
            }
            Navigator.of(context).pop({
              'mealName': mealName,
              'foods': selectedFoods,
            });
          },
          child: const Text('Save'),
        ),
      ],
    );
  }

  void _searchFood(String query) {
    if (query.isEmpty) {
      setState(() {
        searchResults.clear();
      });
      return;
    }

    // Заглушка для поиска
    final mockFoods = [
      {'name': 'Apple'},
      {'name': 'Banana'},
      {'name': 'Rice'},
      {'name': 'Chicken'},
    ];

    setState(() {
      searchResults.clear();
      searchResults.addAll(
        mockFoods.where((food) => food['name']!.toLowerCase().contains(query.toLowerCase())),
      );
    });
  }
}
