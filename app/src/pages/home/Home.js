import React from 'react';
import { StyleSheet, View, Text } from 'react-native';

export default class Home extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <View behavior='padding' style={styles.mapContainer}>
        <Text> Hello World! </Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
});
