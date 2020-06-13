import React from 'react';
import { StyleSheet, View, Image, Text } from 'react-native';
import LoginForm from './LoginForm';

export default class Login extends React.Component {
  render() {
    return (
      <View style={styles.loginContainer}>
        <View style={styles.logoContainer}>
	  <Image
	    style={styles.logo}
	    source={require('../../../assets/icon.png')}
	  />
	  <Text style={styles.title}>Tag-A-Tree</Text>
        </View>
        <View style={styles.formContainer}>
	    <LoginForm/>
        </View>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  loginContainer: {
    flex: 1,
    backgroundColor: '#81ecec'
  },
  logoContainer: {
    flexGrow: 1,
    alignItems: 'center',
    justifyContent: 'center'
  },
  logo: {
    height: 250,
    width: 250,
  },
  formContainer: {
    flex:1,
  }
});
