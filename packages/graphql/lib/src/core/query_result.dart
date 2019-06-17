import 'package:graphql/src/core/graphql_error.dart';

class QueryResult {
  QueryResult({
    this.data,
    this.errors,
    this.loading,
    this.stale,
    this.optimistic = false,
  }) : timestamp = DateTime.now();

  DateTime timestamp;

  /// List<dynamic> or Map<String, dynamic>
  dynamic data;
  List<GraphQLError> errors;
  bool loading;
  // TODO not sure what this is for
  bool stale;
  bool optimistic;

  bool get hasErrors {
    if (errors == null) {
      return false;
    }

    return errors.isNotEmpty;
  }

  void addError(GraphQLError graphQLError) {
    if (errors != null) {
      errors.add(graphQLError);
    } else {
      errors = <GraphQLError>[graphQLError];
    }
  }
}
