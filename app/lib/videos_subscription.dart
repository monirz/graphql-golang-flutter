import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

// import './review.dart';

class VideoFeed extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    print('this is before');
    return Subscription<Map<String, dynamic>>(
      'videoAdded',
      r'''
        subscription videoAdded {
          videoPublished {
            name
          }
        }
      ''',
      builder: ({dynamic loading, dynamic payload, dynamic error}) {
        if (error != null) {
          print('we are getting this error');
          return Text(error.toString());
        }

        if (loading == true) {
          return Center(
            child: const CircularProgressIndicator(),
          );
        }

        print('here' + payload);

        return VideoList(newVideo: payload as Map<String, dynamic>);
      },
    );
  }
}

class VideoList extends StatefulWidget {
  const VideoList({@required this.newVideo});

  final Map<String, dynamic> newVideo;

  @override
  _VideoListState createState() => _VideoListState();
}

class _VideoListState extends State<VideoList> {
  List<Map<String, dynamic>> videos;

  @override
  void initState() {
    videos = widget.newVideo != null ? [widget.newVideo] : [];
    super.initState();
  }

  @override
  void didUpdateWidget(VideoList oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (!videos.contains(widget.newVideo)) {
      setState(() {
        videos.insert(0, widget.newVideo);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return ListView(
      padding: const EdgeInsets.all(8.0),
      children: videos
          .map(displayVideo)
          .map<Widget>((String s) => Card(
                child: Container(
                  padding: const EdgeInsets.all(15.0),
                  height: 150,
                  child: Text(s),
                ),
              ))
          .toList(),
    );
  }
}

const String Function(Object jsonObject) displayVideo = getPrettyJSONString;

String getPrettyJSONString(Object jsonObject) {
  print('is it me?');
  return const JsonEncoder.withIndent('  ').convert(jsonObject);
}

  // Map<String, dynamic> toJson() {
  //   assert(episode != null && stars != null);

  //   return <String, dynamic>{
  //     'episode': episodeToJson(episode),
  //     'stars': stars,
  //     'commentary': commentary,
  //   };
  // }

  // static Review fromJson(Map<String, dynamic> map) => Review(
  //       episode: episodeFromJson(map['episode'] as String),
  //       stars: map['stars'] as int,
  //       commentary: map['commentary'] as String,
  //     );