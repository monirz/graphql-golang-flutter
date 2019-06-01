import 'package:flutter/material.dart';
import 'package:flutter_graphql_demo/client_provider.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import 'package:flutter_graphql_demo/videos_subscription.dart';

void main() {
  runApp(MaterialApp(title: "GQL App", home: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // final HttpLink httpLink =
    //     HttpLink(uri: "http://192.168.1.103:8080/query");
    // final ValueNotifier<GraphQLClient> client = ValueNotifier<GraphQLClient>(
    //   GraphQLClient(
    //     link: httpLink as Link,
    //     cache: OptimisticCache(
    //       dataIdFromObject: typenameDataIdFromObject,
    //     ),
    //   ),
    // );
    return ClientProvider(child: HomePage(),uri: "http://192.168.1.103:8080/query",subscriptionUri: "ws://192.168.1.103:8080/query",);
  }
}

class HomePage extends StatelessWidget {

  final String query = r"""
                query GetVideos($limit: Int){
                          Videos(limit: $limit ){
                                 name
                              }
                          }         
                  """;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("GraphlQL Client"),
      ),
      body: Column(
      mainAxisAlignment: MainAxisAlignment.start,
      crossAxisAlignment: CrossAxisAlignment.start,
      children: <Widget>[
        const ListTile(title: Text('Live Stream of Videos')),
        Expanded(child: VideoFeed()),
      ],
    ),
      // body: Query(
      //   options: QueryOptions(
      //       document: query, variables: <String, dynamic>{"limit": 10}),
      //   builder: (
      //     QueryResult result, {
      //     VoidCallback refetch,
      //   }) {
      //     if (result.loading) {
      //       return Center(child: CircularProgressIndicator());
      //     }
      //     if (result.data == null) {
      //       return Text("No Data Found !");
      //     }
      //     print(result.data["Videos"]);
      //     return VideoFeed();
      //   },
      // ),

    );
  }
}
