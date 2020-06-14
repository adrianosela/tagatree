import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import Tree from '../../objects/Tree'

export default class Home extends React.Component {
  render() {
    return (
      <View style={styles.appContainer}>
        <Text>((Homepage))</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  appContainer: {
    flex: 1,
    backgroundColor: '#FFF'
  }
});
