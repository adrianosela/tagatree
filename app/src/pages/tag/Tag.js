import React from 'react';
import { StyleSheet, Text, View } from 'react-native';

export default class Tag extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <View style={styles.appContainer}>
        <Text style={styles.todoText}>TODO</Text>
      </View>
    );
  }

}

const styles = StyleSheet.create({
  appContainer: {
    flex: 1,
    backgroundColor: '#FFF'
  },
  todoText: {
    fontSize: 200
  }
});
