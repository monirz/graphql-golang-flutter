import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';
import '../video.dart';

class VideoQuery extends StatelessWidget {
  List<Video> videoList;
  VideoQuery({this.videoList});


  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Query(
      options: QueryOptions(
        document: r"""
          query GetVideos($limit: Int){
            Videos(limit: $limit ){
              name
          }
        } 
      """,
      variables: <String, dynamic>{"limit": 3},
      ),
      builder: (QueryResult result, {BoolCallback refetch}) {
        if (result.errors != null) {
          return Text(result.errors.toString());
        }

        if (result.loading) {
          return Center(
            child: const CircularProgressIndicator(),
          );
        }

        for(int i = 0; i < result.data['Videos'].length; i++) {
          videoList.add(Video(name: result.data['Videos'][i]['name']));
        }

        return Container(child: 
          ListView.builder(
            itemCount: result.data['Videos'].length,
            itemBuilder: (context, index) {
              return Card(child: Text(result.data['Videos'][index]['name']),
              );
            },
          )
        ,);
        //VideoQueryList(videoList: videoList,);
        // Column(
        //   children: <Widget>[
        //     Text(getPrettyJSONString(result.data)),
        //     RaisedButton(
        //       onPressed: () => print(refetch()),
        //       child: const Text('REFETCH'),
        //     ),
        //   ],
        // );
      },
    );
  }
}

class VideoQueryList extends StatefulWidget {
  List<Video> videoList;
  VideoQueryList({this.videoList});

  @override
  createState() => VideoQueryListState();
  
}

class VideoQueryListState extends State<VideoQueryList> {

  @override
  void initState() {
    // TODO: implement initState
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Container(child: 
          ListView.builder(
            itemCount: widget.videoList.length,
            itemBuilder: (context, index) {
              return Card(child: Text(widget.videoList[index].name),
              );
            },
          )
        ,);
  }
} 