import React from 'react';
import { StyleSheet, View, Image, Text } from 'react-native';

export default class Login extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.logoContainer}>
	  <Image
	    style={styles.logo}
	    source={require('../../assets/icon.png')}
	  />
	  <Text style={styles.title}>Tag-A-Tree! </Text>
        </View>
        <View style={styles.formContainer}>
        </View>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  loginContainer: {
    flex: 1,
    backgroundColor: '#27ae60'
  },
  logoContainer: {
    alignItems: 'center',
    flexGrow: 1,
    justifyContent: 'center'
  },
  logo: {
    height: 100,
    width: 100,
  },
  formContainer: {
    flex:1,
    backgroundColor: '#9b59b6'
  }
});
