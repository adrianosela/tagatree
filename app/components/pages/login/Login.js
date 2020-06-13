import React from 'react';
import { StyleSheet, View, KeyboardAvoidingView, Image, Text } from 'react-native';
import Logo from '../common/Logo';
import LoginForm from './LoginForm';

export default class Login extends React.Component {
  render() {
    return (
      <KeyboardAvoidingView behavior='padding' style={styles.loginContainer}>
	<Logo />
        <View style={styles.formContainer}>
	    <LoginForm/>
        </View>
      </KeyboardAvoidingView>
    );
  }
}

const styles = StyleSheet.create({
  loginContainer: {
    flex: 1,
    backgroundColor: '#81ecec'
  },
  formContainer: {
    flex:1,
  }
});
