import 'package:flutter/material.dart';
import 'package:flutter_graphql_demo/client_provider.dart';
import 'package:flutter_graphql_demo/queries/video_query.dart';
import 'package:flutter_graphql_demo/queries/videos_subscription.dart';
import './video.dart';


void main() {
  runApp(MaterialApp(title: "GQL App", home: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return  ClientProvider(child: HomePage(),
              uri: "http://192.168.1.13:8090/query",
              subscriptionUri: "ws://192.168.1.13:8090/query",
            );
  }
}

class HomePage extends StatelessWidget {

 List<Video> videoList = [];


  // final String query = r"""
  //               query GetVideos($limit: Int){
  //                         Videos(limit: $limit ){
  //                                name
  //                             }
  //                         }         
  //                 """;

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
        Expanded(child: VideoFeed(),),
        Expanded(child: VideoQuery(videoList: videoList,)),
      ],
    ),

    floatingActionButton: FloatingActionButton(
                  tooltip: 'Add', child: new Icon(Icons.add), onPressed: () {})

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
