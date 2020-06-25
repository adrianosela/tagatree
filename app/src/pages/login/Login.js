import React from 'react';
import PropTypes from 'prop-types';
import { StyleSheet, View, KeyboardAvoidingView } from 'react-native';
import Logo from '../../components/Logo';
import LoginForm from './LoginForm';

export default class Login extends React.Component {
  static propTypes = {
    navigation: PropTypes.shape({
      navigate: PropTypes.func.isRequired,
    }).isRequired,
  }

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <KeyboardAvoidingView behavior='padding' style={styles.loginContainer}>
        <Logo />
        <View style={styles.formContainer}>
          <LoginForm onSuccess={() => this.props.navigation.navigate('Home')}/>
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
