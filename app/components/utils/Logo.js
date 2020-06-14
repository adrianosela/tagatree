import React from 'react';
import { StyleSheet, View, Image, Text } from 'react-native';

export default class Logo extends React.Component {
  render() {
    return (
      <View style={styles.logoContainer}>
	<Image
	  style={styles.logo}
	  source={require('../../assets/icon.png')}
	/>
	<Text style={styles.title}>Tag-A-Tree</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  logoContainer: {
    flexGrow: 1,
    alignItems: 'center',
    justifyContent: 'center'
  },
  logo: {
    height: 200,
    width: 200,
  }
});
