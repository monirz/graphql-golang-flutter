import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

void main() {
  runApp(MaterialApp(title: "GQL App", home: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final HttpLink httpLink =
        HttpLink(uri: "http://192.168.1.13:8090/query");
    final ValueNotifier<GraphQLClient> client = ValueNotifier<GraphQLClient>(
      GraphQLClient(
        link: httpLink as Link,
        cache: OptimisticCache(
          dataIdFromObject: typenameDataIdFromObject,
        ),
      ),
    );
    return GraphQLProvider(
      child: HomePage(),
      client: client,
    );
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
      body: Query(
        options: QueryOptions(
            document: query, pollInterval:20, variables: <String, dynamic>{"limit": 10}),
        builder: (
          QueryResult result, {
          VoidCallback refetch,
        }) {
          if (result.loading) {
            return Center(child: CircularProgressIndicator());
          }
          if (result.data == null) {
            return Text("No Data Found !");
          }
          print(result.data["Videos"]);
          return ListView.builder(
            itemBuilder: (BuildContext context, int index) {
              return ListTile(
                title:
                    Text(result.data['Videos'][index]['name']),
              );
            },
            itemCount: result.data["Videos"].length,
          );
        },
      ),
    );
  }
}
