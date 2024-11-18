class TodayExersice {
  final String name;
  final String amount;
  final String imagePath;

  TodayExersice({
    required this.name,
    required this.amount,
    required this.imagePath,
  });

  static List<TodayExersice> getTodayExersices() {
    List<TodayExersice> todayExersices = [];

    todayExersices.add(
      TodayExersice(name: 'Push Up', amount: "10", imagePath: 'assets/images/push-up.png')
    );

    todayExersices.add(
      TodayExersice(name: 'Sit Up', amount: "20", imagePath: 'assets/images/push-up.png')
    );

    todayExersices.add(
      TodayExersice(name: 'Knee Push Up', amount: "30", imagePath: 'assets/images/push-up.png')
    );

    return todayExersices;
  }
}